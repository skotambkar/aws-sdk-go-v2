// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/enums"
)

type ListCompilationJobsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns the model compilation jobs that were created after
	// a specified time.
	CreationTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns the model compilation jobs that were created before
	// a specified time.
	CreationTimeBefore *time.Time `type:"timestamp"`

	// A filter that returns the model compilation jobs that were modified after
	// a specified time.
	LastModifiedTimeAfter *time.Time `type:"timestamp"`

	// A filter that returns the model compilation jobs that were modified before
	// a specified time.
	LastModifiedTimeBefore *time.Time `type:"timestamp"`

	// The maximum number of model compilation jobs to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A filter that returns the model compilation jobs whose name contains a specified
	// string.
	NameContains *string `type:"string"`

	// If the result of the previous ListCompilationJobs request was truncated,
	// the response includes a NextToken. To retrieve the next set of model compilation
	// jobs, use the token in the next request.
	NextToken *string `type:"string"`

	// The field by which to sort results. The default is CreationTime.
	SortBy enums.ListCompilationJobsSortBy `type:"string" enum:"true"`

	// The sort order for results. The default is Ascending.
	SortOrder enums.SortOrder `type:"string" enum:"true"`

	// A filter that retrieves model compilation jobs with a specific DescribeCompilationJobResponse$CompilationJobStatus
	// status.
	StatusEquals enums.CompilationJobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListCompilationJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCompilationJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCompilationJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListCompilationJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of CompilationJobSummary objects, each describing a model compilation
	// job.
	//
	// CompilationJobSummaries is a required field
	CompilationJobSummaries []CompilationJobSummary `type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this NextToken. To
	// retrieve the next set of model compilation jobs, use this token in the next
	// request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCompilationJobsOutput) String() string {
	return awsutil.Prettify(s)
}