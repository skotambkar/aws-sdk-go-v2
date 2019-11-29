// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PhoneNumberValidateInput struct {
	_ struct{} `type:"structure" payload:"NumberValidateRequest"`

	// Specifies a phone number to validate and retrieve information about.
	//
	// NumberValidateRequest is a required field
	NumberValidateRequest *NumberValidateRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s PhoneNumberValidateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PhoneNumberValidateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PhoneNumberValidateInput"}

	if s.NumberValidateRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("NumberValidateRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PhoneNumberValidateOutput struct {
	_ struct{} `type:"structure" payload:"NumberValidateResponse"`

	// Provides information about a phone number.
	//
	// NumberValidateResponse is a required field
	NumberValidateResponse *NumberValidateResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s PhoneNumberValidateOutput) String() string {
	return awsutil.Prettify(s)
}