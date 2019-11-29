// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ssm/enums"
)

type GetMaintenanceWindowExecutionTaskInput struct {
	_ struct{} `type:"structure"`

	// The ID of the specific task execution in the maintenance window task that
	// should be retrieved.
	//
	// TaskId is a required field
	TaskId *string `min:"36" type:"string" required:"true"`

	// The ID of the maintenance window execution that includes the task.
	//
	// WindowExecutionId is a required field
	WindowExecutionId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMaintenanceWindowExecutionTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMaintenanceWindowExecutionTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMaintenanceWindowExecutionTaskInput"}

	if s.TaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskId"))
	}
	if s.TaskId != nil && len(*s.TaskId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskId", 36))
	}

	if s.WindowExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowExecutionId"))
	}
	if s.WindowExecutionId != nil && len(*s.WindowExecutionId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowExecutionId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMaintenanceWindowExecutionTaskOutput struct {
	_ struct{} `type:"structure"`

	// The time the task execution completed.
	EndTime *time.Time `type:"timestamp"`

	// The defined maximum number of task executions that could be run in parallel.
	MaxConcurrency *string `min:"1" type:"string"`

	// The defined maximum number of task execution errors allowed before scheduling
	// of the task execution would have been stopped.
	MaxErrors *string `min:"1" type:"string"`

	// The priority of the task.
	Priority *int64 `type:"integer"`

	// The role that was assumed when running the task.
	ServiceRole *string `type:"string"`

	// The time the task execution started.
	StartTime *time.Time `type:"timestamp"`

	// The status of the task.
	Status enums.MaintenanceWindowExecutionStatus `type:"string" enum:"true"`

	// The details explaining the Status. Only available for certain status values.
	StatusDetails *string `type:"string"`

	// The ARN of the task that ran.
	TaskArn *string `min:"1" type:"string"`

	// The ID of the specific task execution in the maintenance window task that
	// was retrieved.
	TaskExecutionId *string `min:"36" type:"string"`

	// The parameters passed to the task when it was run.
	//
	// TaskParameters has been deprecated. To specify parameters to pass to a task
	// when it runs, instead use the Parameters option in the TaskInvocationParameters
	// structure. For information about how Systems Manager handles these options
	// for the supported maintenance window task types, see MaintenanceWindowTaskInvocationParameters.
	//
	// The map has the following format:
	//
	// Key: string, between 1 and 255 characters
	//
	// Value: an array of strings, each string is between 1 and 255 characters
	TaskParameters []map[string]MaintenanceWindowTaskParameterValueExpression `type:"list" sensitive:"true"`

	// The type of task that was run.
	Type enums.MaintenanceWindowTaskType `type:"string" enum:"true"`

	// The ID of the maintenance window execution that includes the task.
	WindowExecutionId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s GetMaintenanceWindowExecutionTaskOutput) String() string {
	return awsutil.Prettify(s)
}
