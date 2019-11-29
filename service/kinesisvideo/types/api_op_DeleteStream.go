// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteStreamInput struct {
	_ struct{} `type:"structure"`

	// Optional: The version of the stream that you want to delete.
	//
	// Specify the version as a safeguard to ensure that your are deleting the correct
	// stream. To get the stream version, use the DescribeStream API.
	//
	// If not specified, only the CreationTime is checked before deleting the stream.
	CurrentVersion *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the stream that you want to delete.
	//
	// StreamARN is a required field
	StreamARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteStreamInput"}
	if s.CurrentVersion != nil && len(*s.CurrentVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CurrentVersion", 1))
	}

	if s.StreamARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamARN"))
	}
	if s.StreamARN != nil && len(*s.StreamARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteStreamOutput) String() string {
	return awsutil.Prettify(s)
}