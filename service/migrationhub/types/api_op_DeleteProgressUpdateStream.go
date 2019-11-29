// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteProgressUpdateStreamInput struct {
	_ struct{} `type:"structure"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// The name of the ProgressUpdateStream. Do not store personal data in this
	// field.
	//
	// ProgressUpdateStreamName is a required field
	ProgressUpdateStreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProgressUpdateStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProgressUpdateStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProgressUpdateStreamInput"}

	if s.ProgressUpdateStreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStreamName"))
	}
	if s.ProgressUpdateStreamName != nil && len(*s.ProgressUpdateStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteProgressUpdateStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteProgressUpdateStreamOutput) String() string {
	return awsutil.Prettify(s)
}