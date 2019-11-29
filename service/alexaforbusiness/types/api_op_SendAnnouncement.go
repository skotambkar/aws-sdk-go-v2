// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SendAnnouncementInput struct {
	_ struct{} `type:"structure"`

	// The unique, user-specified identifier for the request that ensures idempotency.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `min:"10" type:"string" required:"true" idempotencyToken:"true"`

	// The announcement content. This can contain only one of the three possible
	// announcement types (text, SSML or audio).
	//
	// Content is a required field
	Content *Content `type:"structure" required:"true"`

	// The filters to use to send an announcement to a specified list of rooms.
	// The supported filter keys are RoomName, ProfileName, RoomArn, and ProfileArn.
	// To send to all rooms, specify an empty RoomFilters list.
	//
	// RoomFilters is a required field
	RoomFilters []Filter `type:"list" required:"true"`

	// The time to live for an announcement. Default is 300. If delivery doesn't
	// occur within this time, the announcement is not delivered.
	TimeToLiveInSeconds *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s SendAnnouncementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendAnnouncementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendAnnouncementInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 10))
	}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.RoomFilters == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomFilters"))
	}
	if s.TimeToLiveInSeconds != nil && *s.TimeToLiveInSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("TimeToLiveInSeconds", 1))
	}
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			invalidParams.AddNested("Content", err.(aws.ErrInvalidParams))
		}
	}
	if s.RoomFilters != nil {
		for i, v := range s.RoomFilters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RoomFilters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SendAnnouncementOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the announcement.
	AnnouncementArn *string `type:"string"`
}

// String returns the string representation
func (s SendAnnouncementOutput) String() string {
	return awsutil.Prettify(s)
}