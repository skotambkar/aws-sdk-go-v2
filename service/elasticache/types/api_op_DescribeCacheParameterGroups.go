// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeCacheParameterGroups operation.
type DescribeCacheParameterGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific cache parameter group to return details for.
	CacheParameterGroupName *string `type:"string"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCacheParameterGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeCacheParameterGroups operation.
type DescribeCacheParameterGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of cache parameter groups. Each element in the list contains detailed
	// information about one cache parameter group.
	CacheParameterGroups []CacheParameterGroup `locationNameList:"CacheParameterGroup" type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCacheParameterGroupsOutput) String() string {
	return awsutil.Prettify(s)
}