// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Lightsail content delivery network (CDN) distribution. A
// distribution is a globally distributed network of caching servers that improve
// the performance of your website or web application hosted on a Lightsail
// instance. For more information, see Content delivery networks in Amazon
// Lightsail
// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-content-delivery-network-distributions).
func (c *Client) CreateDistribution(ctx context.Context, params *CreateDistributionInput, optFns ...func(*Options)) (*CreateDistributionOutput, error) {
	if params == nil {
		params = &CreateDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDistribution", params, optFns, c.addOperationCreateDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDistributionInput struct {

	// The bundle ID to use for the distribution. A distribution bundle describes the
	// specifications of your distribution, such as the monthly cost and monthly
	// network transfer quota. Use the GetDistributionBundles action to get a list of
	// distribution bundle IDs that you can specify.
	//
	// This member is required.
	BundleId *string

	// An object that describes the default cache behavior for the distribution.
	//
	// This member is required.
	DefaultCacheBehavior *types.CacheBehavior

	// The name for the distribution.
	//
	// This member is required.
	DistributionName *string

	// An object that describes the origin resource for the distribution, such as a
	// Lightsail instance or load balancer. The distribution pulls, caches, and serves
	// content from the origin.
	//
	// This member is required.
	Origin *types.InputOrigin

	// An object that describes the cache behavior settings for the distribution.
	CacheBehaviorSettings *types.CacheSettings

	// An array of objects that describe the per-path cache behavior for the
	// distribution.
	CacheBehaviors []types.CacheBehaviorPerPath

	// The IP address type for the distribution. The possible values are ipv4 for IPv4
	// only, and dualstack for IPv4 and IPv6. The default value is dualstack.
	IpAddressType types.IpAddressType

	// The tag keys and optional values to add to the distribution during create. Use
	// the TagResource action to tag a resource after it's created.
	Tags []types.Tag
}

type CreateDistributionOutput struct {

	// An object that describes the distribution created.
	Distribution *types.LightsailDistribution

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDistribution{}, middleware.After)
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
	if err = addOpCreateDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDistribution(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateDistribution",
	}
}
