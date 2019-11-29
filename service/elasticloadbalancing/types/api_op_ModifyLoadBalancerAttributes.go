// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ModifyLoadBalancerAttributes.
type ModifyLoadBalancerAttributesInput struct {
	_ struct{} `type:"structure"`

	// The attributes for the load balancer.
	//
	// LoadBalancerAttributes is a required field
	LoadBalancerAttributes *LoadBalancerAttributes `type:"structure" required:"true"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyLoadBalancerAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyLoadBalancerAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyLoadBalancerAttributesInput"}

	if s.LoadBalancerAttributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerAttributes"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}
	if s.LoadBalancerAttributes != nil {
		if err := s.LoadBalancerAttributes.Validate(); err != nil {
			invalidParams.AddNested("LoadBalancerAttributes", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of ModifyLoadBalancerAttributes.
type ModifyLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the load balancer attributes.
	LoadBalancerAttributes *LoadBalancerAttributes `type:"structure"`

	// The name of the load balancer.
	LoadBalancerName *string `type:"string"`
}

// String returns the string representation
func (s ModifyLoadBalancerAttributesOutput) String() string {
	return awsutil.Prettify(s)
}