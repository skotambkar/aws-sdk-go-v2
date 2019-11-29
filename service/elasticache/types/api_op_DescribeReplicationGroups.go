// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeReplicationGroups operation.
type DescribeReplicationGroupsInput struct {
	_ struct{} `type:"structure"`

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

	// The identifier for the replication group to be described. This parameter
	// is not case sensitive.
	//
	// If you do not specify this parameter, information about all replication groups
	// is returned.
	ReplicationGroupId *string `type:"string"`
}

// String returns the string representation
func (s DescribeReplicationGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeReplicationGroups operation.
type DescribeReplicationGroupsOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`

	// A list of replication groups. Each item in the list contains detailed information
	// about one replication group.
	ReplicationGroups []ReplicationGroup `locationNameList:"ReplicationGroup" type:"list"`
}

// String returns the string representation
func (s DescribeReplicationGroupsOutput) String() string {
	return awsutil.Prettify(s)
}