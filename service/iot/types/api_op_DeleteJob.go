// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteJobInput struct {
	_ struct{} `type:"structure"`

	// (Optional) When true, you can delete a job which is "IN_PROGRESS". Otherwise,
	// you can only delete a job which is in a terminal state ("COMPLETED" or "CANCELED")
	// or an exception will occur. The default is false.
	//
	// Deleting a job which is "IN_PROGRESS", will cause a device which is executing
	// the job to be unable to access job information or update the job execution
	// status. Use caution and ensure that each device executing a job which is
	// deleted is able to recover to a valid state.
	Force *bool `location:"querystring" locationName:"force" type:"boolean"`

	// The ID of the job to be deleted.
	//
	// After a job deletion is completed, you may reuse this jobId when you create
	// a new job. However, this is not recommended, and you must ensure that your
	// devices are not using the jobId to refer to the deleted job.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteJobInput"}

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

type DeleteJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteJobOutput) String() string {
	return awsutil.Prettify(s)
}
