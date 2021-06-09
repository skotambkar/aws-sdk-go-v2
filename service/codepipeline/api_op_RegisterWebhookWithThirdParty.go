// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures a connection between the webhook that was created and the external
// tool with events to be detected.
func (c *Client) RegisterWebhookWithThirdParty(ctx context.Context, params *RegisterWebhookWithThirdPartyInput, optFns ...func(*Options)) (*RegisterWebhookWithThirdPartyOutput, error) {
	if params == nil {
		params = &RegisterWebhookWithThirdPartyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterWebhookWithThirdParty", params, optFns, c.addOperationRegisterWebhookWithThirdPartyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterWebhookWithThirdPartyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterWebhookWithThirdPartyInput struct {

	// The name of an existing webhook created with PutWebhook to register with a
	// supported third party.
	WebhookName *string
}

type RegisterWebhookWithThirdPartyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationRegisterWebhookWithThirdPartyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterWebhookWithThirdParty{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterWebhookWithThirdParty{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterWebhookWithThirdParty(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterWebhookWithThirdParty(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "RegisterWebhookWithThirdParty",
	}
}
