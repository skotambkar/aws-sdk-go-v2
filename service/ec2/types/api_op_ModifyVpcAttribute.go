// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyVpcAttributeInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the instances launched in the VPC get DNS hostnames. If
	// enabled, instances in the VPC get DNS hostnames; otherwise, they do not.
	//
	// You cannot modify the DNS resolution and DNS hostnames attributes in the
	// same request. Use separate requests for each attribute. You can only enable
	// DNS hostnames if you've enabled DNS support.
	EnableDnsHostnames *AttributeBooleanValue `type:"structure"`

	// Indicates whether the DNS resolution is supported for the VPC. If enabled,
	// queries to the Amazon provided DNS server at the 169.254.169.253 IP address,
	// or the reserved IP address at the base of the VPC network range "plus two"
	// succeed. If disabled, the Amazon provided DNS service in the VPC that resolves
	// public DNS hostnames to IP addresses is not enabled.
	//
	// You cannot modify the DNS resolution and DNS hostnames attributes in the
	// same request. Use separate requests for each attribute.
	EnableDnsSupport *AttributeBooleanValue `type:"structure"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `locationName:"vpcId" type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpcAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpcAttributeInput"}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpcAttributeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyVpcAttributeOutput) String() string {
	return awsutil.Prettify(s)
}
