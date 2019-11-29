// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/enums"
)

// Represents the request to update the user pool client.
type UpdateUserPoolClientInput struct {
	_ struct{} `type:"structure"`

	// Set to code to initiate a code grant flow, which provides an authorization
	// code as the response. This code can be exchanged for access tokens with the
	// token endpoint.
	AllowedOAuthFlows []enums.OAuthFlowType `type:"list"`

	// Set to TRUE if the client is allowed to follow the OAuth protocol when interacting
	// with Cognito user pools.
	AllowedOAuthFlowsUserPoolClient *bool `type:"boolean"`

	// A list of allowed OAuth scopes. Currently supported values are "phone", "email",
	// "openid", and "Cognito". In addition to these values, custom scopes created
	// in Resource Servers are also supported.
	AllowedOAuthScopes []string `type:"list"`

	// The Amazon Pinpoint analytics configuration for collecting metrics for this
	// user pool.
	AnalyticsConfiguration *AnalyticsConfigurationType `type:"structure"`

	// A list of allowed redirect (callback) URLs for the identity providers.
	//
	// A redirect URI must:
	//
	//    * Be an absolute URI.
	//
	//    * Be registered with the authorization server.
	//
	//    * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2).
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	CallbackURLs []string `type:"list"`

	// The ID of the client associated with the user pool.
	//
	// ClientId is a required field
	ClientId *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// The client name from the update user pool client request.
	ClientName *string `min:"1" type:"string"`

	// The default redirect URI. Must be in the CallbackURLs list.
	//
	// A redirect URI must:
	//
	//    * Be an absolute URI.
	//
	//    * Be registered with the authorization server.
	//
	//    * Not include a fragment component.
	//
	// See OAuth 2.0 - Redirection Endpoint (https://tools.ietf.org/html/rfc6749#section-3.1.2).
	//
	// Amazon Cognito requires HTTPS over HTTP except for http://localhost for testing
	// purposes only.
	//
	// App callback URLs such as myapp://example are also supported.
	DefaultRedirectURI *string `min:"1" type:"string"`

	// The authentication flows that are supported by the user pool clients. Flow
	// names without the ALLOW_ prefix are deprecated in favor of new names with
	// the ALLOW_ prefix. Note that values with ALLOW_ prefix cannot be used along
	// with values without ALLOW_ prefix.
	//
	// Valid values include:
	//
	//    * ALLOW_ADMIN_USER_PASSWORD_AUTH: Enable admin based user password authentication
	//    flow ADMIN_USER_PASSWORD_AUTH. This setting replaces the ADMIN_NO_SRP_AUTH
	//    setting. With this authentication flow, Cognito receives the password
	//    in the request instead of using the SRP (Secure Remote Password protocol)
	//    protocol to verify passwords.
	//
	//    * ALLOW_CUSTOM_AUTH: Enable Lambda trigger based authentication.
	//
	//    * ALLOW_USER_PASSWORD_AUTH: Enable user password-based authentication.
	//    In this flow, Cognito receives the password in the request instead of
	//    using the SRP protocol to verify passwords.
	//
	//    * ALLOW_USER_SRP_AUTH: Enable SRP based authentication.
	//
	//    * ALLOW_REFRESH_TOKEN_AUTH: Enable authflow to refresh tokens.
	ExplicitAuthFlows []enums.ExplicitAuthFlowsType `type:"list"`

	// A list of allowed logout URLs for the identity providers.
	LogoutURLs []string `type:"list"`

	// Use this setting to choose which errors and responses are returned by Cognito
	// APIs during authentication, account confirmation, and password recovery when
	// the user does not exist in the user pool. When set to ENABLED and the user
	// does not exist, authentication returns an error indicating either the username
	// or password was incorrect, and account confirmation and password recovery
	// return a response indicating a code was sent to a simulated destination.
	// When set to LEGACY, those APIs will return a UserNotFoundException exception
	// if the user does not exist in the user pool.
	//
	// Valid values include:
	//
	//    * ENABLED - This prevents user existence-related errors.
	//
	//    * LEGACY - This represents the old behavior of Cognito where user existence
	//    related errors are not prevented.
	//
	// This setting affects the behavior of following APIs:
	//
	//    * AdminInitiateAuth
	//
	//    * AdminRespondToAuthChallenge
	//
	//    * InitiateAuth
	//
	//    * RespondToAuthChallenge
	//
	//    * ForgotPassword
	//
	//    * ConfirmForgotPassword
	//
	//    * ConfirmSignUp
	//
	//    * ResendConfirmationCode
	//
	// After January 1st 2020, the value of PreventUserExistenceErrors will default
	// to ENABLED for newly created user pool clients if no value is provided.
	PreventUserExistenceErrors enums.PreventUserExistenceErrorTypes `type:"string" enum:"true"`

	// The read-only attributes of the user pool.
	ReadAttributes []string `type:"list"`

	// The time limit, in days, after which the refresh token is no longer valid
	// and cannot be used.
	RefreshTokenValidity *int64 `type:"integer"`

	// A list of provider names for the identity providers that are supported on
	// this client.
	SupportedIdentityProviders []string `type:"list"`

	// The user pool ID for the user pool where you want to update the user pool
	// client.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The writeable attributes of the user pool.
	WriteAttributes []string `type:"list"`
}

// String returns the string representation
func (s UpdateUserPoolClientInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserPoolClientInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserPoolClientInput"}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}
	if s.ClientName != nil && len(*s.ClientName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientName", 1))
	}
	if s.DefaultRedirectURI != nil && len(*s.DefaultRedirectURI) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultRedirectURI", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.AnalyticsConfiguration != nil {
		if err := s.AnalyticsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("AnalyticsConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server to the request to update the user
// pool client.
type UpdateUserPoolClientOutput struct {
	_ struct{} `type:"structure"`

	// The user pool client value from the response from the server when an update
	// user pool client request is made.
	UserPoolClient *UserPoolClientType `type:"structure"`
}

// String returns the string representation
func (s UpdateUserPoolClientOutput) String() string {
	return awsutil.Prettify(s)
}
