// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAccessKeyInput struct {
	_ struct{} `type:"structure"`

	// The access key ID for the access key ID and secret access key you want to
	// delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// AccessKeyId is a required field
	AccessKeyId *string `min:"16" type:"string" required:"true"`

	// The name of the user whose access key pair you want to delete.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteAccessKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccessKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAccessKeyInput"}

	if s.AccessKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessKeyId"))
	}
	if s.AccessKeyId != nil && len(*s.AccessKeyId) < 16 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessKeyId", 16))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAccessKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccessKeyOutput) String() string {
	return awsutil.Prettify(s)
}
