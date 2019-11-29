// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateAddressInput struct {
	_ struct{} `type:"structure"`

	// [EC2-VPC] The allocation ID. This is required for EC2-VPC.
	AllocationId *string `type:"string"`

	// [EC2-VPC] For a VPC in an EC2-Classic account, specify true to allow an Elastic
	// IP address that is already associated with an instance or network interface
	// to be reassociated with the specified instance or network interface. Otherwise,
	// the operation fails. In a VPC in an EC2-VPC-only account, reassociation is
	// automatic, therefore you can specify false to ensure the operation fails
	// if the Elastic IP address is already associated with another resource.
	AllowReassociation *bool `locationName:"allowReassociation" type:"boolean"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance. This is required for EC2-Classic. For EC2-VPC, you
	// can specify either the instance ID or the network interface ID, but not both.
	// The operation fails if you specify an instance ID unless exactly one network
	// interface is attached.
	InstanceId *string `type:"string"`

	// [EC2-VPC] The ID of the network interface. If the instance has more than
	// one network interface, you must specify a network interface ID.
	//
	// For EC2-VPC, you can specify either the instance ID or the network interface
	// ID, but not both.
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string"`

	// [EC2-VPC] The primary or secondary private IP address to associate with the
	// Elastic IP address. If no private IP address is specified, the Elastic IP
	// address is associated with the primary private IP address.
	PrivateIpAddress *string `locationName:"privateIpAddress" type:"string"`

	// The Elastic IP address to associate with the instance. This is required for
	// EC2-Classic.
	PublicIp *string `type:"string"`
}

// String returns the string representation
func (s AssociateAddressInput) String() string {
	return awsutil.Prettify(s)
}

type AssociateAddressOutput struct {
	_ struct{} `type:"structure"`

	// [EC2-VPC] The ID that represents the association of the Elastic IP address
	// with an instance.
	AssociationId *string `locationName:"associationId" type:"string"`
}

// String returns the string representation
func (s AssociateAddressOutput) String() string {
	return awsutil.Prettify(s)
}