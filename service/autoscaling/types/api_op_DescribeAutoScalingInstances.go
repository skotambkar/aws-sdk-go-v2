// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAutoScalingInstancesInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the instances. You can specify up to MaxRecords IDs. If you omit
	// this parameter, all Auto Scaling instances are described. If you specify
	// an ID that does not exist, it is ignored with no error.
	InstanceIds []string `type:"list"`

	// The maximum number of items to return with this call. The default value is
	// 50 and the maximum value is 50.
	MaxRecords *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAutoScalingInstancesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeAutoScalingInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The instances.
	AutoScalingInstances []AutoScalingInstanceDetails `type:"list"`

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this
	// string for the NextToken value when requesting the next set of items. This
	// value is null when there are no more items to return.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeAutoScalingInstancesOutput) String() string {
	return awsutil.Prettify(s)
}
