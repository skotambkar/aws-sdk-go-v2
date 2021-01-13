// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns details about an event source mapping. You can get the identifier of a
// mapping from the output of ListEventSourceMappings.
func (c *Client) GetEventSourceMapping(ctx context.Context, params *GetEventSourceMappingInput, optFns ...func(*Options)) (*GetEventSourceMappingOutput, error) {
	if params == nil {
		params = &GetEventSourceMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventSourceMapping", params, optFns, addOperationGetEventSourceMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventSourceMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventSourceMappingInput struct {

	// The identifier of the event source mapping.
	//
	// This member is required.
	UUID *string
}

// A mapping between an AWS resource and an AWS Lambda function. See
// CreateEventSourceMapping for details.
type GetEventSourceMappingOutput struct {

	// The maximum number of items to retrieve in a single batch.
	BatchSize *int32

	// (Streams) If the function returns an error, split the batch in two and retry.
	// The default value is false.
	BisectBatchOnFunctionError *bool

	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded
	// records.
	DestinationConfig *types.DestinationConfig

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string

	// The ARN of the Lambda function.
	FunctionArn *string

	// (Streams) A list of current response type enums applied to the event source
	// mapping.
	FunctionResponseTypes []types.FunctionResponseType

	// The date that the event source mapping was last updated, or its state changed.
	LastModified *time.Time

	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult *string

	// (Streams and SQS standard queues) The maximum amount of time to gather records
	// before invoking the function, in seconds. The default value is zero.
	MaximumBatchingWindowInSeconds *int32

	// (Streams) Discard records older than the specified age. The default value is
	// infinite (-1). When set to infinite (-1), failed records are retried until the
	// record expires.
	MaximumRecordAgeInSeconds *int32

	// (Streams) Discard records after the specified number of retries. The default
	// value is infinite (-1). When set to infinite (-1), failed records are retried
	// until the record expires.
	MaximumRetryAttempts *int32

	// (Streams) The number of batches to process from each shard concurrently. The
	// default value is 1.
	ParallelizationFactor *int32

	// (MQ) The name of the Amazon MQ broker destination queue to consume.
	Queues []string

	// The Self-Managed Apache Kafka cluster for your event source.
	SelfManagedEventSource *types.SelfManagedEventSource

	// An array of the authentication protocol, or the VPC components to secure your
	// event source.
	SourceAccessConfigurations []types.SourceAccessConfiguration

	// The position in a stream from which to start reading. Required for Amazon
	// Kinesis, Amazon DynamoDB, and Amazon MSK Streams sources. AT_TIMESTAMP is only
	// supported for Amazon Kinesis streams.
	StartingPosition types.EventSourcePosition

	// With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.
	StartingPositionTimestamp *time.Time

	// The state of the event source mapping. It can be one of the following: Creating,
	// Enabling, Enabled, Disabling, Disabled, Updating, or Deleting.
	State *string

	// Indicates whether the last change to the event source mapping was made by a
	// user, or by the Lambda service.
	StateTransitionReason *string

	// The name of the Kafka topic.
	Topics []string

	// (Streams) The duration of a processing window in seconds. The range is between 1
	// second up to 15 minutes.
	TumblingWindowInSeconds *int32

	// The identifier of the event source mapping.
	UUID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEventSourceMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEventSourceMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEventSourceMapping{}, middleware.After)
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
	if err = addOpGetEventSourceMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventSourceMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventSourceMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetEventSourceMapping",
	}
}
