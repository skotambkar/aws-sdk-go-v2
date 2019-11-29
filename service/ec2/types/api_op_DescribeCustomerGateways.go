// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeCustomerGateways.
type DescribeCustomerGatewaysInput struct {
	_ struct{} `type:"structure"`

	// One or more customer gateway IDs.
	//
	// Default: Describes all your customer gateways.
	CustomerGatewayIds []string `locationName:"CustomerGatewayId" locationNameList:"CustomerGatewayId" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// One or more filters.
	//
	//    * bgp-asn - The customer gateway's Border Gateway Protocol (BGP) Autonomous
	//    System Number (ASN).
	//
	//    * customer-gateway-id - The ID of the customer gateway.
	//
	//    * ip-address - The IP address of the customer gateway's Internet-routable
	//    external interface.
	//
	//    * state - The state of the customer gateway (pending | available | deleting
	//    | deleted).
	//
	//    * type - The type of customer gateway. Currently, the only supported type
	//    is ipsec.1.
	//
	//    * tag:<key> - The key/value combination of a tag assigned to the resource.
	//    Use the tag key in the filter name and the tag value as the filter value.
	//    For example, to find all resources that have a tag with the key Owner
	//    and the value TeamA, specify tag:Owner for the filter name and TeamA for
	//    the filter value.
	//
	//    * tag-key - The key of a tag assigned to the resource. Use this filter
	//    to find all resources assigned a tag with a specific key, regardless of
	//    the tag value.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`
}

// String returns the string representation
func (s DescribeCustomerGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeCustomerGateways.
type DescribeCustomerGatewaysOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more customer gateways.
	CustomerGateways []CustomerGateway `locationName:"customerGatewaySet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeCustomerGatewaysOutput) String() string {
	return awsutil.Prettify(s)
}