// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopAccessLoggingInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that you want to stop access logging on.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopAccessLoggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopAccessLoggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopAccessLoggingInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopAccessLoggingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopAccessLoggingOutput) String() string {
	return awsutil.Prettify(s)
}
