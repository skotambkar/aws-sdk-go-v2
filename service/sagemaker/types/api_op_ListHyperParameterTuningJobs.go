// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/enums"
)

type ListHyperParameterTuningJobsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only tuning jobs that were created after the specified
	// time.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only tuning jobs that were created before the specified
	// time.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns only tuning jobs that were modified after the specified
	// time.
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only tuning jobs that were modified before the specified
	// time.
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of tuning jobs to return. The default value is 10.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the tuning job name. This filter returns only tuning jobs whose
	// name contains the specified string.
	NameContains *string `type:"string"`

	// If the result of the previous ListHyperParameterTuningJobs request was truncated,
	// the response includes a NextToken. To retrieve the next set of tuning jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`

	// The field to sort results by. The default is Name.
	SortBy enums.HyperParameterTuningJobSortByOptions `type:"string" enum:"true"`

	// The sort order for results. The default is Ascending.
	SortOrder enums.SortOrder `type:"string" enum:"true"`

	// A filter that returns only tuning jobs with the specified status.
	StatusEquals enums.HyperParameterTuningJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListHyperParameterTuningJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListHyperParameterTuningJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListHyperParameterTuningJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListHyperParameterTuningJobsOutput struct {
	_ struct{} `type:"structure"`

	// A list of HyperParameterTuningJobSummary objects that describe the tuning
	// jobs that the ListHyperParameterTuningJobs request returned.
	//
	// HyperParameterTuningJobSummaries is a required field
	HyperParameterTuningJobSummaries []HyperParameterTuningJobSummary `type:"list" required:"true"`

	// If the result of this ListHyperParameterTuningJobs request was truncated,
	// the response includes a NextToken. To retrieve the next set of tuning jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListHyperParameterTuningJobsOutput) String() string {
	return awsutil.Prettify(s)
}
