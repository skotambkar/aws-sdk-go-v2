// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CountOpenWorkflowExecutionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain containing the workflow executions to count.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// If specified, only workflow executions matching the WorkflowId in the filter
	// are counted.
	//
	// executionFilter, typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	ExecutionFilter *WorkflowExecutionFilter `locationName:"executionFilter" type:"structure"`

	// Specifies the start time criteria that workflow executions must meet in order
	// to be counted.
	//
	// StartTimeFilter is a required field
	StartTimeFilter *ExecutionTimeFilter `locationName:"startTimeFilter" type:"structure" required:"true"`

	// If specified, only executions that have a tag that matches the filter are
	// counted.
	//
	// executionFilter, typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	TagFilter *TagFilter `locationName:"tagFilter" type:"structure"`

	// Specifies the type of the workflow executions to be counted.
	//
	// executionFilter, typeFilter and tagFilter are mutually exclusive. You can
	// specify at most one of these in a request.
	TypeFilter *WorkflowTypeFilter `locationName:"typeFilter" type:"structure"`
}

// String returns the string representation
func (s CountOpenWorkflowExecutionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CountOpenWorkflowExecutionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CountOpenWorkflowExecutionsInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.StartTimeFilter == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTimeFilter"))
	}
	if s.ExecutionFilter != nil {
		if err := s.ExecutionFilter.Validate(); err != nil {
			invalidParams.AddNested("ExecutionFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.StartTimeFilter != nil {
		if err := s.StartTimeFilter.Validate(); err != nil {
			invalidParams.AddNested("StartTimeFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TagFilter != nil {
		if err := s.TagFilter.Validate(); err != nil {
			invalidParams.AddNested("TagFilter", err.(aws.ErrInvalidParams))
		}
	}
	if s.TypeFilter != nil {
		if err := s.TypeFilter.Validate(); err != nil {
			invalidParams.AddNested("TypeFilter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the count of workflow executions returned from CountOpenWorkflowExecutions
// or CountClosedWorkflowExecutions
type CountOpenWorkflowExecutionsOutput struct {
	_ struct{} `type:"structure"`

	// The number of workflow executions.
	//
	// Count is a required field
	Count *int64 `locationName:"count" type:"integer" required:"true"`

	// If set to true, indicates that the actual count was more than the maximum
	// supported by this API and the count returned is the truncated value.
	Truncated *bool `locationName:"truncated" type:"boolean"`
}

// String returns the string representation
func (s CountOpenWorkflowExecutionsOutput) String() string {
	return awsutil.Prettify(s)
}
