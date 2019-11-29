// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutEventsInput struct {
	_ struct{} `type:"structure" payload:"EventsRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies a batch of events to process.
	//
	// EventsRequest is a required field
	EventsRequest *EventsRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEventsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.EventsRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventsRequest"))
	}
	if s.EventsRequest != nil {
		if err := s.EventsRequest.Validate(); err != nil {
			invalidParams.AddNested("EventsRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutEventsOutput struct {
	_ struct{} `type:"structure" payload:"EventsResponse"`

	// Provides information about endpoints and the events that they're associated
	// with.
	//
	// EventsResponse is a required field
	EventsResponse *EventsResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutEventsOutput) String() string {
	return awsutil.Prettify(s)
}