// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ssm/enums"
)

type GetMaintenanceWindowExecutionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the maintenance window execution that includes the task.
	//
	// WindowExecutionId is a required field
	WindowExecutionId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s GetMaintenanceWindowExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMaintenanceWindowExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMaintenanceWindowExecutionInput"}

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

type GetMaintenanceWindowExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The time the maintenance window finished running.
	EndTime *time.Time `type:"timestamp"`

	// The time the maintenance window started running.
	StartTime *time.Time `type:"timestamp"`

	// The status of the maintenance window execution.
	Status enums.MaintenanceWindowExecutionStatus `type:"string" enum:"true"`

	// The details explaining the Status. Only available for certain status values.
	StatusDetails *string `type:"string"`

	// The ID of the task executions from the maintenance window execution.
	TaskIds []string `type:"list"`

	// The ID of the maintenance window execution.
	WindowExecutionId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s GetMaintenanceWindowExecutionOutput) String() string {
	return awsutil.Prettify(s)
}