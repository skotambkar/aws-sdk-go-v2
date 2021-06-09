// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new ledger in your AWS account in the current Region.
func (c *Client) CreateLedger(ctx context.Context, params *CreateLedgerInput, optFns ...func(*Options)) (*CreateLedgerOutput, error) {
	if params == nil {
		params = &CreateLedgerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLedger", params, optFns, c.addOperationCreateLedgerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLedgerInput struct {

	// The name of the ledger that you want to create. The name must be unique among
	// all of your ledgers in the current AWS Region. Naming constraints for ledger
	// names are defined in Quotas in Amazon QLDB
	// (https://docs.aws.amazon.com/qldb/latest/developerguide/limits.html#limits.naming)
	// in the Amazon QLDB Developer Guide.
	//
	// This member is required.
	Name *string

	// The permissions mode to assign to the ledger that you want to create. This
	// parameter can have one of the following values:
	//
	// * ALLOW_ALL: A legacy
	// permissions mode that enables access control with API-level granularity for
	// ledgers. This mode allows users who have the SendCommand API permission for this
	// ledger to run all PartiQL commands (hence, ALLOW_ALL) on any tables in the
	// specified ledger. This mode disregards any table-level or command-level IAM
	// permissions policies that you create for the ledger.
	//
	// * STANDARD: (Recommended)
	// A permissions mode that enables access control with finer granularity for
	// ledgers, tables, and PartiQL commands. By default, this mode denies all user
	// requests to run any PartiQL commands on any tables in this ledger. To allow
	// PartiQL commands to run, you must create IAM permissions policies for specific
	// table resources and PartiQL actions, in addition to the SendCommand API
	// permission for the ledger. For information, see Getting started with the
	// standard permissions mode
	// (https://docs.aws.amazon.com/qldb/latest/developerguide/getting-started-standard-mode.html)
	// in the Amazon QLDB Developer Guide.
	//
	// We strongly recommend using the STANDARD
	// permissions mode to maximize the security of your ledger data.
	//
	// This member is required.
	PermissionsMode types.PermissionsMode

	// The flag that prevents a ledger from being deleted by any user. If not provided
	// on ledger creation, this feature is enabled (true) by default. If deletion
	// protection is enabled, you must first disable it before you can delete the
	// ledger. You can disable it by calling the UpdateLedger operation to set the flag
	// to false.
	DeletionProtection *bool

	// The key-value pairs to add as tags to the ledger that you want to create. Tag
	// keys are case sensitive. Tag values are case sensitive and can be null.
	Tags map[string]*string
}

type CreateLedgerOutput struct {

	// The Amazon Resource Name (ARN) for the ledger.
	Arn *string

	// The date and time, in epoch time format, when the ledger was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1, 1970
	// UTC.)
	CreationDateTime *time.Time

	// The flag that prevents a ledger from being deleted by any user. If not provided
	// on ledger creation, this feature is enabled (true) by default. If deletion
	// protection is enabled, you must first disable it before you can delete the
	// ledger. You can disable it by calling the UpdateLedger operation to set the flag
	// to false.
	DeletionProtection *bool

	// The name of the ledger.
	Name *string

	// The permissions mode of the ledger that you created.
	PermissionsMode types.PermissionsMode

	// The current status of the ledger.
	State types.LedgerState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLedger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLedger{}, middleware.After)
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
	if err = addOpCreateLedgerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLedger(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "CreateLedger",
	}
}
