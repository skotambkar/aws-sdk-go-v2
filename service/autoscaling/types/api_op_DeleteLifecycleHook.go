// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteLifecycleHookInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// The name of the lifecycle hook.
	//
	// LifecycleHookName is a required field
	LifecycleHookName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLifecycleHookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLifecycleHookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLifecycleHookInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.LifecycleHookName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LifecycleHookName"))
	}
	if s.LifecycleHookName != nil && len(*s.LifecycleHookName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LifecycleHookName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteLifecycleHookOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLifecycleHookOutput) String() string {
	return awsutil.Prettify(s)
}