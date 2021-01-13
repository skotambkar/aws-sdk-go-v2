// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a state machine. A state machine consists of a collection of states that
// can do work (Task states), determine to which states to transition next (Choice
// states), stop an execution with an error (Fail states), and so on. State
// machines are specified using a JSON-based, structured language. For more
// information, see Amazon States Language
// (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html)
// in the AWS Step Functions User Guide. This operation is eventually consistent.
// The results are best effort and may not reflect very recent updates and changes.
// CreateStateMachine is an idempotent API. Subsequent requests won’t create a
// duplicate resource if it was already created. CreateStateMachine's idempotency
// check is based on the state machine name, definition, type, LoggingConfiguration
// and TracingConfiguration. If a following request has a different roleArn or
// tags, Step Functions will ignore these differences and treat it as an idempotent
// request of the previous. In this case, roleArn and tags will not be updated,
// even if they are different.
func (c *Client) CreateStateMachine(ctx context.Context, params *CreateStateMachineInput, optFns ...func(*Options)) (*CreateStateMachineOutput, error) {
	if params == nil {
		params = &CreateStateMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStateMachine", params, optFns, addOperationCreateStateMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStateMachineInput struct {

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language
	// (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	//
	// This member is required.
	Definition *string

	// The name of the state machine. A name must not contain:
	//
	// * white space
	//
	// *
	// brackets < > { } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters " # % \ ^
	// | ~ ` $ & , ; : /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable
	// logging with CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and
	// _.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
	//
	// This member is required.
	RoleArn *string

	// Defines what execution history events are logged and where they are logged. By
	// default, the level is set to OFF. For more information see Log Levels
	// (https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html)
	// in the AWS Step Functions User Guide.
	LoggingConfiguration *types.LoggingConfiguration

	// Tags to be added when creating a state machine. An array of key-value pairs. For
	// more information, see Using Cost Allocation Tags
	// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
	// in the AWS Billing and Cost Management User Guide, and Controlling Access Using
	// IAM Tags
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_iam-tags.html). Tags
	// may only contain Unicode letters, digits, white space, or these symbols: _ . : /
	// = + - @.
	Tags []types.Tag

	// Selects whether AWS X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration

	// Determines whether a Standard or Express state machine is created. The default
	// is STANDARD. You cannot update the type of a state machine once it has been
	// created.
	Type types.StateMachineType
}

type CreateStateMachineOutput struct {

	// The date the state machine is created.
	//
	// This member is required.
	CreationDate *time.Time

	// The Amazon Resource Name (ARN) that identifies the created state machine.
	//
	// This member is required.
	StateMachineArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateStateMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateStateMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateStateMachine{}, middleware.After)
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
	if err = addOpCreateStateMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStateMachine(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStateMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "CreateStateMachine",
	}
}
