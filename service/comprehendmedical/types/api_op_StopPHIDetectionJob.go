// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopPHIDetectionJobInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the PHI detection job to stop.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopPHIDetectionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopPHIDetectionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopPHIDetectionJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopPHIDetectionJobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the PHI detection job that was stopped.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StopPHIDetectionJobOutput) String() string {
	return awsutil.Prettify(s)
}
