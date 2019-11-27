// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for the OptInPhoneNumber action.
type OptInPhoneNumberInput struct {
	_ struct{} `type:"structure"`

	// The phone number to opt in.
	//
	// PhoneNumber is a required field
	PhoneNumber *string `locationName:"phoneNumber" type:"string" required:"true"`
}

// String returns the string representation
func (s OptInPhoneNumberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OptInPhoneNumberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OptInPhoneNumberInput"}

	if s.PhoneNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response for the OptInPhoneNumber action.
type OptInPhoneNumberOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s OptInPhoneNumberOutput) String() string {
	return awsutil.Prettify(s)
}
