// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a PutJobFailureResult action.
type PutJobFailureResultInput struct {
	_ struct{} `type:"structure"`

	// The details about the failure of a job.
	//
	// FailureDetails is a required field
	FailureDetails *FailureDetails `locationName:"failureDetails" type:"structure" required:"true"`

	// The unique system-generated ID of the job that failed. This is the same ID
	// returned from PollForJobs.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" type:"string" required:"true"`
}

// String returns the string representation
func (s PutJobFailureResultInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutJobFailureResultInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutJobFailureResultInput"}

	if s.FailureDetails == nil {
		invalidParams.Add(aws.NewErrParamRequired("FailureDetails"))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.FailureDetails != nil {
		if err := s.FailureDetails.Validate(); err != nil {
			invalidParams.AddNested("FailureDetails", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutJobFailureResultOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutJobFailureResultOutput) String() string {
	return awsutil.Prettify(s)
}
