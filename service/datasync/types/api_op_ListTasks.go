// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// ListTasksRequest
type ListTasksInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of tasks to return.
	MaxResults *int64 `type:"integer"`

	// An opaque string that indicates the position at which to begin the next list
	// of tasks.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListTasksInput) String() string {
	return awsutil.Prettify(s)
}

// ListTasksResponse
type ListTasksOutput struct {
	_ struct{} `type:"structure"`

	// An opaque string that indicates the position at which to begin returning
	// the next list of tasks.
	NextToken *string `type:"string"`

	// A list of all the tasks that are returned.
	Tasks []TaskListEntry `type:"list"`
}

// String returns the string representation
func (s ListTasksOutput) String() string {
	return awsutil.Prettify(s)
}