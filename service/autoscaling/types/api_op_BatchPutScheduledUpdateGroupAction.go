// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchPutScheduledUpdateGroupActionInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// One or more scheduled actions. The maximum number allowed is 50.
	//
	// ScheduledUpdateGroupActions is a required field
	ScheduledUpdateGroupActions []ScheduledUpdateGroupActionRequest `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchPutScheduledUpdateGroupActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchPutScheduledUpdateGroupActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchPutScheduledUpdateGroupActionInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.ScheduledUpdateGroupActions == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledUpdateGroupActions"))
	}
	if s.ScheduledUpdateGroupActions != nil {
		for i, v := range s.ScheduledUpdateGroupActions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ScheduledUpdateGroupActions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchPutScheduledUpdateGroupActionOutput struct {
	_ struct{} `type:"structure"`

	// The names of the scheduled actions that could not be created or updated,
	// including an error message.
	FailedScheduledUpdateGroupActions []FailedScheduledUpdateGroupActionRequest `type:"list"`
}

// String returns the string representation
func (s BatchPutScheduledUpdateGroupActionOutput) String() string {
	return awsutil.Prettify(s)
}
