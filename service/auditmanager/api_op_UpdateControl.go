// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a custom control in AWS Audit Manager.
func (c *Client) UpdateControl(ctx context.Context, params *UpdateControlInput, optFns ...func(*Options)) (*UpdateControlOutput, error) {
	if params == nil {
		params = &UpdateControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateControl", params, optFns, c.addOperationUpdateControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateControlInput struct {

	// The identifier for the specified control.
	//
	// This member is required.
	ControlId *string

	// The data mapping sources for the specified control.
	//
	// This member is required.
	ControlMappingSources []types.ControlMappingSource

	// The name of the control to be updated.
	//
	// This member is required.
	Name *string

	// The recommended actions to carry out if the control is not fulfilled.
	ActionPlanInstructions *string

	// The title of the action plan for remediating the control.
	ActionPlanTitle *string

	// The optional description of the control.
	Description *string

	// The steps that to follow to determine if the control has been satisfied.
	TestingInformation *string
}

type UpdateControlOutput struct {

	// The name of the updated control set returned by the UpdateControl API.
	Control *types.Control

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateControl{}, middleware.After)
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
	if err = addOpUpdateControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "UpdateControl",
	}
}
