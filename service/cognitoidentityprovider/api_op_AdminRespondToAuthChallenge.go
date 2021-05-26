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

// Responds to an authentication challenge, as an administrator. Calling this
// action requires developer credentials.
func (c *Client) AdminRespondToAuthChallenge(ctx context.Context, params *AdminRespondToAuthChallengeInput, optFns ...func(*Options)) (*AdminRespondToAuthChallengeOutput, error) {
	if params == nil {
		params = &AdminRespondToAuthChallengeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminRespondToAuthChallenge", params, optFns, c.addOperationAdminRespondToAuthChallengeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminRespondToAuthChallengeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to respond to the authentication challenge, as an administrator.
type AdminRespondToAuthChallengeInput struct {

	// The challenge name. For more information, see AdminInitiateAuth
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminInitiateAuth.html).
	//
	// This member is required.
	ChallengeName types.ChallengeNameType

	// The app client ID.
	//
	// This member is required.
	ClientId *string

	// The ID of the Amazon Cognito user pool.
	//
	// This member is required.
	UserPoolId *string

	// The analytics metadata for collecting Amazon Pinpoint metrics for
	// AdminRespondToAuthChallenge calls.
	AnalyticsMetadata *types.AnalyticsMetadataType

	// The challenge responses. These are inputs corresponding to the value of
	// ChallengeName, for example:
	//
	// * SMS_MFA: SMS_MFA_CODE, USERNAME, SECRET_HASH (if
	// app client is configured with client secret).
	//
	// * PASSWORD_VERIFIER:
	// PASSWORD_CLAIM_SIGNATURE, PASSWORD_CLAIM_SECRET_BLOCK, TIMESTAMP, USERNAME,
	// SECRET_HASH (if app client is configured with client secret).
	//
	// *
	// ADMIN_NO_SRP_AUTH: PASSWORD, USERNAME, SECRET_HASH (if app client is configured
	// with client secret).
	//
	// * NEW_PASSWORD_REQUIRED: NEW_PASSWORD, any other required
	// attributes, USERNAME, SECRET_HASH (if app client is configured with client
	// secret).
	//
	// * MFA_SETUP requires USERNAME, plus you need to use the session value
	// returned by VerifySoftwareToken in the Session parameter.
	//
	// The value of the
	// USERNAME attribute must be the user's actual username, not an alias (such as
	// email address or phone number). To make this easier, the AdminInitiateAuth
	// response includes the actual username value in the USERNAMEUSER_ID_FOR_SRP
	// attribute, even if you specified an alias in your call to AdminInitiateAuth.
	ChallengeResponses map[string]string

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the
	// AdminRespondToAuthChallenge API action, Amazon Cognito invokes any functions
	// that are assigned to the following triggers: pre sign-up, custom message, post
	// authentication, user migration, pre token generation, define auth challenge,
	// create auth challenge, and verify auth challenge response. When Amazon Cognito
	// invokes any of these functions, it passes a JSON payload, which the function
	// receives as input. This payload contains a clientMetadata attribute, which
	// provides the data that you assigned to the ClientMetadata parameter in your
	// AdminRespondToAuthChallenge request. In your function code in AWS Lambda, you
	// can process the clientMetadata value to enhance your workflow for your specific
	// needs. For more information, see Customizing User Pool Workflows with Lambda
	// Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	// * Amazon Cognito does
	// not store the ClientMetadata value. This data is available only to AWS Lambda
	// triggers that are assigned to a user pool to support custom workflows. If your
	// user pool configuration does not include triggers, the ClientMetadata parameter
	// serves no purpose.
	//
	// * Amazon Cognito does not validate the ClientMetadata
	// value.
	//
	// * Amazon Cognito does not encrypt the the ClientMetadata value, so don't
	// use it to provide sensitive information.
	ClientMetadata map[string]string

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	ContextData *types.ContextDataType

	// The session which should be passed both ways in challenge-response calls to the
	// service. If InitiateAuth or RespondToAuthChallenge API call determines that the
	// caller needs to go through another challenge, they return a session with other
	// challenge parameters. This session should be passed as it is to the next
	// RespondToAuthChallenge API call.
	Session *string
}

// Responds to the authentication challenge, as an administrator.
type AdminRespondToAuthChallengeOutput struct {

	// The result returned by the server in response to the authentication request.
	AuthenticationResult *types.AuthenticationResultType

	// The name of the challenge. For more information, see AdminInitiateAuth
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminInitiateAuth.html).
	ChallengeName types.ChallengeNameType

	// The challenge parameters. For more information, see AdminInitiateAuth
	// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminInitiateAuth.html).
	ChallengeParameters map[string]string

	// The session which should be passed both ways in challenge-response calls to the
	// service. If the caller needs to go through another challenge, they return a
	// session with other challenge parameters. This session should be passed as it is
	// to the next RespondToAuthChallenge API call.
	Session *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationAdminRespondToAuthChallengeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminRespondToAuthChallenge{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminRespondToAuthChallenge{}, middleware.After)
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
	if err = addOpAdminRespondToAuthChallengeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminRespondToAuthChallenge(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminRespondToAuthChallenge(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminRespondToAuthChallenge",
	}
}
