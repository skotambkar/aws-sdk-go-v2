// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssignIpv6AddressesInput struct {
	_ struct{} `type:"structure"`

	// The number of IPv6 addresses to assign to the network interface. Amazon EC2
	// automatically selects the IPv6 addresses from the subnet range. You can't
	// use this option if specifying specific IPv6 addresses.
	Ipv6AddressCount *int64 `locationName:"ipv6AddressCount" type:"integer"`

	// One or more specific IPv6 addresses to be assigned to the network interface.
	// You can't use this option if you're specifying a number of IPv6 addresses.
	Ipv6Addresses []string `locationName:"ipv6Addresses" locationNameList:"item" type:"list"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssignIpv6AddressesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssignIpv6AddressesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssignIpv6AddressesInput"}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssignIpv6AddressesOutput struct {
	_ struct{} `type:"structure"`

	// The IPv6 addresses assigned to the network interface.
	AssignedIpv6Addresses []string `locationName:"assignedIpv6Addresses" locationNameList:"item" type:"list"`

	// The ID of the network interface.
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string"`
}

// String returns the string representation
func (s AssignIpv6AddressesOutput) String() string {
	return awsutil.Prettify(s)
}
