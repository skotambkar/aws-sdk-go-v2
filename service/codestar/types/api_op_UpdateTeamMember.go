// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateTeamMemberInput struct {
	_ struct{} `type:"structure"`

	// The ID of the project.
	//
	// ProjectId is a required field
	ProjectId *string `locationName:"projectId" min:"2" type:"string" required:"true"`

	// The role assigned to the user in the project. Project roles have different
	// levels of access. For more information, see Working with Teams (http://docs.aws.amazon.com/codestar/latest/userguide/working-with-teams.html)
	// in the AWS CodeStar User Guide.
	ProjectRole *string `locationName:"projectRole" type:"string"`

	// Whether a team member is allowed to remotely access project resources using
	// the SSH public key associated with the user's profile. Even if this is set
	// to True, the user must associate a public key with their profile before the
	// user can access resources.
	RemoteAccessAllowed *bool `locationName:"remoteAccessAllowed" type:"boolean"`

	// The Amazon Resource Name (ARN) of the user for whom you want to change team
	// membership attributes.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTeamMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTeamMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTeamMemberInput"}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}
	if s.ProjectId != nil && len(*s.ProjectId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectId", 2))
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

type UpdateTeamMemberOutput struct {
	_ struct{} `type:"structure"`

	// The project role granted to the user.
	ProjectRole *string `locationName:"projectRole" type:"string"`

	// Whether a team member is allowed to remotely access project resources using
	// the SSH public key associated with the user's profile.
	RemoteAccessAllowed *bool `locationName:"remoteAccessAllowed" type:"boolean"`

	// The Amazon Resource Name (ARN) of the user whose team membership attributes
	// were updated.
	UserArn *string `locationName:"userArn" min:"32" type:"string"`
}

// String returns the string representation
func (s UpdateTeamMemberOutput) String() string {
	return awsutil.Prettify(s)
}
