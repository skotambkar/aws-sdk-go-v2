// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTransitGatewayInput struct {
	_ struct{} `type:"structure"`

	// A description of the transit gateway.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The transit gateway options.
	Options *TransitGatewayRequestOptions `type:"structure"`

	// The tags to apply to the transit gateway.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateTransitGatewayInput) String() string {
	return awsutil.Prettify(s)
}

type CreateTransitGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Information about the transit gateway.
	TransitGateway *TransitGateway `locationName:"transitGateway" type:"structure"`
}

// String returns the string representation
func (s CreateTransitGatewayOutput) String() string {
	return awsutil.Prettify(s)
}