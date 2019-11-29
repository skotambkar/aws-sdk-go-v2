// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the GetId action.
type GetIdInput struct {
	_ struct{} `type:"structure"`

	// A standard AWS account ID (9+ digits).
	AccountId *string `min:"1" type:"string"`

	// An identity pool ID in the format REGION:GUID.
	//
	// IdentityPoolId is a required field
	IdentityPoolId *string `min:"1" type:"string" required:"true"`

	// A set of optional name-value pairs that map provider names to provider tokens.
	// The available provider names for Logins are as follows:
	//
	//    * Facebook: graph.facebook.com
	//
	//    * Amazon Cognito user pool: cognito-idp.<region>.amazonaws.com/<YOUR_USER_POOL_ID>,
	//    for example, cognito-idp.us-east-1.amazonaws.com/us-east-1_123456789.
	//
	//    * Google: accounts.google.com
	//
	//    * Amazon: www.amazon.com
	//
	//    * Twitter: api.twitter.com
	//
	//    * Digits: www.digits.com
	Logins map[string]string `type:"map"`
}

// String returns the string representation
func (s GetIdInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetIdInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetIdInput"}
	if s.AccountId != nil && len(*s.AccountId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 1))
	}

	if s.IdentityPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityPoolId"))
	}
	if s.IdentityPoolId != nil && len(*s.IdentityPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returned in response to a GetId request.
type GetIdOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier in the format REGION:GUID.
	IdentityId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetIdOutput) String() string {
	return awsutil.Prettify(s)
}