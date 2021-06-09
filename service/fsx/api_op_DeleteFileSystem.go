// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a file system, deleting its contents. After deletion, the file system no
// longer exists, and its data is gone. Any existing automatic backups will also be
// deleted. By default, when you delete an Amazon FSx for Windows File Server file
// system, a final backup is created upon deletion. This final backup is not
// subject to the file system's retention policy, and must be manually deleted. The
// DeleteFileSystem action returns while the file system has the DELETING status.
// You can check the file system deletion status by calling the DescribeFileSystems
// action, which returns a list of file systems in your account. If you pass the
// file system ID for a deleted file system, the DescribeFileSystems returns a
// FileSystemNotFound error. Deleting an Amazon FSx for Lustre file system will
// fail with a 400 BadRequest if a data repository task is in a PENDING or
// EXECUTING state. The data in a deleted file system is also deleted and can't be
// recovered by any means.
func (c *Client) DeleteFileSystem(ctx context.Context, params *DeleteFileSystemInput, optFns ...func(*Options)) (*DeleteFileSystemOutput, error) {
	if params == nil {
		params = &DeleteFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFileSystem", params, optFns, c.addOperationDeleteFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for DeleteFileSystem operation.
type DeleteFileSystemInput struct {

	// The ID of the file system you want to delete.
	//
	// This member is required.
	FileSystemId *string

	// A string of up to 64 ASCII characters that Amazon FSx uses to ensure idempotent
	// deletion. This is automatically filled on your behalf when using the AWS CLI or
	// SDK.
	ClientRequestToken *string

	// The configuration object for the Amazon FSx for Lustre file system being deleted
	// in the DeleteFileSystem operation.
	LustreConfiguration *types.DeleteFileSystemLustreConfiguration

	// The configuration object for the Microsoft Windows file system used in the
	// DeleteFileSystem operation.
	WindowsConfiguration *types.DeleteFileSystemWindowsConfiguration
}

// The response object for the DeleteFileSystem operation.
type DeleteFileSystemOutput struct {

	// The ID of the file system being deleted.
	FileSystemId *string

	// The file system lifecycle for the deletion request. Should be DELETING.
	Lifecycle types.FileSystemLifecycle

	// The response object for the Amazon FSx for Lustre file system being deleted in
	// the DeleteFileSystem operation.
	LustreResponse *types.DeleteFileSystemLustreResponse

	// The response object for the Microsoft Windows file system used in the
	// DeleteFileSystem operation.
	WindowsResponse *types.DeleteFileSystemWindowsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteFileSystem{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteFileSystemMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFileSystem(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteFileSystem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteFileSystem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteFileSystemInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteFileSystemMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteFileSystem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DeleteFileSystem",
	}
}
