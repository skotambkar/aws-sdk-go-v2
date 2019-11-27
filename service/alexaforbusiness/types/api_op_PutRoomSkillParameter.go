// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutRoomSkillParameterInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the room associated with the room skill parameter. Required.
	RoomArn *string `type:"string"`

	// The updated room skill parameter. Required.
	//
	// RoomSkillParameter is a required field
	RoomSkillParameter *RoomSkillParameter `type:"structure" required:"true"`

	// The ARN of the skill associated with the room skill parameter. Required.
	//
	// SkillId is a required field
	SkillId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutRoomSkillParameterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutRoomSkillParameterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutRoomSkillParameterInput"}

	if s.RoomSkillParameter == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomSkillParameter"))
	}

	if s.SkillId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SkillId"))
	}
	if s.RoomSkillParameter != nil {
		if err := s.RoomSkillParameter.Validate(); err != nil {
			invalidParams.AddNested("RoomSkillParameter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutRoomSkillParameterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutRoomSkillParameterOutput) String() string {
	return awsutil.Prettify(s)
}
