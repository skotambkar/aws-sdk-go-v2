// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// CancelTaskExecutionRequest
type CancelTaskExecutionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the task execution to cancel.
	//
	// TaskExecutionArn is a required field
	TaskExecutionArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CancelTaskExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelTaskExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelTaskExecutionInput"}

	if s.TaskExecutionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskExecutionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelTaskExecutionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelTaskExecutionOutput) String() string {
	return awsutil.Prettify(s)
}
