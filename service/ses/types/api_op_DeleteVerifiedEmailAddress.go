// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete an email address from the list of email addresses
// you have attempted to verify under your AWS account.
type DeleteVerifiedEmailAddressInput struct {
	_ struct{} `type:"structure"`

	// An email address to be removed from the list of verified addresses.
	//
	// EmailAddress is a required field
	EmailAddress *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVerifiedEmailAddressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVerifiedEmailAddressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVerifiedEmailAddressInput"}

	if s.EmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteVerifiedEmailAddressOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVerifiedEmailAddressOutput) String() string {
	return awsutil.Prettify(s)
}
