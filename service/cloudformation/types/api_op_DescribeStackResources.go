// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for DescribeStackResources action.
type DescribeStackResourcesInput struct {
	_ struct{} `type:"structure"`

	// The logical name of the resource as specified in the template.
	//
	// Default: There is no default value.
	LogicalResourceId *string `type:"string"`

	// The name or unique identifier that corresponds to a physical instance ID
	// of a resource supported by AWS CloudFormation.
	//
	// For example, for an Amazon Elastic Compute Cloud (EC2) instance, PhysicalResourceId
	// corresponds to the InstanceId. You can pass the EC2 InstanceId to DescribeStackResources
	// to find which stack the instance belongs to and what other resources are
	// part of the stack.
	//
	// Required: Conditional. If you do not specify PhysicalResourceId, you must
	// specify StackName.
	//
	// Default: There is no default value.
	PhysicalResourceId *string `type:"string"`

	// The name or the unique stack ID that is associated with the stack, which
	// are not always interchangeable:
	//
	//    * Running stacks: You can specify either the stack's name or its unique
	//    stack ID.
	//
	//    * Deleted stacks: You must specify the unique stack ID.
	//
	// Default: There is no default value.
	//
	// Required: Conditional. If you do not specify StackName, you must specify
	// PhysicalResourceId.
	StackName *string `type:"string"`
}

// String returns the string representation
func (s DescribeStackResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// The output for a DescribeStackResources action.
type DescribeStackResourcesOutput struct {
	_ struct{} `type:"structure"`

	// A list of StackResource structures.
	StackResources []StackResource `type:"list"`
}

// String returns the string representation
func (s DescribeStackResourcesOutput) String() string {
	return awsutil.Prettify(s)
}
