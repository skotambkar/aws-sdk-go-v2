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

// Publishes a CloudFront function by copying the function code from the
// DEVELOPMENT stage to LIVE. This automatically updates all cache behaviors that
// are using this function to use the newly published copy in the LIVE stage. When
// a function is published to the LIVE stage, you can attach the function to a
// distribution’s cache behavior, using the function’s Amazon Resource Name (ARN).
// To publish a function, you must provide the function’s name and version (ETag
// value). To get these values, you can use ListFunctions and DescribeFunction.
func (c *Client) PublishFunction(ctx context.Context, params *PublishFunctionInput, optFns ...func(*Options)) (*PublishFunctionOutput, error) {
	if params == nil {
		params = &PublishFunctionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishFunction", params, optFns, c.addOperationPublishFunctionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishFunctionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishFunctionInput struct {

	// The current version (ETag value) of the function that you are publishing, which
	// you can get using DescribeFunction.
	//
	// This member is required.
	IfMatch *string

	// The name of the function that you are publishing.
	//
	// This member is required.
	Name *string
}

type PublishFunctionOutput struct {

	// Contains configuration information and metadata about a CloudFront function.
	FunctionSummary *types.FunctionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPublishFunctionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPublishFunction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPublishFunction{}, middleware.After)
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
	if err = addOpPublishFunctionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublishFunction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPublishFunction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "PublishFunction",
	}
}
