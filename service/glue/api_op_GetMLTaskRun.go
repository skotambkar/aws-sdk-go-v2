// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets details for a specific task run on a machine learning transform. Machine
// learning task runs are asynchronous tasks that AWS Glue runs on your behalf as
// part of various machine learning workflows. You can check the stats of any task
// run by calling GetMLTaskRun with the TaskRunID and its parent transform's
// TransformID.
func (c *Client) GetMLTaskRun(ctx context.Context, params *GetMLTaskRunInput, optFns ...func(*Options)) (*GetMLTaskRunOutput, error) {
	if params == nil {
		params = &GetMLTaskRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMLTaskRun", params, optFns, addOperationGetMLTaskRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMLTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMLTaskRunInput struct {

	// The unique identifier of the task run.
	//
	// This member is required.
	TaskRunId *string

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string
}

type GetMLTaskRunOutput struct {

	// The date and time when this task run was completed.
	CompletedOn *time.Time

	// The error strings that are associated with the task run.
	ErrorString *string

	// The amount of time (in seconds) that the task run consumed resources.
	ExecutionTime int32

	// The date and time when this task run was last modified.
	LastModifiedOn *time.Time

	// The names of the log groups that are associated with the task run.
	LogGroupName *string

	// The list of properties that are associated with the task run.
	Properties *types.TaskRunProperties

	// The date and time when this task run started.
	StartedOn *time.Time

	// The status for this task run.
	Status types.TaskStatusType

	// The unique run identifier associated with this run.
	TaskRunId *string

	// The unique identifier of the task run.
	TransformId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMLTaskRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMLTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMLTaskRun{}, middleware.After)
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
	if err = addOpGetMLTaskRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMLTaskRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMLTaskRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetMLTaskRun",
	}
}
