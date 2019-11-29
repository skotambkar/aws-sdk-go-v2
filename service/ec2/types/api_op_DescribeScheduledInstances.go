// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeScheduledInstances.
type DescribeScheduledInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters.
	//
	//    * availability-zone - The Availability Zone (for example, us-west-2a).
	//
	//    * instance-type - The instance type (for example, c4.large).
	//
	//    * network-platform - The network platform (EC2-Classic or EC2-VPC).
	//
	//    * platform - The platform (Linux/UNIX or Windows).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return in a single call. This value can
	// be between 5 and 300. The default value is 100. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of results.
	NextToken *string `type:"string"`

	// The Scheduled Instance IDs.
	ScheduledInstanceIds []string `locationName:"ScheduledInstanceId" locationNameList:"ScheduledInstanceId" type:"list"`

	// The time period for the first schedule to start.
	SlotStartTimeRange *SlotStartTimeRangeRequest `type:"structure"`
}

// String returns the string representation
func (s DescribeScheduledInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeScheduledInstances.
type DescribeScheduledInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The token required to retrieve the next set of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the Scheduled Instances.
	ScheduledInstanceSet []ScheduledInstance `locationName:"scheduledInstanceSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeScheduledInstancesOutput) String() string {
	return awsutil.Prettify(s)
}