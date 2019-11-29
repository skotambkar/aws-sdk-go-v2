// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInvitationConfigurationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetInvitationConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

type GetInvitationConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The email ID of the organization or individual contact that the enrolled
	// user can use.
	ContactEmail *string `min:"1" type:"string"`

	// The name of the organization sending the enrollment invite to a user.
	OrganizationName *string `min:"1" type:"string"`

	// The list of private skill IDs that you want to recommend to the user to enable
	// in the invitation.
	PrivateSkillIds []string `type:"list"`
}

// String returns the string representation
func (s GetInvitationConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
