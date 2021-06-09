// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a unique generated shared secret key code for the user account. The
// request takes an access token or a session string, but not both. Calling
// AssociateSoftwareToken immediately disassociates the existing software token
// from the user account. If the user doesn't subsequently verify the software
// token, their account is essentially set up to authenticate without MFA. If MFA
// config is set to Optional at the user pool level, the user can then login
// without MFA. However, if MFA is set to Required for the user pool, the user will
// be asked to setup a new software token MFA during sign in.
func (c *Client) AssociateSoftwareToken(ctx context.Context, params *AssociateSoftwareTokenInput, optFns ...func(*Options)) (*AssociateSoftwareTokenOutput, error) {
	if params == nil {
		params = &AssociateSoftwareTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateSoftwareToken", params, optFns, c.addOperationAssociateSoftwareTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateSoftwareTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSoftwareTokenInput struct {

	// The access token.
	AccessToken *string

	// The session which should be passed both ways in challenge-response calls to the
	// service. This allows authentication of the user as part of the MFA setup
	// process.
	Session *string
}

type AssociateSoftwareTokenOutput struct {

	// A unique generated shared secret code that is used in the TOTP algorithm to
	// generate a one time code.
	SecretCode *string

	// The session which should be passed both ways in challenge-response calls to the
	// service. This allows authentication of the user as part of the MFA setup
	// process.
	Session *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAssociateSoftwareTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateSoftwareToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateSoftwareToken{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSoftwareToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateSoftwareToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AssociateSoftwareToken",
	}
}
