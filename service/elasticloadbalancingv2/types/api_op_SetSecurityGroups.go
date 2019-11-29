// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SetSecurityGroupsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the load balancer.
	//
	// LoadBalancerArn is a required field
	LoadBalancerArn *string `type:"string" required:"true"`

	// The IDs of the security groups.
	//
	// SecurityGroups is a required field
	SecurityGroups []string `type:"list" required:"true"`
}

// String returns the string representation
func (s SetSecurityGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetSecurityGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetSecurityGroupsInput"}

	if s.LoadBalancerArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerArn"))
	}

	if s.SecurityGroups == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityGroups"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetSecurityGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the security groups associated with the load balancer.
	SecurityGroupIds []string `type:"list"`
}

// String returns the string representation
func (s SetSecurityGroupsOutput) String() string {
	return awsutil.Prettify(s)
}
