// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

// Contains the parameters for DescribeSpotPriceHistory.
type DescribeSpotPriceHistoryInput struct {
	_ struct{} `type:"structure"`

	// Filters the results by the specified Availability Zone.
	AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The date and time, up to the current date, from which to stop retrieving
	// the price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// One or more filters.
	//
	//    * availability-zone - The Availability Zone for which prices should be
	//    returned.
	//
	//    * instance-type - The type of instance (for example, m3.medium).
	//
	//    * product-description - The product description for the Spot price (Linux/UNIX
	//    | SUSE Linux | Windows | Linux/UNIX (Amazon VPC) | SUSE Linux (Amazon
	//    VPC) | Windows (Amazon VPC)).
	//
	//    * spot-price - The Spot price. The value must match exactly (or use wildcards;
	//    greater than or less than comparison is not supported).
	//
	//    * timestamp - The time stamp of the Spot price history, in UTC format
	//    (for example, YYYY-MM-DDTHH:MM:SSZ). You can use wildcards (* and ?).
	//    Greater than or less than comparison is not supported.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// Filters the results by the specified instance types.
	InstanceTypes []enums.InstanceType `locationName:"InstanceType" type:"list"`

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Filters the results by the specified basic product descriptions.
	ProductDescriptions []string `locationName:"ProductDescription" type:"list"`

	// The date and time, up to the past 90 days, from which to start retrieving
	// the price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`
}

// String returns the string representation
func (s DescribeSpotPriceHistoryInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output of DescribeSpotPriceHistory.
type DescribeSpotPriceHistoryOutput struct {
	_ struct{} `type:"structure"`

	// The token required to retrieve the next set of results. This value is null
	// or an empty string when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The historical Spot prices.
	SpotPriceHistory []SpotPrice `locationName:"spotPriceHistorySet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeSpotPriceHistoryOutput) String() string {
	return awsutil.Prettify(s)
}