// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetTransitGatewayRouteTablePropagationsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters. The possible values are:
	//
	//    * resource-id - The ID of the resource.
	//
	//    * resource-type - The resource type (vpc | vpn).
	//
	//    * transit-gateway-attachment-id - The ID of the attachment.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The ID of the transit gateway route table.
	//
	// TransitGatewayRouteTableId is a required field
	TransitGatewayRouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetTransitGatewayRouteTablePropagationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTransitGatewayRouteTablePropagationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTransitGatewayRouteTablePropagationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if s.TransitGatewayRouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTransitGatewayRouteTablePropagationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the route table propagations.
	TransitGatewayRouteTablePropagations []TransitGatewayRouteTablePropagation `locationName:"transitGatewayRouteTablePropagations" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s GetTransitGatewayRouteTablePropagationsOutput) String() string {
	return awsutil.Prettify(s)
}