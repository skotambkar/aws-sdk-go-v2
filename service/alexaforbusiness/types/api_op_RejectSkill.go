// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RejectSkillInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the skill.
	//
	// SkillId is a required field
	SkillId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RejectSkillInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectSkillInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RejectSkillInput"}

	if s.SkillId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SkillId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RejectSkillOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RejectSkillOutput) String() string {
	return awsutil.Prettify(s)
}