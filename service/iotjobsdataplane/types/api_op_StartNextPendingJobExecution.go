// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartNextPendingJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// A collection of name/value pairs that describe the status of the job execution.
	// If not specified, the statusDetails are unchanged.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// Specifies the amount of time this device has to finish execution of this
	// job. If the job execution status is not set to a terminal state before this
	// timer expires, or before the timer is reset (by calling UpdateJobExecution,
	// setting the status to IN_PROGRESS and specifying a new timeout value in field
	// stepTimeoutInMinutes) the job execution status will be automatically set
	// to TIMED_OUT. Note that setting this timeout has no effect on that job execution
	// timeout which may have been specified when the job was created (CreateJob
	// using field timeoutConfig).
	StepTimeoutInMinutes *int64 `locationName:"stepTimeoutInMinutes" type:"long"`

	// The name of the thing associated with the device.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartNextPendingJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartNextPendingJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartNextPendingJobExecutionInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartNextPendingJobExecutionOutput struct {
	_ struct{} `type:"structure"`

	// A JobExecution object.
	Execution *JobExecution `locationName:"execution" type:"structure"`
}

// String returns the string representation
func (s StartNextPendingJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}
