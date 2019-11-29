// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMLTaskRunsInput struct {
	_ struct{} `type:"structure"`

	// The filter criteria, in the TaskRunFilterCriteria structure, for the task
	// run.
	Filter *TaskRunFilterCriteria `type:"structure"`

	// The maximum number of results to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token for pagination of the results. The default is empty.
	NextToken *string `type:"string"`

	// The sorting criteria, in the TaskRunSortCriteria structure, for the task
	// run.
	Sort *TaskRunSortCriteria `type:"structure"`

	// The unique identifier of the machine learning transform.
	//
	// TransformId is a required field
	TransformId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMLTaskRunsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMLTaskRunsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMLTaskRunsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.TransformId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransformId"))
	}
	if s.TransformId != nil && len(*s.TransformId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TransformId", 1))
	}
	if s.Sort != nil {
		if err := s.Sort.Validate(); err != nil {
			invalidParams.AddNested("Sort", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMLTaskRunsOutput struct {
	_ struct{} `type:"structure"`

	// A pagination token, if more results are available.
	NextToken *string `type:"string"`

	// A list of task runs that are associated with the transform.
	TaskRuns []TaskRun `type:"list"`
}

// String returns the string representation
func (s GetMLTaskRunsOutput) String() string {
	return awsutil.Prettify(s)
}