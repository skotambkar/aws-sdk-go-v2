// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates stream metadata, such as the device name and media type. You must
// provide the stream name or the Amazon Resource Name (ARN) of the stream. To make
// sure that you have the latest version of the stream before updating it, you can
// specify the stream version. Kinesis Video Streams assigns a version to each
// stream. When you update a stream, Kinesis Video Streams assigns a new version
// number. To get the latest stream version, use the DescribeStream API.
// UpdateStream is an asynchronous operation, and takes time to complete.
func (c *Client) UpdateStream(ctx context.Context, params *UpdateStreamInput, optFns ...func(*Options)) (*UpdateStreamOutput, error) {
	if params == nil {
		params = &UpdateStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStream", params, optFns, addOperationUpdateStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStreamInput struct {

	// The version of the stream whose metadata you want to update.
	//
	// This member is required.
	CurrentVersion *string

	// The name of the device that is writing to the stream. In the current
	// implementation, Kinesis Video Streams does not use this name.
	DeviceName *string

	// The stream's media type. Use MediaType to specify the type of content that the
	// stream contains to the consumers of the stream. For more information about media
	// types, see Media Types
	// (http://www.iana.org/assignments/media-types/media-types.xhtml). If you choose
	// to specify the MediaType, see Naming Requirements
	// (https://tools.ietf.org/html/rfc6838#section-4.2). To play video on the console,
	// you must specify the correct video type. For example, if the video in the stream
	// is H.264, specify video/h264 as the MediaType.
	MediaType *string

	// The ARN of the stream whose metadata you want to update.
	StreamARN *string

	// The name of the stream whose metadata you want to update. The stream name is an
	// identifier for the stream, and must be unique for each account and region.
	StreamName *string
}

type UpdateStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStream{}, middleware.After)
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
	if err = addOpUpdateStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "UpdateStream",
	}
}
