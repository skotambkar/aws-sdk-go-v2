// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UnassignIpv6AddressesInput struct {
	_ struct{} `type:"structure"`

	// The IPv6 addresses to unassign from the network interface.
	//
	// Ipv6Addresses is a required field
	Ipv6Addresses []string `locationName:"ipv6Addresses" locationNameList:"item" type:"list" required:"true"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s UnassignIpv6AddressesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnassignIpv6AddressesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UnassignIpv6AddressesInput"}

	if s.Ipv6Addresses == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ipv6Addresses"))
	}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UnassignIpv6AddressesOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the network interface.
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string"`

	// The IPv6 addresses that have been unassigned from the network interface.
	UnassignedIpv6Addresses []string `locationName:"unassignedIpv6Addresses" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s UnassignIpv6AddressesOutput) String() string {
	return awsutil.Prettify(s)
}