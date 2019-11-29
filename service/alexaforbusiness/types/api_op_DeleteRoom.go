// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRoomInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the room to delete. Required.
	RoomArn *string `type:"string"`
}

// String returns the string representation
func (s DeleteRoomInput) String() string {
	return awsutil.Prettify(s)
}

type DeleteRoomOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRoomOutput) String() string {
	return awsutil.Prettify(s)
}