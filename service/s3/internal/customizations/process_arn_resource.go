package customizations

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/transport/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsarn "github.com/aws/aws-sdk-go-v2/aws/arn"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared"
	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared/arn"
)

// processARNResourceMiddleware is used to process an ARN resource.
type processARNResourceMiddleware struct {
	// GetARN points to a function that processes an input and returns ARN as string ptr,
	// and bool indicating if ARN is supported or set.
	GetARN func(interface{}) (*string, bool)

	// UseARNRegion indicates if region parsed from an ARN should be used.
	UseARNRegion bool

	// UseAccelerate indicates if s3 transfer acceleration is enabled
	UseAccelerate bool

	// UseDualstack instructs if s3 dualstack endpoint config is enabled
	UseDualstack bool
}

// ID returns the middleware ID.
func (*processARNResourceMiddleware) ID() string { return "S3:ProcessARNResourceMiddleware" }

func (m *processARNResourceMiddleware) HandleSerialize(
	ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler,
) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*http.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown request type %T", req)
	}

	// get arn from an arnable field.
	v, ok := m.GetARN(in.Parameters)
	if !ok || v == nil {
		return next.HandleSerialize(ctx, in)
	}

	resource, err := parseEndpointARN(*v)
	if err != nil {
		return out, metadata, fmt.Errorf("Error parsing resource arn %w", err)
	}

	resourceRequest := s3shared.ResourceRequest{
		Resource: resource,
		Request:  req,
		Options: s3shared.ResourceRequestOptions{
			RequestRegion: awsmiddleware.GetRegion(ctx),
			SigningRegion: awsmiddleware.GetSigningRegion(ctx),
			PartitionID:   awsmiddleware.GetPartitionID(ctx),
			UseARNRegion:  m.UseARNRegion,
			// TODO: change HasCustomEndpoint value to use HostImmutable from ctx
			HasCustomEndpoint: false,
		},
	}

	// validate resource request
	if err := validateResourceRequest(resourceRequest); err != nil {
		return out, metadata, err
	}

	// switch to correct endpoint updater
	switch tv := resource.(type) {
	case arn.AccessPointARN:
		// validations
		// check if accelerate
		if m.UseAccelerate {
			return out, metadata, s3shared.NewClientConfiguredForAccelerateError(tv,
				resourceRequest.Options.PartitionID, resourceRequest.Options.RequestRegion, nil)
		}

		// TODO : Do we provide an option to disable host prefix behavior ? If yes, that option should be set to false

		resolveRegion := tv.Region
		resolveService := tv.Service
		// check if request region is FIPS
		if resourceRequest.UseFips() {
			// if use arn region is enabled and request signing region is not same as arn region
			if m.UseARNRegion && resourceRequest.IsCrossRegion() {
				// FIPS with cross region is not supported, the SDK must fail
				// because there is no well defined method for SDK to construct a
				// correct FIPS endpoint.
				return out, metadata,
					s3shared.NewClientConfiguredForCrossRegionFIPSError(
						tv,
						resourceRequest.Options.PartitionID,
						resourceRequest.Options.RequestRegion,
						nil,
					)
			}

			// if use arn region is NOT set, we should use the request region
			resolveRegion = resourceRequest.Options.RequestRegion
		}

		// resolve regional endpoint for resolved region.
		// we do not support custom endpoint resolver for this middleware behavior
		var endpointResolver aws.EndpointResolver
		endpoint, err := endpointResolver.ResolveEndpoint(resolveService, resolveRegion)
		if err != nil {
			return out, metadata, s3shared.NewFailedToResolveEndpointError(
				tv,
				resourceRequest.Options.PartitionID,
				resourceRequest.Options.RequestRegion,
				err,
			)
		}

		req.URL, err = url.Parse(endpoint.URL)
		if err != nil {
			return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
		}

		const serviceEndpointLabel = "s3-accesspoint"
		cfgHost := req.URL.Host
		if strings.HasPrefix(cfgHost, "s3") {
			// replace service hostlabel "s3" to "s3-accesspoint"
			resourceRequest.Request.URL.Host = serviceEndpointLabel + cfgHost[len("s3"):]
		}

		// redirect signer to use resolved endpoint signing name and region
		ctx = awsmiddleware.SetSigningName(ctx, endpoint.SigningName)
		ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)

		// add host prefix for s3-accesspoint
		accessPointHostPrefix := "{" + tv.AccessPointName + "}-{" + tv.AccountID + "}."
		req.URL.Host = accessPointHostPrefix + req.URL.Host
		if len(req.Host) > 0 {
			req.Host = accessPointHostPrefix + req.Host
		}

		// TODO: Need another middleware to remove arnable field from path
		// remove the serialized arn in place in place of /{Bucket}
		// removeBucketFromPath(req.URL, tv.String())

	// process outpost accesspoint ARN
	case arn.OutpostAccessPointARN:
		// validations
		// check if accelerate
		if m.UseAccelerate {
			return out, metadata, s3shared.NewClientConfiguredForAccelerateError(tv,
				resourceRequest.Options.PartitionID, resourceRequest.Options.RequestRegion, nil)
		}

		// check if dual stack
		if m.UseDualstack {
			return out, metadata, s3shared.NewClientConfiguredForDualStackError(tv,
				resourceRequest.Options.PartitionID, resourceRequest.Options.RequestRegion, nil)
		}

		// TODO : Do we provide an option to disable host prefix behavior ? If yes, that option should be set to false

		// check if resource arn region is FIPS
		if resourceRequest.ResourceConfiguredForFIPS() {
			return out, metadata, s3shared.NewInvalidARNWithFIPSError(tv, nil)
		}

		resolveRegion := tv.Region
		resolveService := tv.Service
		endpointsID := resolveService
		if strings.EqualFold(resolveService, "s3-outposts") {
			// assign endpoints ID as "S3"
			endpointsID = "s3"
		}

		// resolve regional endpoint for resolved region.
		// we do not support custom endpoint resolver for this middleware behavior
		var endpointResolver aws.EndpointResolver
		endpoint, err := endpointResolver.ResolveEndpoint(endpointsID, resolveRegion)
		if err != nil {
			return out, metadata, s3shared.NewFailedToResolveEndpointError(
				tv,
				resourceRequest.Options.PartitionID,
				resourceRequest.Options.RequestRegion,
				err,
			)
		}

		// assign resolved endpoint url to request url
		req.URL, err = url.Parse(endpoint.URL)
		if err != nil {
			return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
		}

		cfgHost := req.URL.Host
		if strings.HasPrefix(cfgHost, endpointsID) {
			// replace service endpointID label with resolved service
			resourceRequest.Request.URL.Host = resolveService + cfgHost[len(endpointsID):]
		}

		// redirect signer to use resolved endpoint signing name and region
		ctx = awsmiddleware.SetSigningName(ctx, endpoint.SigningName)
		ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)

		// add host prefix for s3-outposts
		outpostAPHostPrefix := "{" + tv.AccessPointName + "}-{" + tv.AccountID + "}." + "{" + tv.OutpostID + "}"
		req.URL.Host = outpostAPHostPrefix + req.URL.Host
		if len(req.Host) > 0 {
			req.Host = outpostAPHostPrefix + req.Host
		}

		// validate the endpoint host
		if err := http.ValidateEndpointHost(req.URL.Host); err != nil {
			return out, metadata, s3shared.NewInvalidARNError(tv, err)
		}

		// TODO: Need another middleware to remove arnable field from path
		// remove the serialized arn in place in place of /{Bucket}
		// removeBucketFromPath(req.URL, tv.String())
	default:
		return out, metadata, s3shared.NewInvalidARNError(resource, nil)
	}

	return next.HandleSerialize(ctx, in)
}

