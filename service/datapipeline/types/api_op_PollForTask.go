// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for PollForTask.
type PollForTaskInput struct {
	_ struct{} `type:"structure"`

	// The public DNS name of the calling task runner.
	Hostname *string `locationName:"hostname" min:"1" type:"string"`

	// Identity information for the EC2 instance that is hosting the task runner.
	// You can get this value from the instance using http://169.254.169.254/latest/meta-data/instance-id.
	// For more information, see Instance Metadata (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AESDG-chapter-instancedata.html)
	// in the Amazon Elastic Compute Cloud User Guide. Passing in this value proves
	// that your task runner is running on an EC2 instance, and ensures the proper
	// AWS Data Pipeline service charges are applied to your pipeline.
	InstanceIdentity *InstanceIdentity `locationName:"instanceIdentity" type:"structure"`

	// The type of task the task runner is configured to accept and process. The
	// worker group is set as a field on objects in the pipeline when they are created.
	// You can only specify a single value for workerGroup in the call to PollForTask.
	// There are no wildcard values permitted in workerGroup; the string must be
	// an exact, case-sensitive, match.
	//
	// WorkerGroup is a required field
	WorkerGroup *string `locationName:"workerGroup" type:"string" required:"true"`
}

// String returns the string representation
func (s PollForTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PollForTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PollForTaskInput"}
	if s.Hostname != nil && len(*s.Hostname) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Hostname", 1))
	}

	if s.WorkerGroup == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkerGroup"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of PollForTask.
type PollForTaskOutput struct {
	_ struct{} `type:"structure"`

	// The information needed to complete the task that is being assigned to the
	// task runner. One of the fields returned in this object is taskId, which contains
	// an identifier for the task being assigned. The calling task runner uses taskId
	// in subsequent calls to ReportTaskProgress and SetTaskStatus.
	TaskObject *TaskObject `locationName:"taskObject" type:"structure"`
}

// String returns the string representation
func (s PollForTaskOutput) String() string {
	return awsutil.Prettify(s)
}
