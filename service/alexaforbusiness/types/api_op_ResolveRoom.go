// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResolveRoomInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the skill that was requested. Required.
	//
	// SkillId is a required field
	SkillId *string `type:"string" required:"true"`

	// The ARN of the user. Required.
	//
	// UserId is a required field
	UserId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResolveRoomInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResolveRoomInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResolveRoomInput"}

	if s.SkillId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SkillId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ResolveRoomOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the room from which the skill request was invoked.
	RoomArn *string `type:"string"`

	// The name of the room from which the skill request was invoked.
	RoomName *string `min:"1" type:"string"`

	// Response to get the room profile request. Required.
	RoomSkillParameters []RoomSkillParameter `type:"list"`
}

// String returns the string representation
func (s ResolveRoomOutput) String() string {
	return awsutil.Prettify(s)
}
