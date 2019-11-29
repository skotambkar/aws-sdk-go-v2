// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/enums"
)

type CreateAcceleratorInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether an accelerator is enabled. The value is true or false.
	// The default value is true.
	//
	// If the value is set to true, an accelerator cannot be deleted. If set to
	// false, the accelerator can be deleted.
	Enabled *bool `type:"boolean"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency—that
	// is, the uniqueness—of an accelerator.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `type:"string" required:"true"`

	// The value for the address type must be IPv4.
	IpAddressType enums.IpAddressType `type:"string" enum:"true"`

	// The name of an accelerator. The name can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens (-), and must not begin
	// or end with a hyphen.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAcceleratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAcceleratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAcceleratorInput"}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateAcceleratorOutput struct {
	_ struct{} `type:"structure"`

	// The accelerator that is created by specifying a listener and the supported
	// IP address types.
	Accelerator *Accelerator `type:"structure"`
}

// String returns the string representation
func (s CreateAcceleratorOutput) String() string {
	return awsutil.Prettify(s)
}