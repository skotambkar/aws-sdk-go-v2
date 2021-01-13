// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates settings for a server. This operation is synchronous.
func (c *Client) UpdateServer(ctx context.Context, params *UpdateServerInput, optFns ...func(*Options)) (*UpdateServerOutput, error) {
	if params == nil {
		params = &UpdateServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServer", params, optFns, addOperationUpdateServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServerInput struct {

	// The name of the server to update.
	//
	// This member is required.
	ServerName *string

	// Sets the number of automated backups that you want to keep.
	BackupRetentionCount *int32

	// Setting DisableAutomatedBackup to true disables automated or scheduled backups.
	// Automated backups are enabled by default.
	DisableAutomatedBackup *bool

	// DDD:HH:MM (weekly start time) or HH:MM (daily start time). Time windows always
	// use coordinated universal time (UTC). Valid strings for day of week (DDD) are:
	// Mon, Tue, Wed, Thr, Fri, Sat, or Sun.
	PreferredBackupWindow *string

	// DDD:HH:MM (weekly start time) or HH:MM (daily start time). Time windows always
	// use coordinated universal time (UTC). Valid strings for day of week (DDD) are:
	// Mon, Tue, Wed, Thr, Fri, Sat, or Sun.
	PreferredMaintenanceWindow *string
}

type UpdateServerOutput struct {

	// Contains the response to a UpdateServer request.
	Server *types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServer{}, middleware.After)
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
	if err = addOpUpdateServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "UpdateServer",
	}
}
