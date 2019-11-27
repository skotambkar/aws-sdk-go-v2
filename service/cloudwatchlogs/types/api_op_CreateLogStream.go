// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateLogStreamInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The name of the log stream.
	//
	// LogStreamName is a required field
	LogStreamName *string `locationName:"logStreamName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLogStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLogStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLogStreamInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if s.LogStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogStreamName"))
	}
	if s.LogStreamName != nil && len(*s.LogStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogStreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLogStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateLogStreamOutput) String() string {
	return awsutil.Prettify(s)
}
