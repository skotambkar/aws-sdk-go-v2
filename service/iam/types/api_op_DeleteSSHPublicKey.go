// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteSSHPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the SSH public key.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// SSHPublicKeyId is a required field
	SSHPublicKeyId *string `min:"20" type:"string" required:"true"`

	// The name of the IAM user associated with the SSH public key.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSSHPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSSHPublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSSHPublicKeyInput"}

	if s.SSHPublicKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SSHPublicKeyId"))
	}
	if s.SSHPublicKeyId != nil && len(*s.SSHPublicKeyId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("SSHPublicKeyId", 20))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSSHPublicKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSSHPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}