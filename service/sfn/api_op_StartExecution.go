// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a state machine execution. StartExecution is idempotent. If
// StartExecution is called with the same name and input as a running execution,
// the call will succeed and return the same response as the original request. If
// the execution is closed or if the input is different, it will return a 400
// ExecutionAlreadyExists error. Names can be reused after 90 days.
func (c *Client) StartExecution(ctx context.Context, params *StartExecutionInput, optFns ...func(*Options)) (*StartExecutionOutput, error) {
	if params == nil {
		params = &StartExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartExecution", params, optFns, addOperationStartExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartExecutionInput struct {

	// The Amazon Resource Name (ARN) of the state machine to execute.
	//
	// This member is required.
	StateMachineArn *string

	// The string that contains the JSON input data for the execution, for example:
	// "input": "{\"first_name\" : \"test\"}" If you don't include any JSON input data,
	// you still must include the two braces, for example: "input": "{}" Length
	// constraints apply to the payload size, and are expressed as bytes in UTF-8
	// encoding.
	Input *string

	// The name of the execution. This name must be unique for your AWS account,
	// region, and state machine for 90 days. For more information, see  Limits Related
	// to State Machine Executions
	// (https://docs.aws.amazon.com/step-functions/latest/dg/limits.html#service-limits-state-machine-executions)
	// in the AWS Step Functions Developer Guide. A name must not contain:
	//
	// * white
	// space
	//
	// * brackets < > { } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters "
	// # % \ ^ | ~ ` $ & , ; : /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To
	// enable logging with CloudWatch Logs, the name should only contain 0-9, A-Z, a-z,
	// - and _.
	Name *string

	// Passes the AWS X-Ray trace header. The trace header can also be passed in the
	// request payload.
	TraceHeader *string
}

type StartExecutionOutput struct {

	// The Amazon Resource Name (ARN) that identifies the execution.
	//
	// This member is required.
	ExecutionArn *string

	// The date the execution is started.
	//
	// This member is required.
	StartDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStartExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartExecution{}, middleware.After)
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
	if err = addOpStartExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "StartExecution",
	}
}
