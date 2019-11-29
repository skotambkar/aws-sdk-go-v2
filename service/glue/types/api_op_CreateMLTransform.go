// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/glue/enums"
)

type CreateMLTransformInput struct {
	_ struct{} `type:"structure"`

	// A description of the machine learning transform that is being defined. The
	// default is an empty string.
	Description *string `type:"string"`

	// A list of AWS Glue table definitions used by the transform.
	//
	// InputRecordTables is a required field
	InputRecordTables []GlueTable `type:"list" required:"true"`

	// The number of AWS Glue data processing units (DPUs) that are allocated to
	// task runs for this transform. You can allocate from 2 to 100 DPUs; the default
	// is 10. A DPU is a relative measure of processing power that consists of 4
	// vCPUs of compute capacity and 16 GB of memory. For more information, see
	// the AWS Glue pricing page (https://aws.amazon.com/glue/pricing/).
	//
	// When the WorkerType field is set to a value other than Standard, the MaxCapacity
	// field is set automatically and becomes read-only.
	MaxCapacity *float64 `type:"double"`

	// The maximum number of times to retry a task for this transform after a task
	// run fails.
	MaxRetries *int64 `type:"integer"`

	// The unique name that you give the transform when you create it.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The number of workers of a defined workerType that are allocated when this
	// task runs.
	NumberOfWorkers *int64 `type:"integer"`

	// The algorithmic parameters that are specific to the transform type used.
	// Conditionally dependent on the transform type.
	//
	// Parameters is a required field
	Parameters *TransformParameters `type:"structure" required:"true"`

	// The name or Amazon Resource Name (ARN) of the IAM role with the required
	// permissions. Ensure that this role has permission to your Amazon Simple Storage
	// Service (Amazon S3) sources, targets, temporary directory, scripts, and any
	// libraries that are used by the task run for this transform.
	//
	// Role is a required field
	Role *string `type:"string" required:"true"`

	// The timeout of the task run for this transform in minutes. This is the maximum
	// time that a task run for this transform can consume resources before it is
	// terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours).
	Timeout *int64 `min:"1" type:"integer"`

	// The type of predefined worker that is allocated when this task runs. Accepts
	// a value of Standard, G.1X, or G.2X.
	//
	//    * For the Standard worker type, each worker provides 4 vCPU, 16 GB of
	//    memory and a 50GB disk, and 2 executors per worker.
	//
	//    * For the G.1X worker type, each worker provides 4 vCPU, 16 GB of memory
	//    and a 64GB disk, and 1 executor per worker.
	//
	//    * For the G.2X worker type, each worker provides 8 vCPU, 32 GB of memory
	//    and a 128GB disk, and 1 executor per worker.
	WorkerType enums.WorkerType `type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateMLTransformInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMLTransformInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMLTransformInput"}

	if s.InputRecordTables == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputRecordTables"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Parameters == nil {
		invalidParams.Add(aws.NewErrParamRequired("Parameters"))
	}

	if s.Role == nil {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}
	if s.Timeout != nil && *s.Timeout < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Timeout", 1))
	}
	if s.InputRecordTables != nil {
		for i, v := range s.InputRecordTables {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "InputRecordTables", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			invalidParams.AddNested("Parameters", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateMLTransformOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier that is generated for the transform.
	TransformId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateMLTransformOutput) String() string {
	return awsutil.Prettify(s)
}
