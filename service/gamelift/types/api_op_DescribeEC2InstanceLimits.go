// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/enums"
)

// Represents the input for a request action.
type DescribeEC2InstanceLimitsInput struct {
	_ struct{} `type:"structure"`

	// Name of an EC2 instance type that is supported in Amazon GameLift. A fleet
	// instance type determines the computing resources of each instance in the
	// fleet, including CPU, memory, storage, and networking capacity. Amazon GameLift
	// supports the following EC2 instance types. See Amazon EC2 Instance Types
	// (http://aws.amazon.com/ec2/instance-types/) for detailed descriptions. Leave
	// this parameter blank to retrieve limits for all types.
	EC2InstanceType enums.EC2InstanceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeEC2InstanceLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the returned data in response to a request action.
type DescribeEC2InstanceLimitsOutput struct {
	_ struct{} `type:"structure"`

	// Object that contains the maximum number of instances for the specified instance
	// type.
	EC2InstanceLimits []EC2InstanceLimit `type:"list"`
}

// String returns the string representation
func (s DescribeEC2InstanceLimitsOutput) String() string {
	return awsutil.Prettify(s)
}