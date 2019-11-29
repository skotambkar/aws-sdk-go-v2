// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RemoveAccountFromOrganizationInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) of the member account that you want to remove
	// from the organization.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for an account ID string
	// requires exactly 12 digits.
	//
	// AccountId is a required field
	AccountId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveAccountFromOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveAccountFromOrganizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveAccountFromOrganizationInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveAccountFromOrganizationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveAccountFromOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}