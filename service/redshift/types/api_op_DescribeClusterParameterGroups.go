// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeClusterParameterGroupsInput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterParameterGroups request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The name of a specific parameter group for which to return details. By default,
	// details about all parameter groups and the default parameter group are returned.
	ParameterGroupName *string `type:"string"`

	// A tag key or keys for which you want to return all matching cluster parameter
	// groups that are associated with the specified key or keys. For example, suppose
	// that you have parameter groups that are tagged with keys called owner and
	// environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the parameter groups that have either or
	// both of these tag keys associated with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching cluster parameter
	// groups that are associated with the specified tag value or values. For example,
	// suppose that you have parameter groups that are tagged with values called
	// admin and test. If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the parameter groups that have either or
	// both of these tag values associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeClusterParameterGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output from the DescribeClusterParameterGroups action.
type DescribeClusterParameterGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// A list of ClusterParameterGroup instances. Each instance describes one cluster
	// parameter group.
	ParameterGroups []ClusterParameterGroup `locationNameList:"ClusterParameterGroup" type:"list"`
}

// String returns the string representation
func (s DescribeClusterParameterGroupsOutput) String() string {
	return awsutil.Prettify(s)
}
