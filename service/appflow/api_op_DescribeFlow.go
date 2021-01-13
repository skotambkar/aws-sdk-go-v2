// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides a description of the specified flow.
func (c *Client) DescribeFlow(ctx context.Context, params *DescribeFlowInput, optFns ...func(*Options)) (*DescribeFlowOutput, error) {
	if params == nil {
		params = &DescribeFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFlow", params, optFns, addOperationDescribeFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFlowInput struct {

	// The specified name of the flow. Spaces are not allowed. Use underscores (_) or
	// hyphens (-) only.
	//
	// This member is required.
	FlowName *string
}

type DescribeFlowOutput struct {

	// Specifies when the flow was created.
	CreatedAt *time.Time

	// The ARN of the user who created the flow.
	CreatedBy *string

	// A description of the flow.
	Description *string

	// The configuration that controls how Amazon AppFlow transfers data to the
	// destination connector.
	DestinationFlowConfigList []types.DestinationFlowConfig

	// The flow's Amazon Resource Name (ARN).
	FlowArn *string

	// The specified name of the flow. Spaces are not allowed. Use underscores (_) or
	// hyphens (-) only.
	FlowName *string

	// Indicates the current status of the flow.
	FlowStatus types.FlowStatus

	// Contains an error message if the flow status is in a suspended or error state.
	// This applies only to scheduled or event-triggered flows.
	FlowStatusMessage *string

	// The ARN (Amazon Resource Name) of the Key Management Service (KMS) key you
	// provide for encryption. This is required if you do not want to use the Amazon
	// AppFlow-managed KMS key. If you don't provide anything here, Amazon AppFlow uses
	// the Amazon AppFlow-managed KMS key.
	KmsArn *string

	// Describes the details of the most recent flow run.
	LastRunExecutionDetails *types.ExecutionDetails

	// Specifies when the flow was last updated.
	LastUpdatedAt *time.Time

	// Specifies the user name of the account that performed the most recent update.
	LastUpdatedBy *string

	// The configuration that controls how Amazon AppFlow retrieves data from the
	// source connector.
	SourceFlowConfig *types.SourceFlowConfig

	// The tags used to organize, track, or control access for your flow.
	Tags map[string]string

	// A list of tasks that Amazon AppFlow performs while transferring the data in the
	// flow run.
	Tasks []types.Task

	// The trigger settings that determine how and when the flow runs.
	TriggerConfig *types.TriggerConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFlow{}, middleware.After)
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
	if err = addOpDescribeFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "DescribeFlow",
	}
}
