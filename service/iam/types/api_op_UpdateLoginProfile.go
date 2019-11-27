// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateLoginProfileInput struct {
	_ struct{} `type:"structure"`

	// The new password for the specified IAM user.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// However, the format can be further restricted by the account administrator
	// by setting a password policy on the AWS account. For more information, see
	// UpdateAccountPasswordPolicy.
	Password *string `min:"1" type:"string" sensitive:"true"`

	// Allows this new password to be used only once by requiring the specified
	// IAM user to set a new password on next sign-in.
	PasswordResetRequired *bool `type:"boolean"`

	// The name of the user whose password you want to update.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateLoginProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLoginProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLoginProfileInput"}
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

type UpdateLoginProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateLoginProfileOutput) String() string {
	return awsutil.Prettify(s)
}
