// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateVirtualMFADeviceInput struct {
	_ struct{} `type:"structure"`

	// The path for the virtual MFA device. For more information about paths, see
	// IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/).
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of either a forward slash (/) by itself
	// or a string that must begin and end with forward slashes. In addition, it
	// can contain any ASCII character from the ! (\u0021) through the DEL character
	// (\u007F), including most punctuation characters, digits, and upper and lowercased
	// letters.
	Path *string `min:"1" type:"string"`

	// The name of the virtual MFA device. Use with path to uniquely identify a
	// virtual MFA device.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// VirtualMFADeviceName is a required field
	VirtualMFADeviceName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVirtualMFADeviceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVirtualMFADeviceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVirtualMFADeviceInput"}
	if s.Path != nil && len(*s.Path) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Path", 1))
	}

	if s.VirtualMFADeviceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualMFADeviceName"))
	}
	if s.VirtualMFADeviceName != nil && len(*s.VirtualMFADeviceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualMFADeviceName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful CreateVirtualMFADevice request.
type CreateVirtualMFADeviceOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing details about the new virtual MFA device.
	//
	// VirtualMFADevice is a required field
	VirtualMFADevice *VirtualMFADevice `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVirtualMFADeviceOutput) String() string {
	return awsutil.Prettify(s)
}
