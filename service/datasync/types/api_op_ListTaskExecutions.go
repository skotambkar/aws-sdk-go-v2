// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// ListTaskExecutions
type ListTaskExecutionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of executed tasks to list.
	MaxResults *int64 `type:"integer"`

	// An opaque string that indicates the position at which to begin the next list
	// of the executed tasks.
	NextToken *string `type:"string"`

	// The Amazon Resource Name (ARN) of the task whose tasks you want to list.
	TaskArn *string `type:"string"`
}

// String returns the string representation
func (s ListTaskExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// ListTaskExecutionsResponse
type ListTaskExecutionsOutput struct {
	_ struct{} `type:"structure"`

	// An opaque string that indicates the position at which to begin returning
	// the next list of executed tasks.
	NextToken *string `type:"string"`

	// A list of executed tasks.
	TaskExecutions []TaskExecutionListEntry `type:"list"`
}

// String returns the string representation
func (s ListTaskExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}