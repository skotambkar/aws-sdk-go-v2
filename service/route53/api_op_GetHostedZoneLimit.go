// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the specified limit for a specified hosted zone, for example, the maximum
// number of records that you can create in the hosted zone. For the default limit,
// see Limits
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To request a higher limit, open a case
// (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-route53).
func (c *Client) GetHostedZoneLimit(ctx context.Context, params *GetHostedZoneLimitInput, optFns ...func(*Options)) (*GetHostedZoneLimitOutput, error) {
	if params == nil {
		params = &GetHostedZoneLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHostedZoneLimit", params, optFns, addOperationGetHostedZoneLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHostedZoneLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the request to create a hosted
// zone.
type GetHostedZoneLimitInput struct {

	// The ID of the hosted zone that you want to get a limit for.
	//
	// This member is required.
	HostedZoneId *string

	// The limit that you want to get. Valid values include the following:
	//
	// *
	// MAX_RRSETS_BY_ZONE: The maximum number of records that you can create in the
	// specified hosted zone.
	//
	// * MAX_VPCS_ASSOCIATED_BY_ZONE: The maximum number of
	// Amazon VPCs that you can associate with the specified private hosted zone.
	//
	// This member is required.
	Type types.HostedZoneLimitType
}

// A complex type that contains the requested limit.
type GetHostedZoneLimitOutput struct {

	// The current number of entities that you have created of the specified type. For
	// example, if you specified MAX_RRSETS_BY_ZONE for the value of Type in the
	// request, the value of Count is the current number of records that you have
	// created in the specified hosted zone.
	//
	// This member is required.
	Count int64

	// The current setting for the specified limit. For example, if you specified
	// MAX_RRSETS_BY_ZONE for the value of Type in the request, the value of Limit is
	// the maximum number of records that you can create in the specified hosted zone.
	//
	// This member is required.
	Limit *types.HostedZoneLimit

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetHostedZoneLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetHostedZoneLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetHostedZoneLimit{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetHostedZoneLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHostedZoneLimit(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetHostedZoneLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHostedZoneLimit",
	}
}
