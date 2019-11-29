// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRoomInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The room name.
	Name *string `type:"string" sensitive:"true"`

	// The room ID.
	//
	// RoomId is a required field
	RoomId *string `location:"uri" locationName:"roomId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateRoomInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRoomInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRoomInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.RoomId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRoomOutput struct {
	_ struct{} `type:"structure"`

	// The room details.
	Room *Room `type:"structure"`
}

// String returns the string representation
func (s UpdateRoomOutput) String() string {
	return awsutil.Prettify(s)
}
