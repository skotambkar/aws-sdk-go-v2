// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeCacheSubnetGroups operation.
type DescribeCacheSubnetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache subnet group to return details for.
	CacheSubnetGroupName *string `type:"string"`

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
func (s DescribeCacheSubnetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeCacheSubnetGroups operation.
type DescribeCacheSubnetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of cache subnet groups. Each element in the list contains detailed
	// information about one group.
	CacheSubnetGroups []CacheSubnetGroup `locationNameList:"CacheSubnetGroup" type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCacheSubnetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}