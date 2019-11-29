// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a PollForJobs action.
type PollForJobsInput struct {
	_ struct{} `type:"structure"`

	// Represents information about an action type.
	//
	// ActionTypeId is a required field
	ActionTypeId *ActionTypeId `locationName:"actionTypeId" type:"structure" required:"true"`

	// The maximum number of jobs to return in a poll for jobs call.
	MaxBatchSize *int64 `locationName:"maxBatchSize" min:"1" type:"integer"`

	// A map of property names and values. For an action type with no queryable
	// properties, this value must be null or an empty map. For an action type with
	// a queryable property, you must supply that property as a key in the map.
	// Only jobs whose action configuration matches the mapped value are returned.
	QueryParam map[string]string `locationName:"queryParam" type:"map"`
}

// String returns the string representation
func (s PollForJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PollForJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PollForJobsInput"}

	if s.ActionTypeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActionTypeId"))
	}
	if s.MaxBatchSize != nil && *s.MaxBatchSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxBatchSize", 1))
	}
	if s.ActionTypeId != nil {
		if err := s.ActionTypeId.Validate(); err != nil {
			invalidParams.AddNested("ActionTypeId", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a PollForJobs action.
type PollForJobsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the jobs to take action on.
	Jobs []Job `locationName:"jobs" type:"list"`
}

// String returns the string representation
func (s PollForJobsOutput) String() string {
	return awsutil.Prettify(s)
}
