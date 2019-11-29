// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TestIdentityProviderInput struct {
	_ struct{} `type:"structure"`

	// A system-assigned identifier for a specific server. That server's user authentication
	// method is tested with a user name and password.
	//
	// ServerId is a required field
	ServerId *string `min:"19" type:"string" required:"true"`

	// This request parameter is the name of the user account to be tested.
	//
	// UserName is a required field
	UserName *string `min:"3" type:"string" required:"true"`

	// The password of the user account to be tested.
	UserPassword *string `type:"string" sensitive:"true"`
}

// String returns the string representation
func (s TestIdentityProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestIdentityProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestIdentityProviderInput"}

	if s.ServerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerId"))
	}
	if s.ServerId != nil && len(*s.ServerId) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerId", 19))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TestIdentityProviderOutput struct {
	_ struct{} `type:"structure"`

	// A message that indicates whether the test was successful or not.
	Message *string `type:"string"`

	// The response that is returned from your API Gateway.
	Response *string `type:"string"`

	// The HTTP status code that is the response from your API Gateway.
	//
	// StatusCode is a required field
	StatusCode *int64 `type:"integer" required:"true"`

	// The endpoint of the service used to authenticate a user.
	//
	// Url is a required field
	Url *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TestIdentityProviderOutput) String() string {
	return awsutil.Prettify(s)
}
