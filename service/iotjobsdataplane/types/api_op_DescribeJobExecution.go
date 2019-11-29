// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// Optional. A number that identifies a particular job execution on a particular
	// device. If not specified, the latest job execution is returned.
	ExecutionNumber *int64 `location:"querystring" locationName:"executionNumber" type:"long"`

	// Optional. When set to true, the response contains the job document. The default
	// is false.
	IncludeJobDocument *bool `location:"querystring" locationName:"includeJobDocument" type:"boolean"`

	// The unique identifier assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`

	// The thing name associated with the device the job execution is running on.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeJobExecutionInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

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

type DescribeJobExecutionOutput struct {
	_ struct{} `type:"structure"`

	// Contains data about a job execution.
	Execution *JobExecution `locationName:"execution" type:"structure"`
}

// String returns the string representation
func (s DescribeJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}