// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutAutoScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	// Specifies the definition of the automatic scaling policy.
	//
	// AutoScalingPolicy is a required field
	AutoScalingPolicy *AutoScalingPolicy `type:"structure" required:"true"`

	// Specifies the ID of a cluster. The instance group to which the automatic
	// scaling policy is applied is within this cluster.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`

	// Specifies the ID of the instance group to which the automatic scaling policy
	// is applied.
	//
	// InstanceGroupId is a required field
	InstanceGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutAutoScalingPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutAutoScalingPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutAutoScalingPolicyInput"}

	if s.AutoScalingPolicy == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingPolicy"))
	}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if s.InstanceGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceGroupId"))
	}
	if s.AutoScalingPolicy != nil {
		if err := s.AutoScalingPolicy.Validate(); err != nil {
			invalidParams.AddNested("AutoScalingPolicy", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutAutoScalingPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The automatic scaling policy definition.
	AutoScalingPolicy *AutoScalingPolicyDescription `type:"structure"`

	// The Amazon Resource Name of the cluster.
	ClusterArn *string `min:"20" type:"string"`

	// Specifies the ID of a cluster. The instance group to which the automatic
	// scaling policy is applied is within this cluster.
	ClusterId *string `type:"string"`

	// Specifies the ID of the instance group to which the scaling policy is applied.
	InstanceGroupId *string `type:"string"`
}

// String returns the string representation
func (s PutAutoScalingPolicyOutput) String() string {
	return awsutil.Prettify(s)
}
