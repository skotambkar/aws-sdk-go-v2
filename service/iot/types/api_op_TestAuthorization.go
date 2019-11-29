// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TestAuthorizationInput struct {
	_ struct{} `type:"structure"`

	// A list of authorization info objects. Simulating authorization will create
	// a response for each authInfo object in the list.
	//
	// AuthInfos is a required field
	AuthInfos []AuthInfo `locationName:"authInfos" min:"1" type:"list" required:"true"`

	// The MQTT client ID.
	ClientId *string `location:"querystring" locationName:"clientId" type:"string"`

	// The Cognito identity pool ID.
	CognitoIdentityPoolId *string `locationName:"cognitoIdentityPoolId" type:"string"`

	// When testing custom authorization, the policies specified here are treated
	// as if they are attached to the principal being authorized.
	PolicyNamesToAdd []string `locationName:"policyNamesToAdd" type:"list"`

	// When testing custom authorization, the policies specified here are treated
	// as if they are not attached to the principal being authorized.
	PolicyNamesToSkip []string `locationName:"policyNamesToSkip" type:"list"`

	// The principal.
	Principal *string `locationName:"principal" type:"string"`
}

// String returns the string representation
func (s TestAuthorizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestAuthorizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestAuthorizationInput"}

	if s.AuthInfos == nil {
		invalidParams.Add(aws.NewErrParamRequired("AuthInfos"))
	}
	if s.AuthInfos != nil && len(s.AuthInfos) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthInfos", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TestAuthorizationOutput struct {
	_ struct{} `type:"structure"`

	// The authentication results.
	AuthResults []AuthResult `locationName:"authResults" type:"list"`
}

// String returns the string representation
func (s TestAuthorizationOutput) String() string {
	return awsutil.Prettify(s)
}