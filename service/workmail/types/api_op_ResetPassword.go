// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResetPasswordInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the organization that contains the user for which the password
	// is reset.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The new password for the user.
	//
	// Password is a required field
	Password *string `type:"string" required:"true" sensitive:"true"`

	// The identifier of the user for whom the password is reset.
	//
	// UserId is a required field
	UserId *string `min:"12" type:"string" required:"true"`
}

// String returns the string representation
func (s ResetPasswordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetPasswordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResetPasswordInput"}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResetPasswordOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ResetPasswordOutput) String() string {
	return awsutil.Prettify(s)
}
