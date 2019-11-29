// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccountPasswordPolicyInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetAccountPasswordPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a successful GetAccountPasswordPolicy request.
type GetAccountPasswordPolicyOutput struct {
	_ struct{} `type:"structure"`

	// A structure that contains details about the account's password policy.
	//
	// PasswordPolicy is a required field
	PasswordPolicy *PasswordPolicy `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetAccountPasswordPolicyOutput) String() string {
	return awsutil.Prettify(s)
}
