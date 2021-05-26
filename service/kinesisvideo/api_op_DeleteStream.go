// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Kinesis video stream and the data contained in the stream. This method
// marks the stream for deletion, and makes the data in the stream inaccessible
// immediately. To ensure that you have the latest version of the stream before
// deleting it, you can specify the stream version. Kinesis Video Streams assigns a
// version to each stream. When you update a stream, Kinesis Video Streams assigns
// a new version number. To get the latest stream version, use the DescribeStream
// API. This operation requires permission for the KinesisVideo:DeleteStream
// action.
func (c *Client) DeleteStream(ctx context.Context, params *DeleteStreamInput, optFns ...func(*Options)) (*DeleteStreamOutput, error) {
	if params == nil {
		params = &DeleteStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStream", params, optFns, c.addOperationDeleteStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteStreamInput struct {

	// The Amazon Resource Name (ARN) of the stream that you want to delete.
	//
	// This member is required.
	StreamARN *string

	// Optional: The version of the stream that you want to delete. Specify the version
	// as a safeguard to ensure that your are deleting the correct stream. To get the
	// stream version, use the DescribeStream API. If not specified, only the
	// CreationTime is checked before deleting the stream.
	CurrentVersion *string
}

type DeleteStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteStream{}, middleware.After)
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
	if err = addOpDeleteStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "DeleteStream",
	}
}
