// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSkillGroupInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the skill group for which to get details. Required.
	SkillGroupArn *string `type:"string"`
}

// String returns the string representation
func (s GetSkillGroupInput) String() string {
	return awsutil.Prettify(s)
}

type GetSkillGroupOutput struct {
	_ struct{} `type:"structure"`

	// The details of the skill group requested. Required.
	SkillGroup *SkillGroup `type:"structure"`
}

// String returns the string representation
func (s GetSkillGroupOutput) String() string {
	return awsutil.Prettify(s)
}