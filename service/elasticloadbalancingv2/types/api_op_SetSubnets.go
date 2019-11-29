// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SetSubnetsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// LoadBalancerArn is a required field
	LoadBalancerArn *string `type:"string" required:"true"`

	// The IDs of the public subnets. You must specify subnets from at least two
	// Availability Zones. You can specify only one subnet per Availability Zone.
	// You must specify either subnets or subnet mappings.
	//
	// You cannot specify Elastic IP addresses for your subnets.
	SubnetMappings []SubnetMapping `type:"list"`

	// The IDs of the public subnets. You must specify subnets from at least two
	// Availability Zones. You can specify only one subnet per Availability Zone.
	// You must specify either subnets or subnet mappings.
	Subnets []string `type:"list"`
}

// String returns the string representation
func (s SetSubnetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetSubnetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetSubnetsInput"}

	if s.LoadBalancerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetSubnetsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the subnet and Availability Zone.
	AvailabilityZones []AvailabilityZone `type:"list"`
}

// String returns the string representation
func (s SetSubnetsOutput) String() string {
	return awsutil.Prettify(s)
}