// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchCreateAttendeeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Attendees != nil {
		v := s.Attendees

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attendees", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MeetingId != nil {
		v := *s.MeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchCreateAttendeeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Attendees != nil {
		v := s.Attendees

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Attendees", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Errors != nil {
		v := s.Errors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Errors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchCreateAttendee = "BatchCreateAttendee"

// BatchCreateAttendeeRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates up to 100 new attendees for an active Amazon Chime SDK meeting. For
// more information about the Amazon Chime SDK, see Using the Amazon Chime SDK
// (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html) in the Amazon
// Chime Developer Guide.
//
//    // Example sending a request using BatchCreateAttendeeRequest.
//    req := client.BatchCreateAttendeeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchCreateAttendee
func (c *Client) BatchCreateAttendeeRequest(input *BatchCreateAttendeeInput) BatchCreateAttendeeRequest {
	op := &aws.Operation{
		Name:       opBatchCreateAttendee,
		HTTPMethod: "POST",
		HTTPPath:   "/meetings/{meetingId}/attendees?operation=batch-create",
	}

	if input == nil {
		input = &BatchCreateAttendeeInput{}
	}

	req := c.newRequest(op, input, &BatchCreateAttendeeOutput{})
	return BatchCreateAttendeeRequest{Request: req, Input: input, Copy: c.BatchCreateAttendeeRequest}
}

// BatchCreateAttendeeRequest is the request type for the
// BatchCreateAttendee API operation.
type BatchCreateAttendeeRequest struct {
	*aws.Request
	Input *BatchCreateAttendeeInput
	Copy  func(*BatchCreateAttendeeInput) BatchCreateAttendeeRequest
}

// Send marshals and sends the BatchCreateAttendee API request.
func (r BatchCreateAttendeeRequest) Send(ctx context.Context) (*BatchCreateAttendeeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchCreateAttendeeResponse{
		BatchCreateAttendeeOutput: r.Request.Data.(*BatchCreateAttendeeOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchCreateAttendeeResponse is the response type for the
// BatchCreateAttendee API operation.
type BatchCreateAttendeeResponse struct {
	*BatchCreateAttendeeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchCreateAttendee request.
func (r *BatchCreateAttendeeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
