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

// Creates a streaming image resource in a studio.
func (c *Client) CreateStreamingImage(ctx context.Context, params *CreateStreamingImageInput, optFns ...func(*Options)) (*CreateStreamingImageOutput, error) {
	if params == nil {
		params = &CreateStreamingImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStreamingImage", params, optFns, c.addOperationCreateStreamingImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamingImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A collection of streaming images.
type CreateStreamingImageInput struct {

	// The ID of an EC2 machine image with which to create this streaming image.
	//
	// This member is required.
	Ec2ImageId *string

	// A friendly name for a streaming image resource.
	//
	// This member is required.
	Name *string

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

	// A human-readable description of the streaming image.
	Description *string

	// A collection of labels, in the form of key:value pairs, that apply to this
	// resource.
	Tags map[string]string
}

type CreateStreamingImageOutput struct {
	StreamingImage *types.StreamingImage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateStreamingImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateStreamingImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateStreamingImage{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateStreamingImageMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateStreamingImageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamingImage(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateStreamingImage struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateStreamingImage) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateStreamingImage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateStreamingImageInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateStreamingImageInput ")
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
func addIdempotencyToken_opCreateStreamingImageMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateStreamingImage{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateStreamingImage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "nimble",
		OperationName: "CreateStreamingImage",
	}
}
