// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTransitGatewayRouteInput struct {
	_ struct{} `type:"structure"`

	// The CIDR range for the route. This must match the CIDR for the route exactly.
	//
	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitGatewayRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitGatewayRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTransitGatewayRouteInput"}

	if s.DestinationCidrBlock == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationCidrBlock"))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTransitGatewayRouteOutput struct {
	_ struct{} `type:"structure"`

	// Information about the route.
	Route *TransitGatewayRoute `locationName:"route" type:"structure"`
}

// String returns the string representation
func (s DeleteTransitGatewayRouteOutput) String() string {
	return awsutil.Prettify(s)
}