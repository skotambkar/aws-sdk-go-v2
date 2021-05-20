// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an AWS App Runner connection resource. App Runner requires a connection
// resource when you create App Runner services that access private repositories
// from certain third-party providers. You can share a connection across multiple
// services. A connection resource is needed to access GitHub repositories. GitHub
// requires a user interface approval process through the App Runner console before
// you can use the connection.
func (c *Client) CreateConnection(ctx context.Context, params *CreateConnectionInput, optFns ...func(*Options)) (*CreateConnectionOutput, error) {
	if params == nil {
		params = &CreateConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConnection", params, optFns, addOperationCreateConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectionInput struct {

	// A name for the new connection. It must be unique across all App Runner
	// connections for the AWS account in the AWS Region.
	//
	// This member is required.
	ConnectionName *string

	// The source repository provider.
	//
	// This member is required.
	ProviderType types.ProviderType

	// A list of metadata items that you can associate with your connection resource. A
	// tag is a key-value pair.
	Tags []types.Tag
}

type CreateConnectionOutput struct {

	// A description of the App Runner connection that's created by this request.
	//
	// This member is required.
	Connection *types.Connection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateConnection{}, middleware.After)
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
	if err = addOpCreateConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "CreateConnection",
	}
}
