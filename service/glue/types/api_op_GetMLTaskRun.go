// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/glue/enums"
)

type GetMLTaskRunInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the task run.
	//
	// TaskRunId is a required field
	TaskRunId *string `min:"1" type:"string" required:"true"`

	// The unique identifier of the machine learning transform.
	//
	// TransformId is a required field
	TransformId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMLTaskRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMLTaskRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMLTaskRunInput"}

	if s.TaskRunId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskRunId"))
	}
	if s.TaskRunId != nil && len(*s.TaskRunId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskRunId", 1))
	}

	if s.TransformId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransformId"))
	}
	if s.TransformId != nil && len(*s.TransformId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TransformId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMLTaskRunOutput struct {
	_ struct{} `type:"structure"`

	// The date and time when this task run was completed.
	CompletedOn *time.Time `type:"timestamp"`

	// The error strings that are associated with the task run.
	ErrorString *string `type:"string"`

	// The amount of time (in seconds) that the task run consumed resources.
	ExecutionTime *int64 `type:"integer"`

	// The date and time when this task run was last modified.
	LastModifiedOn *time.Time `type:"timestamp"`

	// The names of the log groups that are associated with the task run.
	LogGroupName *string `type:"string"`

	// The list of properties that are associated with the task run.
	Properties *TaskRunProperties `type:"structure"`

	// The date and time when this task run started.
	StartedOn *time.Time `type:"timestamp"`

	// The status for this task run.
	Status enums.TaskStatusType `type:"string" enum:"true"`

	// The unique run identifier associated with this run.
	TaskRunId *string `min:"1" type:"string"`

	// The unique identifier of the task run.
	TransformId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetMLTaskRunOutput) String() string {
	return awsutil.Prettify(s)
}