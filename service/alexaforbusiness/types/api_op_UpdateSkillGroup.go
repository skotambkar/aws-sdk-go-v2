// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSkillGroupInput struct {
	_ struct{} `type:"structure"`

	// The updated description for the skill group.
	Description *string `min:"1" type:"string"`

	// The ARN of the skill group to update.
	SkillGroupArn *string `type:"string"`

	// The updated name for the skill group.
	SkillGroupName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateSkillGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSkillGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSkillGroupInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.SkillGroupName != nil && len(*s.SkillGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SkillGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSkillGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateSkillGroupOutput) String() string {
	return awsutil.Prettify(s)
}
