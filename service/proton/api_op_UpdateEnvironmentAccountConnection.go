// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// In an environment account, update an environment account connection to use a new
// IAM role. For more information, see Environment account connections in the AWS
// Proton Administration guide.
func (c *Client) UpdateEnvironmentAccountConnection(ctx context.Context, params *UpdateEnvironmentAccountConnectionInput, optFns ...func(*Options)) (*UpdateEnvironmentAccountConnectionOutput, error) {
	if params == nil {
		params = &UpdateEnvironmentAccountConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEnvironmentAccountConnection", params, optFns, c.addOperationUpdateEnvironmentAccountConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEnvironmentAccountConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEnvironmentAccountConnectionInput struct {

	// The ID of the environment account connection to update.
	//
	// This member is required.
	Id *string

	// The Amazon Resource Name (ARN) of the IAM service role that is associated with
	// the environment account connection to update.
	//
	// This member is required.
	RoleArn *string
}

type UpdateEnvironmentAccountConnectionOutput struct {

	// The environment account connection detail data that's returned by AWS Proton.
	//
	// This member is required.
	EnvironmentAccountConnection *types.EnvironmentAccountConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationUpdateEnvironmentAccountConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateEnvironmentAccountConnection{}, middleware.After)
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
	if err = addOpUpdateEnvironmentAccountConnectionValidationMiddleware(stack); err != nil {
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
