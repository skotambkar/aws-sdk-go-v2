// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAutoScalingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// Specifies that the group is to be deleted along with all instances associated
	// with the group, without waiting for all instances to be terminated. This
	// parameter also deletes any lifecycle actions associated with the group.
	ForceDelete *bool `type:"boolean"`
}

// String returns the string representation
func (s DeleteAutoScalingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAutoScalingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAutoScalingGroupInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAutoScalingGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAutoScalingGroupOutput) String() string {
	return awsutil.Prettify(s)
}