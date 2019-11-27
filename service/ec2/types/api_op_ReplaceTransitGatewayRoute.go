// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ReplaceTransitGatewayRouteInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether traffic matching this route is to be dropped.
	Blackhole *bool `type:"boolean"`

	// The CIDR range used for the destination match. Routing decisions are based
	// on the most specific match.
	//
	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the attachment.
	TransitGatewayAttachmentId *string `type:"string"`

	// The ID of the route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceTransitGatewayRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceTransitGatewayRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReplaceTransitGatewayRouteInput"}

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

type ReplaceTransitGatewayRouteOutput struct {
	_ struct{} `type:"structure"`

	// Information about the modified route.
	Route *TransitGatewayRoute `locationName:"route" type:"structure"`
}

// String returns the string representation
func (s ReplaceTransitGatewayRouteOutput) String() string {
	return awsutil.Prettify(s)
}
