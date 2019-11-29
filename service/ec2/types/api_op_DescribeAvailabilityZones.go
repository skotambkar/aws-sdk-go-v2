// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAvailabilityZonesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * message - Information about the Availability Zone.
	//
	//    * region-name - The name of the Region for the Availability Zone (for
	//    example, us-east-1).
	//
	//    * state - The state of the Availability Zone (available | information
	//    | impaired | unavailable).
	//
	//    * zone-id - The ID of the Availability Zone (for example, use1-az1).
	//
	//    * zone-name - The name of the Availability Zone (for example, us-east-1a).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The IDs of the Availability Zones.
	ZoneIds []string `locationName:"ZoneId" locationNameList:"ZoneId" type:"list"`

	// The names of the Availability Zones.
	ZoneNames []string `locationName:"ZoneName" locationNameList:"ZoneName" type:"list"`
}

// String returns the string representation
func (s DescribeAvailabilityZonesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeAvailabilityZonesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Availability Zones.
	AvailabilityZones []AvailabilityZone `locationName:"availabilityZoneInfo" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeAvailabilityZonesOutput) String() string {
	return awsutil.Prettify(s)
}
