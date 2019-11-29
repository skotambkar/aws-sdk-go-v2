// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteMeetingInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMeetingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMeetingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMeetingInput"}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMeetingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMeetingOutput) String() string {
	return awsutil.Prettify(s)
}
