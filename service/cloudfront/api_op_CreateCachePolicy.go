// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a cache policy. After you create a cache policy, you can attach it to
// one or more cache behaviors. When it’s attached to a cache behavior, the cache
// policy determines the following:
//
// * The values that CloudFront includes in the
// cache key. These values can include HTTP headers, cookies, and URL query
// strings. CloudFront uses the cache key to find an object in its cache that it
// can return to the viewer.
//
// * The default, minimum, and maximum time to live
// (TTL) values that you want objects to stay in the CloudFront cache.
//
// The
// headers, cookies, and query strings that are included in the cache key are
// automatically included in requests that CloudFront sends to the origin.
// CloudFront sends a request when it can’t find an object in its cache that
// matches the request’s cache key. If you want to send values to the origin but
// not include them in the cache key, use OriginRequestPolicy. For more information
// about cache policies, see Controlling the cache key
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateCachePolicy(ctx context.Context, params *CreateCachePolicyInput, optFns ...func(*Options)) (*CreateCachePolicyOutput, error) {
	if params == nil {
		params = &CreateCachePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCachePolicy", params, optFns, addOperationCreateCachePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCachePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCachePolicyInput struct {

	// A cache policy configuration.
	//
	// This member is required.
	CachePolicyConfig *types.CachePolicyConfig
}

type CreateCachePolicyOutput struct {

	// A cache policy.
	CachePolicy *types.CachePolicy

	// The current version of the cache policy.
	ETag *string

	// The fully qualified URI of the cache policy just created.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCachePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateCachePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateCachePolicy{}, middleware.After)
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
	if err = addOpCreateCachePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCachePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCachePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateCachePolicy",
	}
}
