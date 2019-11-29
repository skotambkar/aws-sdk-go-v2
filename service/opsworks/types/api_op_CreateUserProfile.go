// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateUserProfileInput struct {
	_ struct{} `type:"structure"`

	// Whether users can specify their own SSH public key through the My Settings
	// page. For more information, see Setting an IAM User's Public SSH Key (https://docs.aws.amazon.com/opsworks/latest/userguide/security-settingsshkey.html).
	AllowSelfManagement *bool `type:"boolean"`

	// The user's IAM ARN; this can also be a federated user's ARN.
	//
	// IamUserArn is a required field
	IamUserArn *string `type:"string" required:"true"`

	// The user's public SSH key.
	SshPublicKey *string `type:"string"`

	// The user's SSH user name. The allowable characters are [a-z], [A-Z], [0-9],
	// '-', and '_'. If the specified name includes other punctuation marks, AWS
	// OpsWorks Stacks removes them. For example, my.name will be changed to myname.
	// If you do not specify an SSH user name, AWS OpsWorks Stacks generates one
	// from the IAM user name.
	SshUsername *string `type:"string"`
}

// String returns the string representation
func (s CreateUserProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserProfileInput"}

	if s.IamUserArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("IamUserArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a CreateUserProfile request.
type CreateUserProfileOutput struct {
	_ struct{} `type:"structure"`

	// The user's IAM ARN.
	IamUserArn *string `type:"string"`
}

// String returns the string representation
func (s CreateUserProfileOutput) String() string {
	return awsutil.Prettify(s)
}
