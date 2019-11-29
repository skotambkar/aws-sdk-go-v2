// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAttendeeInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK attendee ID.
	//
	// AttendeeId is a required field
	AttendeeId *string `location:"uri" locationName:"attendeeId" type:"string" required:"true"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAttendeeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAttendeeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAttendeeInput"}

	if s.AttendeeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttendeeId"))
	}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAttendeeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK attendee information.
	Attendee *Attendee `type:"structure"`
}

// String returns the string representation
func (s GetAttendeeOutput) String() string {
	return awsutil.Prettify(s)
}
