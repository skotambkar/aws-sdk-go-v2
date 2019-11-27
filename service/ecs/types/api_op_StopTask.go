// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopTaskInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the task to stop. If you do not specify a cluster, the default cluster is
	// assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// An optional message specified when a task is stopped. For example, if you
	// are using a custom scheduler, you can use this parameter to specify the reason
	// for stopping the task here, and the message appears in subsequent DescribeTasks
	// API operations on this task. Up to 255 characters are allowed in this message.
	Reason *string `locationName:"reason" type:"string"`

	// The task ID or full Amazon Resource Name (ARN) of the task to stop.
	//
	// Task is a required field
	Task *string `locationName:"task" type:"string" required:"true"`
}

// String returns the string representation
func (s StopTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopTaskInput"}

	if s.Task == nil {
		invalidParams.Add(aws.NewErrParamRequired("Task"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopTaskOutput struct {
	_ struct{} `type:"structure"`

	// The task that was stopped.
	Task *Task `locationName:"task" type:"structure"`
}

// String returns the string representation
func (s StopTaskOutput) String() string {
	return awsutil.Prettify(s)
}
