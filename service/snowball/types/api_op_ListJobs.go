// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListJobsInput struct {
	_ struct{} `type:"structure"`

	// The number of JobListEntry objects to return.
	MaxResults *int64 `type:"integer"`

	// HTTP requests are stateless. To identify what object comes "next" in the
	// list of JobListEntry objects, you have the option of specifying NextToken
	// as the starting point for your returned list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJobsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListJobsOutput struct {
	_ struct{} `type:"structure"`

	// Each JobListEntry object contains a job's state, a job's ID, and a value
	// that indicates whether the job is a job part, in the case of export jobs.
	JobListEntries []JobListEntry `type:"list"`

	// HTTP requests are stateless. If you use this automatically generated NextToken
	// value in your next ListJobs call, your returned JobListEntry objects will
	// start from this point in the array.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListJobsOutput) String() string {
	return awsutil.Prettify(s)
}