// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeFleetsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters.
	//
	//    * activity-status - The progress of the EC2 Fleet ( error | pending-fulfillment
	//    | pending-termination | fulfilled).
	//
	//    * excess-capacity-termination-policy - Indicates whether to terminate
	//    running instances if the target capacity is decreased below the current
	//    EC2 Fleet size (true | false).
	//
	//    * fleet-state - The state of the EC2 Fleet (submitted | active | deleted
	//    | failed | deleted-running | deleted-terminating | modifying).
	//
	//    * replace-unhealthy-instances - Indicates whether EC2 Fleet should replace
	//    unhealthy instances (true | false).
	//
	//    * type - The type of request (instant | request | maintain).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The ID of the EC2 Fleets.
	FleetIds []string `locationName:"FleetId" type:"list"`

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeFleetsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeFleetsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the EC2 Fleets.
	Fleets []FleetData `locationName:"fleetSet" locationNameList:"item" type:"list"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeFleetsOutput) String() string {
	return awsutil.Prettify(s)
}
