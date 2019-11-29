// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSubnetInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone for the subnet.
	//
	// Default: AWS selects one for you. If you create more than one subnet in your
	// VPC, we may not necessarily select a different zone for each subnet.
	AvailabilityZone *string `type:"string"`

	// The AZ ID of the subnet.
	AvailabilityZoneId *string `type:"string"`

	// The IPv4 network range for the subnet, in CIDR notation. For example, 10.0.0.0/24.
	//
	// CidrBlock is a required field
	CidrBlock *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The IPv6 network range for the subnet, in CIDR notation. The subnet size
	// must use a /64 prefix length.
	Ipv6CidrBlock *string `type:"string"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSubnetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSubnetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSubnetInput"}

	if s.CidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("CidrBlock"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSubnetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the subnet.
	Subnet *Subnet `locationName:"subnet" type:"structure"`
}

// String returns the string representation
func (s CreateSubnetOutput) String() string {
	return awsutil.Prettify(s)
}
