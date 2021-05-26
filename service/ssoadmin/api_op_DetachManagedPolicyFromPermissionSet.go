// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detaches the attached IAM managed policy ARN from the specified permission set.
func (c *Client) DetachManagedPolicyFromPermissionSet(ctx context.Context, params *DetachManagedPolicyFromPermissionSetInput, optFns ...func(*Options)) (*DetachManagedPolicyFromPermissionSetOutput, error) {
	if params == nil {
		params = &DetachManagedPolicyFromPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachManagedPolicyFromPermissionSet", params, optFns, c.addOperationDetachManagedPolicyFromPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachManagedPolicyFromPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachManagedPolicyFromPermissionSetInput struct {

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The IAM managed policy ARN to be attached to a permission set.
	//
	// This member is required.
	ManagedPolicyArn *string

	// The ARN of the PermissionSet from which the policy should be detached.
	//
	// This member is required.
	PermissionSetArn *string
}

type DetachManagedPolicyFromPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDetachManagedPolicyFromPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetachManagedPolicyFromPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachManagedPolicyFromPermissionSet{}, middleware.After)
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
	if err = addOpDetachManagedPolicyFromPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachManagedPolicyFromPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachManagedPolicyFromPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "DetachManagedPolicyFromPermissionSet",
	}
}
