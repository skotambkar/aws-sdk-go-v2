// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/enums"
)

type ListTrainingJobsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only training jobs created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only training jobs created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns only training jobs modified after the specified time
	// (timestamp).
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns only training jobs modified before the specified time
	// (timestamp).
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of training jobs to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the training job name. This filter returns only training jobs
	// whose name contains the specified string.
	NameContains *string `type:"string"`

	// If the result of the previous ListTrainingJobs request was truncated, the
	// response includes a NextToken. To retrieve the next set of training jobs,
	// use the token in the next request.
	NextToken *string `type:"string"`

	// The field to sort results by. The default is CreationTime.
	SortBy enums.SortBy `type:"string" enum:"true"`

	// The sort order for results. The default is Ascending.
	SortOrder enums.SortOrder `type:"string" enum:"true"`

	// A filter that retrieves only training jobs with a specific status.
	StatusEquals enums.TrainingJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListTrainingJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTrainingJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTrainingJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTrainingJobsOutput struct {
	_ struct{} `type:"structure"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of training jobs, use it in the subsequent request.
	NextToken *string `type:"string"`

	// An array of TrainingJobSummary objects, each listing a training job.
	//
	// TrainingJobSummaries is a required field
	TrainingJobSummaries []TrainingJobSummary `type:"list" required:"true"`
}

// String returns the string representation
func (s ListTrainingJobsOutput) String() string {
	return awsutil.Prettify(s)
}
