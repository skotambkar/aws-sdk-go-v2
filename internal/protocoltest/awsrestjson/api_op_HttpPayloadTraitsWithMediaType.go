// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This examples uses a @mediaType trait on the payload to force a custom
// content-type to be serialized.
func (c *Client) HttpPayloadTraitsWithMediaType(ctx context.Context, params *HttpPayloadTraitsWithMediaTypeInput, optFns ...func(*Options)) (*HttpPayloadTraitsWithMediaTypeOutput, error) {
	if params == nil {
		params = &HttpPayloadTraitsWithMediaTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpPayloadTraitsWithMediaType", params, optFns, c.addOperationHttpPayloadTraitsWithMediaTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpPayloadTraitsWithMediaTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadTraitsWithMediaTypeInput struct {

	// This value conforms to the media type: text/plain
	Blob []byte

	Foo *string
}

type HttpPayloadTraitsWithMediaTypeOutput struct {

	// This value conforms to the media type: text/plain
	Blob []byte

	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationHttpPayloadTraitsWithMediaTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpHttpPayloadTraitsWithMediaType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpPayloadTraitsWithMediaType{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadTraitsWithMediaType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opHttpPayloadTraitsWithMediaType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadTraitsWithMediaType",
	}
}