// validate if s3 resource and request config is compatible.
func validateResourceRequest(resourceRequest s3shared.ResourceRequest) error {
	// check if resourceRequest leads to a cross partition error
	v, err := resourceRequest.IsCrossPartition()
	if err != nil {
		return err
	}
	if v {
		// if cross partition
		return s3shared.NewClientPartitionMismatchError(resourceRequest.Resource,
			resourceRequest.Options.PartitionID, resourceRequest.Options.RequestRegion, nil)
	}

	// check if resourceRequest leads to a cross region error
	if !resourceRequest.AllowCrossRegion() && resourceRequest.IsCrossRegion() {
		// if cross region, but not use ARN region is not enabled
		return s3shared.NewClientRegionMismatchError(resourceRequest.Resource,
			resourceRequest.Options.PartitionID, resourceRequest.Options.RequestRegion, nil)
	}

	// check if resourceRequest leads to a custom endpoint error
	if resourceRequest.HasCustomEndpoint() {
		// if a custom endpoint was provided, return an error
		return s3shared.NewInvalidARNWithCustomEndpointError(resourceRequest.Resource, nil)
	}

	return nil
}

// Used by shapes with members decorated as endpoint ARN.
func parseEndpointARN(v string) (arn.Resource, error) {
	return arn.ParseResource(v, accessPointResourceParser)
}

func accessPointResourceParser(a awsarn.ARN) (arn.Resource, error) {
	resParts := arn.SplitResource(a.Resource)
	switch resParts[0] {
	case "accesspoint":
		if a.Service != "s3" {
			return arn.AccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "service is not s3"}
		}
		return arn.ParseAccessPointResource(a, resParts[1:])
	case "outpost":
		if a.Service != "s3-outposts" {
			return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "service is not s3-outposts"}
		}
		return parseOutpostAccessPointResource(a, resParts[1:])
	default:
		return nil, arn.InvalidARNError{ARN: a, Reason: "unknown resource type"}
	}
}

func parseOutpostAccessPointResource(a awsarn.ARN, resParts []string) (arn.OutpostAccessPointARN, error) {
	// outpost accesspoint arn is only valid if service is s3-outposts
	if a.Service != "s3-outposts" {
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "service is not s3-outposts"}
	}

	if len(resParts) == 0 {
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "outpost resource-id not set"}
	}

	if len(resParts) < 3 {
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{
			ARN: a, Reason: "access-point resource not set in Outpost ARN",
		}
	}

	resID := strings.TrimSpace(resParts[0])
	if len(resID) == 0 {
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "outpost resource-id not set"}
	}

	var outpostAccessPointARN = arn.OutpostAccessPointARN{}
	switch resParts[1] {
	case "accesspoint":
		accessPointARN, err := arn.ParseAccessPointResource(a, resParts[2:])
		if err != nil {
			return arn.OutpostAccessPointARN{}, err
		}
		// set access-point arn
		outpostAccessPointARN.AccessPointARN = accessPointARN
	default:
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "access-point resource not set in Outpost ARN"}
	}

	// set outpost id
	outpostAccessPointARN.OutpostID = resID
	return outpostAccessPointARN, nil
}
