// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the GetOpenIdToken action.
type GetOpenIdTokenInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier in the format REGION:GUID.
	//
	// IdentityId is a required field
	IdentityId *string `min:"1" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	// When using graph.facebook.com and www.amazon.com, supply the access_token
	// returned from the provider's authflow. For accounts.google.com, an Amazon
	// Cognito user pool provider, or any other OpenId Connect provider, always
	// include the id_token.
	Logins map[string]string `type:"map"`
}

// String returns the string representation
func (s GetOpenIdTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOpenIdTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOpenIdTokenInput"}

	if s.IdentityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityId"))
	}
	if s.IdentityId != nil && len(*s.IdentityId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returned in response to a successful GetOpenIdToken request.
type GetOpenIdTokenOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier in the format REGION:GUID. Note that the IdentityId returned
	// may not match the one passed on input.
	IdentityId *string `min:"1" type:"string"`

	// An OpenID token, valid for 10 minutes.
	Token *string `type:"string"`
}

// String returns the string representation
func (s GetOpenIdTokenOutput) String() string {
	return awsutil.Prettify(s)
}
