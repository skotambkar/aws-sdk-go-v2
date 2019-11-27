// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RevokeInvitationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the enrollment invitation to revoke. Required.
	EnrollmentId *string `type:"string"`

	// The ARN of the user for whom to revoke an enrollment invitation. Required.
	UserArn *string `type:"string"`
}

// String returns the string representation
func (s RevokeInvitationInput) String() string {
	return awsutil.Prettify(s)
}

type RevokeInvitationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RevokeInvitationOutput) String() string {
	return awsutil.Prettify(s)
}
