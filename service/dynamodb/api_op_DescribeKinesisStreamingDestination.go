// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the status of Kinesis streaming.
func (c *Client) DescribeKinesisStreamingDestination(ctx context.Context, params *DescribeKinesisStreamingDestinationInput, optFns ...func(*Options)) (*DescribeKinesisStreamingDestinationOutput, error) {
	if params == nil {
		params = &DescribeKinesisStreamingDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeKinesisStreamingDestination", params, optFns, c.addOperationDescribeKinesisStreamingDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeKinesisStreamingDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeKinesisStreamingDestinationInput struct {

	// The name of the table being described.
	//
	// This member is required.
	TableName *string
}

type DescribeKinesisStreamingDestinationOutput struct {

	// The list of replica structures for the table being described.
	KinesisDataStreamDestinations []types.KinesisDataStreamDestination

	// The name of the table being described.
	TableName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeKinesisStreamingDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeKinesisStreamingDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeKinesisStreamingDestination{}, middleware.After)
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
	if err = addOpDescribeKinesisStreamingDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeKinesisStreamingDestination(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeKinesisStreamingDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "DescribeKinesisStreamingDestination",
	}
}
