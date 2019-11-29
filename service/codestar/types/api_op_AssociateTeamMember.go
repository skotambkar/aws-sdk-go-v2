// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateTeamMemberInput struct {
	_ struct{} `type:"structure"`

	// A user- or system-generated token that identifies the entity that requested
	// the team member association to the project. This token can be used to repeat
	// the request.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string"`

	// The ID of the project to which you will add the IAM user.
	//
	// ProjectId is a required field
	ProjectId *string `locationName:"projectId" min:"2" type:"string" required:"true"`

	// The AWS CodeStar project role that will apply to this user. This role determines
	// what actions a user can take in an AWS CodeStar project.
	//
	// ProjectRole is a required field
	ProjectRole *string `locationName:"projectRole" type:"string" required:"true"`

	// Whether the team member is allowed to use an SSH public/private key pair
	// to remotely access project resources, for example Amazon EC2 instances.
	RemoteAccessAllowed *bool `locationName:"remoteAccessAllowed" type:"boolean"`

	// The Amazon Resource Name (ARN) for the IAM user you want to add to the AWS
	// CodeStar project.
	//
	// UserArn is a required field
	UserArn *string `locationName:"userArn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateTeamMemberInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateTeamMemberInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateTeamMemberInput"}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}
	if s.ProjectId != nil && len(*s.ProjectId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectId", 2))
	}

	if s.ProjectRole == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectRole"))
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

type AssociateTeamMemberOutput struct {
	_ struct{} `type:"structure"`

	// The user- or system-generated token from the initial request that can be
	// used to repeat the request.
	ClientRequestToken *string `locationName:"clientRequestToken" min:"1" type:"string"`
}

// String returns the string representation
func (s AssociateTeamMemberOutput) String() string {
	return awsutil.Prettify(s)
}
