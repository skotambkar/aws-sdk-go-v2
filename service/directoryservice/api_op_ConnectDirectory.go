// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an AD Connector to connect to an on-premises directory. Before you call
// ConnectDirectory, ensure that all of the required permissions have been
// explicitly granted through a policy. For details about what permissions are
// required to run the ConnectDirectory operation, see AWS Directory Service API
// Permissions: Actions, Resources, and Conditions Reference
// (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
func (c *Client) ConnectDirectory(ctx context.Context, params *ConnectDirectoryInput, optFns ...func(*Options)) (*ConnectDirectoryOutput, error) {
	if params == nil {
		params = &ConnectDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConnectDirectory", params, optFns, c.addOperationConnectDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConnectDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the ConnectDirectory operation.
type ConnectDirectoryInput struct {

	// A DirectoryConnectSettings object that contains additional information for the
	// operation.
	//
	// This member is required.
	ConnectSettings *types.DirectoryConnectSettings

	// The fully qualified name of the on-premises directory, such as corp.example.com.
	//
	// This member is required.
	Name *string

	// The password for the on-premises user account.
	//
	// This member is required.
	Password *string

	// The size of the directory.
	//
	// This member is required.
	Size types.DirectorySize

	// A description for the directory.
	Description *string

	// The NetBIOS name of the on-premises directory, such as CORP.
	ShortName *string

	// The tags to be assigned to AD Connector.
	Tags []types.Tag
}

// Contains the results of the ConnectDirectory operation.
type ConnectDirectoryOutput struct {

	// The identifier of the new directory.
	DirectoryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationConnectDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpConnectDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpConnectDirectory{}, middleware.After)
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
	if err = addOpConnectDirectoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConnectDirectory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConnectDirectory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ConnectDirectory",
	}
}
