// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the user's multi-factor authentication (MFA) preference, including which
// MFA options are enabled and if any are preferred. Only one factor can be set as
// preferred. The preferred MFA factor will be used to authenticate a user if
// multiple factors are enabled. If multiple options are enabled and no preference
// is set, a challenge to choose an MFA option will be returned during sign in.
func (c *Client) AdminSetUserMFAPreference(ctx context.Context, params *AdminSetUserMFAPreferenceInput, optFns ...func(*Options)) (*AdminSetUserMFAPreferenceOutput, error) {
	if params == nil {
		params = &AdminSetUserMFAPreferenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminSetUserMFAPreference", params, optFns, c.addOperationAdminSetUserMFAPreferenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminSetUserMFAPreferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminSetUserMFAPreferenceInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The user pool username or alias.
	//
	// This member is required.
	Username *string

	// The SMS text message MFA settings.
	SMSMfaSettings *types.SMSMfaSettingsType

	// The time-based one-time password software token MFA settings.
	SoftwareTokenMfaSettings *types.SoftwareTokenMfaSettingsType
}

type AdminSetUserMFAPreferenceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAdminSetUserMFAPreferenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminSetUserMFAPreference{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminSetUserMFAPreference{}, middleware.After)
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
	if err = addOpAdminSetUserMFAPreferenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminSetUserMFAPreference(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminSetUserMFAPreference(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminSetUserMFAPreference",
	}
}
