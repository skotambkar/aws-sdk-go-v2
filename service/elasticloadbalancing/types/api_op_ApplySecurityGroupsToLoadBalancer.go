// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for ApplySecurityGroupsToLoadBalancer.
type ApplySecurityGroupsToLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`

	// The IDs of the security groups to associate with the load balancer. Note
	// that you cannot specify the name of the security group.
	//
	// SecurityGroups is a required field
	SecurityGroups []string `type:"list" required:"true"`
}

// String returns the string representation
func (s ApplySecurityGroupsToLoadBalancerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplySecurityGroupsToLoadBalancerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplySecurityGroupsToLoadBalancerInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if s.SecurityGroups == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityGroups"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of ApplySecurityGroupsToLoadBalancer.
type ApplySecurityGroupsToLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the security groups associated with the load balancer.
	SecurityGroups []string `type:"list"`
}

// String returns the string representation
func (s ApplySecurityGroupsToLoadBalancerOutput) String() string {
	return awsutil.Prettify(s)
}
