// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVpcEndpointServiceConfigurationsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * service-name - The name of the service.
	//
	//    * service-id - The ID of the service.
	//
	//    * service-state - The state of the service (Pending | Available | Deleting
	//    | Deleted | Failed).
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

	// The maximum number of results to return for the request in a single page.
	// The remaining results of the initial request can be seen by sending another
	// request with the returned NextToken value. This value can be between 5 and
	// 1000; if MaxResults is given a value larger than 1000, only 1000 results
	// are returned.
	MaxResults *int64 `type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`

	// The IDs of one or more services.
	ServiceIds []string `locationName:"ServiceId" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcEndpointServiceConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeVpcEndpointServiceConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about one or more services.
	ServiceConfigurations []ServiceConfiguration `locationName:"serviceConfigurationSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcEndpointServiceConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}
