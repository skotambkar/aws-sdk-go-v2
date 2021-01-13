// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Designates an account as the delegated Amazon Macie administrator account for an
// AWS organization.
func (c *Client) EnableOrganizationAdminAccount(ctx context.Context, params *EnableOrganizationAdminAccountInput, optFns ...func(*Options)) (*EnableOrganizationAdminAccountOutput, error) {
	if params == nil {
		params = &EnableOrganizationAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableOrganizationAdminAccount", params, optFns, addOperationEnableOrganizationAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableOrganizationAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableOrganizationAdminAccountInput struct {

	// The AWS account ID for the account to designate as the delegated Amazon Macie
	// administrator account for the organization.
	//
	// This member is required.
	AdminAccountId *string

	// A unique, case-sensitive token that you provide to ensure the idempotency of the
	// request.
	ClientToken *string
}

type EnableOrganizationAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableOrganizationAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEnableOrganizationAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableOrganizationAdminAccount{}, middleware.After)
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
	if err = addIdempotencyToken_opEnableOrganizationAdminAccountMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpEnableOrganizationAdminAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableOrganizationAdminAccount(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpEnableOrganizationAdminAccount struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpEnableOrganizationAdminAccount) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpEnableOrganizationAdminAccount) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*EnableOrganizationAdminAccountInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *EnableOrganizationAdminAccountInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opEnableOrganizationAdminAccountMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpEnableOrganizationAdminAccount{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opEnableOrganizationAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "EnableOrganizationAdminAccount",
	}
}
