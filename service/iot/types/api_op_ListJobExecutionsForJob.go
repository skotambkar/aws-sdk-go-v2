// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iot/enums"
)

type ListJobExecutionsForJobInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier you assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// The maximum number of results to be returned per request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The status of the job.
	Status enums.JobExecutionStatus `location:"querystring" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListJobExecutionsForJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJobExecutionsForJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJobExecutionsForJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListJobExecutionsForJobOutput struct {
	_ struct{} `type:"structure"`

	// A list of job execution summaries.
	ExecutionSummaries []JobExecutionSummaryForJob `locationName:"executionSummaries" type:"list"`

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListJobExecutionsForJobOutput) String() string {
	return awsutil.Prettify(s)
}