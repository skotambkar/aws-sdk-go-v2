// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/enums"
)

type UpdateAcceleratorInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the accelerator to update.
	//
	// AcceleratorArn is a required field
	AcceleratorArn *string `type:"string" required:"true"`

	// Indicates whether an accelerator is enabled. The value is true or false.
	// The default value is true.
	//
	// If the value is set to true, the accelerator cannot be deleted. If set to
	// false, the accelerator can be deleted.
	Enabled *bool `type:"boolean"`

	// The value for the address type must be IPv4.
	IpAddressType enums.IpAddressType `type:"string" enum:"true"`

	// The name of the accelerator. The name can have a maximum of 32 characters,
	// must contain only alphanumeric characters or hyphens (-), and must not begin
	// or end with a hyphen.
	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateAcceleratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAcceleratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAcceleratorInput"}

	if s.AcceleratorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceleratorArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAcceleratorOutput struct {
	_ struct{} `type:"structure"`

	// Information about the updated accelerator.
	Accelerator *Accelerator `type:"structure"`
}

// String returns the string representation
func (s UpdateAcceleratorOutput) String() string {
	return awsutil.Prettify(s)
}