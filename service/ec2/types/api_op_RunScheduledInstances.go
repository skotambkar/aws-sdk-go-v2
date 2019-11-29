// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for RunScheduledInstances.
type RunScheduledInstancesInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that ensures the idempotency of the request.
	// For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The number of instances.
	//
	// Default: 1
	InstanceCount *int64 `type:"integer"`

	// The launch specification. You must match the instance type, Availability
	// Zone, network, and platform of the schedule that you purchased.
	//
	// LaunchSpecification is a required field
	LaunchSpecification *ScheduledInstancesLaunchSpecification `type:"structure" required:"true"`

	// The Scheduled Instance ID.
	//
	// ScheduledInstanceId is a required field
	ScheduledInstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RunScheduledInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RunScheduledInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RunScheduledInstancesInput"}

	if s.LaunchSpecification == nil {
		invalidParams.Add(aws.NewErrParamRequired("LaunchSpecification"))
	}

	if s.ScheduledInstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledInstanceId"))
	}
	if s.LaunchSpecification != nil {
		if err := s.LaunchSpecification.Validate(); err != nil {
			invalidParams.AddNested("LaunchSpecification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of RunScheduledInstances.
type RunScheduledInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the newly launched instances.
	InstanceIdSet []string `locationName:"instanceIdSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s RunScheduledInstancesOutput) String() string {
	return awsutil.Prettify(s)
}
