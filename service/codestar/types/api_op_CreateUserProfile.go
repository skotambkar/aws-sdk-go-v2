// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateUserProfileInput struct {
	_ struct{} `type:"structure"`

	// The name that will be displayed as the friendly name for the user in AWS
	// CodeStar.
	//
	// DisplayName is a required field
	DisplayName *string `locationName:"displayName" min:"1" type:"string" required:"true" sensitive:"true"`

	// The email address that will be displayed as part of the user's profile in
	// AWS CodeStar.
	//
	// EmailAddress is a required field
	EmailAddress *string `locationName:"emailAddress" min:"3" type:"string" required:"true" sensitive:"true"`

	// The SSH public key associated with the user in AWS CodeStar. If a project
	// owner allows the user remote access to project resources, this public key
	// will be used along with the user's private key for SSH access.
	SshPublicKey *string `locationName:"sshPublicKey" type:"string"`

	// The Amazon Resource Name (ARN) of the user in IAM.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserProfileInput"}

	if s.DisplayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DisplayName"))
	}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}

	if s.EmailAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailAddress"))
	}
	if s.EmailAddress != nil && len(*s.EmailAddress) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("EmailAddress", 3))
	}

	if s.UserArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserArn"))
	}
	if s.UserArn != nil && len(*s.UserArn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("UserArn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateUserProfileOutput struct {
	_ struct{} `type:"structure"`

	// The date the user profile was created, in timestamp format.
	CreatedTimestamp *time.Time `locationName:"createdTimestamp" type:"timestamp"`

	// The name that is displayed as the friendly name for the user in AWS CodeStar.
	DisplayName *string `locationName:"displayName" min:"1" type:"string" sensitive:"true"`

	// The email address that is displayed as part of the user's profile in AWS
	// CodeStar.
	EmailAddress *string `locationName:"emailAddress" min:"3" type:"string" sensitive:"true"`

	// The date the user profile was last modified, in timestamp format.
	LastModifiedTimestamp *time.Time `locationName:"lastModifiedTimestamp" type:"timestamp"`

	// The SSH public key associated with the user in AWS CodeStar. This is the
	// public portion of the public/private keypair the user can use to access project
	// resources if a project owner allows the user remote access to those resources.
	SshPublicKey *string `locationName:"sshPublicKey" type:"string"`

	// The Amazon Resource Name (ARN) of the user in IAM.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserProfileOutput) String() string {
	return awsutil.Prettify(s)
}
