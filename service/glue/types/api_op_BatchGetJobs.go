// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchGetJobsInput struct {
	_ struct{} `type:"structure"`

	// A list of job names, which might be the names returned from the ListJobs
	// operation.
	//
	// JobNames is a required field
	JobNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetJobsInput"}

	if s.JobNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchGetJobsOutput struct {
	_ struct{} `type:"structure"`

	// A list of job definitions.
	Jobs []Job `type:"list"`

	// A list of names of jobs not found.
	JobsNotFound []string `type:"list"`
}

// String returns the string representation
func (s BatchGetJobsOutput) String() string {
	return awsutil.Prettify(s)
}