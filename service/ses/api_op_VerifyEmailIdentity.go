// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an email address to the list of identities for your Amazon SES account in
// the current AWS region and attempts to verify it. As a result of executing this
// operation, a verification email is sent to the specified address. You can
// execute this operation no more than once per second.
func (c *Client) VerifyEmailIdentity(ctx context.Context, params *VerifyEmailIdentityInput, optFns ...func(*Options)) (*VerifyEmailIdentityOutput, error) {
	if params == nil {
		params = &VerifyEmailIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyEmailIdentity", params, optFns, c.addOperationVerifyEmailIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyEmailIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to begin email address verification with Amazon SES. For
// information about email address verification, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/verify-email-addresses.html).
type VerifyEmailIdentityInput struct {

	// The email address to be verified.
	//
	// This member is required.
	EmailAddress *string
}

// An empty element returned on a successful request.
type VerifyEmailIdentityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationVerifyEmailIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpVerifyEmailIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpVerifyEmailIdentity{}, middleware.After)
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
	if err = addOpVerifyEmailIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyEmailIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifyEmailIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "VerifyEmailIdentity",
	}
}
