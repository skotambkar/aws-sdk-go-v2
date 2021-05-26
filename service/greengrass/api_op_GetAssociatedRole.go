// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the role associated with a particular group.
func (c *Client) GetAssociatedRole(ctx context.Context, params *GetAssociatedRoleInput, optFns ...func(*Options)) (*GetAssociatedRoleOutput, error) {
	if params == nil {
		params = &GetAssociatedRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAssociatedRole", params, optFns, c.addOperationGetAssociatedRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAssociatedRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssociatedRoleInput struct {

	// The ID of the Greengrass group.
	//
	// This member is required.
	GroupId *string
}

type GetAssociatedRoleOutput struct {

	// The time when the role was associated with the group.
	AssociatedAt *string

	// The ARN of the role that is associated with the group.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetAssociatedRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAssociatedRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAssociatedRole{}, middleware.After)
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
	if err = addOpGetAssociatedRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssociatedRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAssociatedRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetAssociatedRole",
	}
}
