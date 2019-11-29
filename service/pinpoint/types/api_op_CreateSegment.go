// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSegmentInput struct {
	_ struct{} `type:"structure" payload:"WriteSegmentRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the configuration, dimension, and other settings for a segment.
	// A WriteSegmentRequest object can include a Dimensions object or a SegmentGroups
	// object, but not both.
	//
	// WriteSegmentRequest is a required field
	WriteSegmentRequest *WriteSegmentRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateSegmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSegmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSegmentInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.WriteSegmentRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("WriteSegmentRequest"))
	}
	if s.WriteSegmentRequest != nil {
		if err := s.WriteSegmentRequest.Validate(); err != nil {
			invalidParams.AddNested("WriteSegmentRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSegmentOutput struct {
	_ struct{} `type:"structure" payload:"SegmentResponse"`

	// Provides information about the configuration, dimension, and other settings
	// for a segment.
	//
	// SegmentResponse is a required field
	SegmentResponse *SegmentResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateSegmentOutput) String() string {
	return awsutil.Prettify(s)
}