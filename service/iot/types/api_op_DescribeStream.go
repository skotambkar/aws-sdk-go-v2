// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStreamInput struct {
	_ struct{} `type:"structure"`

	// The stream ID.
	//
	// StreamId is a required field
	StreamId *string `location:"uri" locationName:"streamId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStreamInput"}

	if s.StreamId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamId"))
	}
	if s.StreamId != nil && len(*s.StreamId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStreamOutput struct {
	_ struct{} `type:"structure"`

	// Information about the stream.
	StreamInfo *StreamInfo `locationName:"streamInfo" type:"structure"`
}

// String returns the string representation
func (s DescribeStreamOutput) String() string {
	return awsutil.Prettify(s)
}
