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
	s3endpoints "github.com/aws/aws-sdk-go-v2/service/s3control/internal/endpoints/s3"
)

const (
	// outpost id header
	outpostIDHeader = "x-amz-outpost-id"

	// account id header
	accountIDHeader = "x-amz-account-id"
)

// processARNResourceMiddleware is used to process an ARN resource.
type processARNResourceMiddleware struct {
	// UpdateARNField points to a function that takes in a copy of input, updates the ARN field with
	// the provided value and returns the input copy, along with a bool indicating if field supports ARN
	UpdateARNField func(interface{}, string) (interface{}, bool)

	// UseARNRegion indicates if region parsed from an ARN should be used.
	UseARNRegion bool

	// UseDualstack instructs if s3 dualstack endpoint config is enabled
	UseDualstack bool

	// EndpointResolver used to resolve endpoints. This may be a custom endpoint resolver
	EndpointResolver EndpointResolver

	// EndpointResolverOptions used by endpoint resolver
	EndpointResolverOptions EndpointResolverOptions
}

// ID returns the middleware ID.
func (*processARNResourceMiddleware) ID() string { return "S3Control:ProcessARNResourceMiddleware" }

func (m *processARNResourceMiddleware) HandleSerialize(
	ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler,
) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	// check if arn was provided, if not skip this middleware
	arnValue, ok := s3shared.GetARNResourceFromContext(ctx)
	if !ok {
		return next.HandleSerialize(ctx, in)
	}

	// TODO: verify behavior in case arn region resolves to custom endpoint that is mutable

	req, ok := in.Request.(*http.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown request type %T", req)
	}

	// parse arn into an endpoint arn wrt to service
	resource, err := parseEndpointARN(arnValue)
	if err != nil {
		return out, metadata, err
	}

	resourceRequest := s3shared.ResourceRequest{
		Resource:      resource,
		RequestRegion: awsmiddleware.GetRegion(ctx),
		SigningRegion: awsmiddleware.GetSigningRegion(ctx),
		PartitionID:   awsmiddleware.GetPartitionID(ctx),
		UseARNRegion:  m.UseARNRegion,
	}

	// validate resource request
	if err := validateResourceRequest(resourceRequest); err != nil {
		return out, metadata, err
	}

	// switch to correct endpoint updater
	switch tv := resource.(type) {
	case arn.OutpostAccessPointARN:
		// validations
		// check if dual stack
		if m.UseDualstack {
			return out, metadata, s3shared.NewClientConfiguredForDualStackError(tv,
				resourceRequest.PartitionID, resourceRequest.RequestRegion, nil)
		}

		// check if resource arn region is FIPS
		if resourceRequest.ResourceConfiguredForFIPS() {
			return out, metadata, s3shared.NewInvalidARNWithFIPSError(tv, nil)
		}

		// TODO : Do we provide an option to disable host prefix behavior ? If yes, that option should be set to false

		// update the arnable field with access point name
		input, ok := m.UpdateARNField(in.Parameters, tv.AccessPointName)
		if !ok {
			return out, metadata, fmt.Errorf("error updating arnable field while serializing")
		}
		in.Parameters = input

		// Build outpost access point resource
		resolveRegion := tv.Region
		resolveService := tv.Service
		// check if request region is FIPS and ARN region usage is not allowed
		if resourceRequest.UseFips() && !m.UseARNRegion {
			return out, metadata, s3shared.NewInvalidARNWithFIPSError(tv, nil)
		}

		endpointsID := resolveService
		if resolveService == "s3-outposts" {
			endpointsID = "s3"
		}

		// resolve regional endpoint for resolved region.
		var endpoint aws.Endpoint
		if endpointsID == "s3" {
			// TODO: no way to know if customer has provided a custom resolver anticipating this behavior
			//  By default a provided endpoint resolver has a fallback to s3-control default resolver
			//  that would prevent endpointID switch and cross region endpoint resolver behavior.

			// use s3 endpoint resolver
			endpoint, err = s3endpoints.New().ResolveEndpoint(resolveRegion, s3endpoints.Options{
				DisableHTTPS: m.EndpointResolverOptions.DisableHTTPS,
			})
		} else {
			endpoint, err = m.EndpointResolver.ResolveEndpoint(resolveRegion, m.EndpointResolverOptions)
		}

		if err != nil {
			return out, metadata, s3shared.NewFailedToResolveEndpointError(
				tv,
				resourceRequest.PartitionID,
				resourceRequest.RequestRegion,
				err,
			)
		}

		req.URL, err = url.Parse(endpoint.URL)
		if err != nil {
			return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
		}

		// redirect signer to use resolved endpoint signing name and region
		if len(endpoint.SigningName) != 0 {
			ctx = awsmiddleware.SetSigningName(ctx, endpoint.SigningName)
		} else {
			// assign resolved service from arn as signing name
			ctx = awsmiddleware.SetSigningName(ctx, resolveService)
		}

		if len(endpoint.SigningRegion) != 0 {
			// redirect signer to use resolved endpoint signing name and region
			ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)
		} else {
			ctx = awsmiddleware.SetSigningRegion(ctx, resolveRegion)
		}

		// skip arn processing, if arn region resolves to a immutable endpoint
		if endpoint.HostnameImmutable {
			return next.HandleSerialize(ctx, in)
		}

		// Add outpostID header
		req.Header.Add(outpostIDHeader, tv.OutpostID)

		// add url host as s3-outposts
		cfgHost := req.URL.Host
		if strings.HasPrefix(cfgHost, endpointsID) {
			req.URL.Host = resolveService + cfgHost[len(endpointsID):]
		}

		// validate the endpoint host
		if err := http.ValidateEndpointHost(req.URL.Host); err != nil {
			return out, metadata, s3shared.NewInvalidARNError(tv, err)
		}

		// TODO: Disable endpoint host prefix for s3-control

	// process outpost accesspoint ARN
	case arn.OutpostBucketARN:
		// validations
		// check if dual stack
		if m.UseDualstack {
			return out, metadata, s3shared.NewClientConfiguredForDualStackError(tv,
				resourceRequest.PartitionID, resourceRequest.RequestRegion, nil)
		}

		// check if resource arn region is FIPS
		if resourceRequest.ResourceConfiguredForFIPS() {
			return out, metadata, s3shared.NewInvalidARNWithFIPSError(tv, nil)
		}

		// TODO : Do we provide an option to disable host prefix behavior ? If yes, that option should be set to false

		// update the arnable field with bucket name
		input, ok := m.UpdateARNField(in.Parameters, tv.BucketName)
		if !ok {
			return out, metadata, fmt.Errorf("error updating arnable field while serializing")
		}
		in.Parameters = input

		// Build endpoint from outpost bucket arn
		resolveRegion := tv.Region
		resolveService := tv.Service
		// Outpost bucket resource uses `s3-control` as serviceEndpointLabel
		endpointsID := "s3-control"

		// resolve regional endpoint for resolved region.
		var endpoint aws.Endpoint
		endpoint, err = m.EndpointResolver.ResolveEndpoint(resolveRegion, m.EndpointResolverOptions)
		if err != nil {
			return out, metadata, s3shared.NewFailedToResolveEndpointError(
				tv,
				resourceRequest.PartitionID,
				resourceRequest.RequestRegion,
				err,
			)
		}

		// assign resolved endpoint url to request url
		req.URL, err = url.Parse(endpoint.URL)
		if err != nil {
			return out, metadata, fmt.Errorf("failed to parse endpoint URL: %w", err)
		}

		if len(endpoint.SigningName) != 0 {
			ctx = awsmiddleware.SetSigningName(ctx, endpoint.SigningName)
		} else {
			// assign resolved service from arn as signing name
			ctx = awsmiddleware.SetSigningName(ctx, resolveService)
		}

		if len(endpoint.SigningRegion) != 0 {
			// redirect signer to use resolved endpoint signing name and region
			ctx = awsmiddleware.SetSigningRegion(ctx, endpoint.SigningRegion)
		} else {
			ctx = awsmiddleware.SetSigningRegion(ctx, resolveRegion)
		}

		// skip arn processing, if arn region resolves to a immutable endpoint
		if endpoint.HostnameImmutable {
			return next.HandleSerialize(ctx, in)
		}

		// Add outpostID header
		req.Header.Add(outpostIDHeader, tv.OutpostID)

		cfgHost := req.URL.Host
		if strings.HasPrefix(cfgHost, endpointsID) {
			// replace service endpointID label with resolved service
			req.URL.Host = resolveService + cfgHost[len(endpointsID):]
		}

		// validate the endpoint host
		if err := http.ValidateEndpointHost(req.URL.Host); err != nil {
			return out, metadata, s3shared.NewInvalidARNError(tv, err)
		}

		// TODO: disable host prefix for s3-control
	default:
		return out, metadata, s3shared.NewInvalidARNError(resource, nil)
	}

	// Add account-id header for the request if not present.
	// SDK must always send the x-amz-account-id header for all requests
	// where an accountId has been extracted from an ARN or the accountId field modeled as a header.
	if h := req.Header.Get(accountIDHeader); len(h) == 0 {
		req.Header.Add(accountIDHeader, resource.GetARN().AccountID)
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
			resourceRequest.PartitionID, resourceRequest.RequestRegion, nil)
	}

	// check if resourceRequest leads to a cross region error
	if !resourceRequest.AllowCrossRegion() && resourceRequest.IsCrossRegion() {
		// if cross region, but not use ARN region is not enabled
		return s3shared.NewClientRegionMismatchError(resourceRequest.Resource,
			resourceRequest.PartitionID, resourceRequest.RequestRegion, nil)
	}

	// resource configured with FIPS as region is not supported by outposts
	if resourceRequest.ResourceConfiguredForFIPS() {
		return s3shared.NewInvalidARNWithFIPSError(resourceRequest.Resource, nil)
	}

	return nil
}

// Used by shapes with members decorated as endpoint ARN.
func parseEndpointARN(v awsarn.ARN) (arn.Resource, error) {
	return arn.ParseResource(v, resourceParser)
}

func resourceParser(a awsarn.ARN) (arn.Resource, error) {
	resParts := arn.SplitResource(a.Resource)
	switch resParts[0] {
	case "outpost":
		return arn.ParseOutpostARNResource(a, resParts[1:])
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
