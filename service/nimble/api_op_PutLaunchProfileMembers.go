// Code generated by smithy-go-codegen DO NOT EDIT.

package nimble

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/nimble/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Add/update users with given persona to launch profile membership.
func (c *Client) PutLaunchProfileMembers(ctx context.Context, params *PutLaunchProfileMembersInput, optFns ...func(*Options)) (*PutLaunchProfileMembersOutput, error) {
	if params == nil {
		params = &PutLaunchProfileMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLaunchProfileMembers", params, optFns, c.addOperationPutLaunchProfileMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLaunchProfileMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Information about a launch profile membership.
type PutLaunchProfileMembersInput struct {

	// The ID of the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// The launch profile ID.
	//
	// This member is required.
	LaunchProfileId *string

	// A list of members.
	//
	// This member is required.
	Members []types.NewLaunchProfileMember

	// The studio ID.
	//
	// This member is required.
	StudioId *string

	// To make an idempotent API request using one of these actions, specify a client
	// token in the request. You should not reuse the same client token for other API
	// requests. If you retry a request that completed successfully using the same
	// client token and the same parameters, the retry succeeds without performing any
	// further actions. If you retry a successful request using the same client token,
	// but one or more of the parameters are different, the retry fails with a
	// ValidationException error.
	ClientToken *string
}

type PutLaunchProfileMembersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationPutLaunchProfileMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutLaunchProfileMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutLaunchProfileMembers{}, middleware.After)
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
	if err = addIdempotencyToken_opPutLaunchProfileMembersMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutLaunchProfileMembersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLaunchProfileMembers(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPutLaunchProfileMembers struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPutLaunchProfileMembers) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPutLaunchProfileMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PutLaunchProfileMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PutLaunchProfileMembersInput ")
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
func addIdempotencyToken_opPutLaunchProfileMembersMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPutLaunchProfileMembers{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPutLaunchProfileMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "PutLaunchProfileMembers",
	}
}
