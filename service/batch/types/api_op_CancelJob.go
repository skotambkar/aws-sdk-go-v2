// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CancelJobInput struct {
	_ struct{} `type:"structure"`

	// The AWS Batch job ID of the job to cancel.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" type:"string" required:"true"`

	// A message to attach to the job that explains the reason for canceling it.
	// This message is returned by future DescribeJobs operations on the job. This
	// message is also recorded in the AWS Batch activity logs.
	//
	// Reason is a required field
	Reason *string `locationName:"reason" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if s.Reason == nil {
		invalidParams.Add(aws.NewErrParamRequired("Reason"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CancelJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelJobOutput) String() string {
	return awsutil.Prettify(s)
}
