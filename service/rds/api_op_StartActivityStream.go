// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a database activity stream to monitor activity on the database. For more
// information, see Database Activity Streams
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/DBActivityStreams.html)
// in the Amazon Aurora User Guide.
func (c *Client) StartActivityStream(ctx context.Context, params *StartActivityStreamInput, optFns ...func(*Options)) (*StartActivityStreamOutput, error) {
	if params == nil {
		params = &StartActivityStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartActivityStream", params, optFns, c.addOperationStartActivityStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartActivityStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartActivityStreamInput struct {

	// The AWS KMS key identifier for encrypting messages in the database activity
	// stream. The AWS KMS key identifier is the key ARN, key ID, alias ARN, or alias
	// name for the AWS KMS customer master key (CMK).
	//
	// This member is required.
	KmsKeyId *string

	// Specifies the mode of the database activity stream. Database events such as a
	// change or access generate an activity stream event. The database session can
	// handle these events either synchronously or asynchronously.
	//
	// This member is required.
	Mode types.ActivityStreamMode

	// The Amazon Resource Name (ARN) of the DB cluster, for example,
	// arn:aws:rds:us-east-1:12345667890:cluster:das-cluster.
	//
	// This member is required.
	ResourceArn *string

	// Specifies whether or not the database activity stream is to start as soon as
	// possible, regardless of the maintenance window for the database.
	ApplyImmediately *bool
}

type StartActivityStreamOutput struct {

	// Indicates whether or not the database activity stream will start as soon as
	// possible, regardless of the maintenance window for the database.
	ApplyImmediately bool

	// The name of the Amazon Kinesis data stream to be used for the database activity
	// stream.
	KinesisStreamName *string

	// The AWS KMS key identifier for encryption of messages in the database activity
	// stream.
	KmsKeyId *string

	// The mode of the database activity stream.
	Mode types.ActivityStreamMode

	// The status of the database activity stream.
	Status types.ActivityStreamStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationStartActivityStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStartActivityStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStartActivityStream{}, middleware.After)
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
	if err = addOpStartActivityStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartActivityStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartActivityStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StartActivityStream",
	}
}
