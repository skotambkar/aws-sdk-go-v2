// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchCreateAttendeeInput struct {
	_ struct{} `type:"structure"`

	// The request containing the attendees to create.
	//
	// Attendees is a required field
	Attendees []CreateAttendeeRequestItem `type:"list" required:"true"`

	// The Amazon Chime SDK meeting ID.
	//
	// MeetingId is a required field
	MeetingId *string `location:"uri" locationName:"meetingId" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchCreateAttendeeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchCreateAttendeeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchCreateAttendeeInput"}

	if s.Attendees == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attendees"))
	}

	if s.MeetingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeetingId"))
	}
	if s.Attendees != nil {
		for i, v := range s.Attendees {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attendees", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchCreateAttendeeOutput struct {
	_ struct{} `type:"structure"`

	// The attendee information, including attendees IDs and join tokens.
	Attendees []Attendee `type:"list"`

	// If the action fails for one or more of the attendees in the request, a list
	// of the attendees is returned, along with error codes and error messages.
	Errors []CreateAttendeeError `type:"list"`
}

// String returns the string representation
func (s BatchCreateAttendeeOutput) String() string {
	return awsutil.Prettify(s)
}