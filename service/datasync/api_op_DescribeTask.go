// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata about a task.
func (c *Client) DescribeTask(ctx context.Context, params *DescribeTaskInput, optFns ...func(*Options)) (*DescribeTaskOutput, error) {
	if params == nil {
		params = &DescribeTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTask", params, optFns, c.addOperationDescribeTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTaskRequest
type DescribeTaskInput struct {

	// The Amazon Resource Name (ARN) of the task to describe.
	//
	// This member is required.
	TaskArn *string
}

// DescribeTaskResponse
type DescribeTaskOutput struct {

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that was used
	// to monitor and log events in the task. For more information on these groups, see
	// Working with Log Groups and Log Streams in the Amazon CloudWatch User Guide.
	CloudWatchLogGroupArn *string

	// The time that the task was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the task execution that is syncing files.
	CurrentTaskExecutionArn *string

	// The Amazon Resource Name (ARN) of the AWS storage resource's location.
	DestinationLocationArn *string

	// The Amazon Resource Name (ARN) of the destination ENIs (Elastic Network
	// Interface) that was created for your subnet.
	DestinationNetworkInterfaceArns []string

	// Errors that AWS DataSync encountered during execution of the task. You can use
	// this error code to help troubleshoot issues.
	ErrorCode *string

	// Detailed description of an error that was encountered during the task execution.
	// You can use this information to help troubleshoot issues.
	ErrorDetail *string

	// A list of filter rules that determines which files to exclude from a task. The
	// list should contain a single filter string that consists of the patterns to
	// exclude. The patterns are delimited by "|" (that is, a pipe), for example:
	// "/folder1|/folder2"
	Excludes []types.FilterRule

	// The name of the task that was described.
	Name *string

	// The set of configuration options that control the behavior of a single execution
	// of the task that occurs when you call StartTaskExecution. You can configure
	// these options to preserve metadata such as user ID (UID) and group (GID), file
	// permissions, data integrity verification, and so on. For each individual task
	// execution, you can override these options by specifying the overriding
	// OverrideOptions value to StartTaskExecution
	// (https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html)
	// operation.
	Options *types.Options

	// The schedule used to periodically transfer files from a source to a destination
	// location.
	Schedule *types.TaskSchedule

	// The Amazon Resource Name (ARN) of the source file system's location.
	SourceLocationArn *string

	// The Amazon Resource Name (ARN) of the source ENIs (Elastic Network Interface)
	// that was created for your subnet.
	SourceNetworkInterfaceArns []string

	// The status of the task that was described. For detailed information about task
	// execution statuses, see Understanding Task Statuses in the AWS DataSync User
	// Guide.
	Status types.TaskStatus

	// The Amazon Resource Name (ARN) of the task that was described.
	TaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTask{}, middleware.After)
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
	if err = addOpDescribeTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeTask",
	}
}
