// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteRouteInput struct {
	_ struct{} `type:"structure"`

	// The IPv4 CIDR range for the route. The value you specify must match the CIDR
	// for the route exactly.
	DestinationCidrBlock *string `locationName:"destinationCidrBlock" type:"string"`

	// The IPv6 CIDR range for the route. The value you specify must match the CIDR
	// for the route exactly.
	DestinationIpv6CidrBlock *string `locationName:"destinationIpv6CidrBlock" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the route table.
	//
	// RouteTableId is a required field
	RouteTableId *string `locationName:"routeTableId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRouteInput"}

	if s.RouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRouteOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRouteOutput) String() string {
	return awsutil.Prettify(s)
}
