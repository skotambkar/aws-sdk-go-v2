// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a specified playback authorization key pair and returns the arn and
// fingerprint. The privateKey held by the caller can be used to generate viewer
// authorization tokens, to grant viewers access to private channels. For more
// information, see Setting Up Private Channels
// (https://docs.aws.amazon.com/ivs/latest/userguide/private-channels.html) in the
// Amazon IVS User Guide.
func (c *Client) GetPlaybackKeyPair(ctx context.Context, params *GetPlaybackKeyPairInput, optFns ...func(*Options)) (*GetPlaybackKeyPairOutput, error) {
	if params == nil {
		params = &GetPlaybackKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPlaybackKeyPair", params, optFns, c.addOperationGetPlaybackKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPlaybackKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPlaybackKeyPairInput struct {

	// ARN of the key pair to be returned.
	//
	// This member is required.
	Arn *string
}

type GetPlaybackKeyPairOutput struct {

	// A key pair used to sign and validate a playback authorization token.
	KeyPair *types.PlaybackKeyPair

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetPlaybackKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPlaybackKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPlaybackKeyPair{}, middleware.After)
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
	if err = addOpGetPlaybackKeyPairValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPlaybackKeyPair(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPlaybackKeyPair(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "GetPlaybackKeyPair",
	}
}
