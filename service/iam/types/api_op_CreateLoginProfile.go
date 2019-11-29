// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateLoginProfileInput struct {
	_ struct{} `type:"structure"`

	// The new password for the user.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of characters. That string can include almost
	// any printable ASCII character from the space (\u0020) through the end of
	// the ASCII character range (\u00FF). You can also include the tab (\u0009),
	// line feed (\u000A), and carriage return (\u000D) characters. Any of these
	// characters are valid in a password. However, many tools, such as the AWS
	// Management Console, might restrict the ability to type certain characters
	// because they have special meaning within that tool.
	//
	// Password is a required field
	Password *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// Specifies whether the user is required to set a new password on next sign-in.
	PasswordResetRequired *bool `type:"boolean"`

	// The name of the IAM user to create a password for. The user must already
	// exist.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLoginProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLoginProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLoginProfileInput"}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}
	if s.Password != nil && len(*s.Password) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Password", 1))
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

// Contains the response to a successful CreateLoginProfile request.
type CreateLoginProfileOutput struct {
	_ struct{} `type:"structure"`

	// A structure containing the user name and password create date.
	//
	// LoginProfile is a required field
	LoginProfile *LoginProfile `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateLoginProfileOutput) String() string {
	return awsutil.Prettify(s)
}