// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

type DescribeVpcAttributeInput struct {
	_ struct{} `type:"structure"`

	// The VPC attribute.
	//
	// Attribute is a required field
	Attribute enums.VpcAttributeName `type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVpcAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVpcAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeVpcAttributeOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the instances launched in the VPC get DNS hostnames. If
	// this attribute is true, instances in the VPC get DNS hostnames; otherwise,
	// they do not.
	EnableDnsHostnames *AttributeBooleanValue `locationName:"enableDnsHostnames" type:"structure"`

	// Indicates whether DNS resolution is enabled for the VPC. If this attribute
	// is true, the Amazon DNS server resolves DNS hostnames for your instances
	// to their corresponding IP addresses; otherwise, it does not.
	EnableDnsSupport *AttributeBooleanValue `locationName:"enableDnsSupport" type:"structure"`

	// The ID of the VPC.
	VpcId *string `locationName:"vpcId" type:"string"`
}

// String returns the string representation
func (s DescribeVpcAttributeOutput) String() string {
	return awsutil.Prettify(s)
}
