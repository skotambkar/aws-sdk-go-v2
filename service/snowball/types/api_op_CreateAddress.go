// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateAddressInput struct {
	_ struct{} `type:"structure"`

	// The address that you want the Snowball shipped to.
	//
	// Address is a required field
	Address *Address `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateAddressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAddressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAddressInput"}

	if s.Address == nil {
		invalidParams.Add(aws.NewErrParamRequired("Address"))
	}
	if s.Address != nil {
		if err := s.Address.Validate(); err != nil {
			invalidParams.AddNested("Address", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateAddressOutput struct {
	_ struct{} `type:"structure"`

	// The automatically generated ID for a specific address. You'll use this ID
	// when you create a job to specify which address you want the Snowball for
	// that job shipped to.
	AddressId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateAddressOutput) String() string {
	return awsutil.Prettify(s)
}