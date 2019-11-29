// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateJobInput struct {
	_ struct{} `type:"structure"`

	// Allows you to create criteria to abort a job.
	AbortConfig *AbortConfig `locationName:"abortConfig" type:"structure"`

	// A short text description of the job.
	Description *string `locationName:"description" type:"string"`

	// Allows you to create a staged rollout of the job.
	JobExecutionsRolloutConfig *JobExecutionsRolloutConfig `locationName:"jobExecutionsRolloutConfig" type:"structure"`

	// The ID of the job to be updated.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// Configuration information for pre-signed S3 URLs.
	PresignedUrlConfig *PresignedUrlConfig `locationName:"presignedUrlConfig" type:"structure"`

	// Specifies the amount of time each device has to finish its execution of the
	// job. The timer is started when the job execution status is set to IN_PROGRESS.
	// If the job execution status is not set to another terminal state before the
	// time expires, it will be automatically set to TIMED_OUT.
	TimeoutConfig *TimeoutConfig `locationName:"timeoutConfig" type:"structure"`
}

// String returns the string representation
func (s UpdateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if s.AbortConfig != nil {
		if err := s.AbortConfig.Validate(); err != nil {
			invalidParams.AddNested("AbortConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.JobExecutionsRolloutConfig != nil {
		if err := s.JobExecutionsRolloutConfig.Validate(); err != nil {
			invalidParams.AddNested("JobExecutionsRolloutConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.PresignedUrlConfig != nil {
		if err := s.PresignedUrlConfig.Validate(); err != nil {
			invalidParams.AddNested("PresignedUrlConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateJobOutput) String() string {
	return awsutil.Prettify(s)
}