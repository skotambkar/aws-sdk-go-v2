// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSubnetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	//
	// The value for MaxResults must be between 20 and 100.
	MaxResults *int64 `type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `type:"string"`

	// The name of the subnet group.
	SubnetGroupNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeSubnetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeSubnetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string `type:"string"`

	// An array of subnet groups. Each element in the array represents a single
	// subnet group.
	SubnetGroups []SubnetGroup `type:"list"`
}

// String returns the string representation
func (s DescribeSubnetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}