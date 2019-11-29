// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAggregateIdFormatInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s DescribeAggregateIdFormatInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeAggregateIdFormatOutput struct {
	_ struct{} `type:"structure"`

	// Information about each resource's ID format.
	Statuses []IdFormat `locationName:"statusSet" locationNameList:"item" type:"list"`

	// Indicates whether all resource types in the Region are configured to use
	// longer IDs. This value is only true if all users are configured to use longer
	// IDs for all resources types in the Region.
	UseLongIdsAggregated *bool `locationName:"useLongIdsAggregated" type:"boolean"`
}

// String returns the string representation
func (s DescribeAggregateIdFormatOutput) String() string {
	return awsutil.Prettify(s)
}
