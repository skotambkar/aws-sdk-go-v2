// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeSnapshotsMessage operation.
type DescribeSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// A user-supplied cluster identifier. If this parameter is specified, only
	// snapshots associated with that specific cluster are described.
	CacheClusterId *string `type:"string"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 50
	//
	// Constraints: minimum 20; maximum 50.
	MaxRecords *int64 `type:"integer"`

	// A user-supplied replication group identifier. If this parameter is specified,
	// only snapshots associated with that specific replication group are described.
	ReplicationGroupId *string `type:"string"`

	// A Boolean value which if true, the node group (shard) configuration is included
	// in the snapshot description.
	ShowNodeGroupConfig *bool `type:"boolean"`

	// A user-supplied name of the snapshot. If this parameter is specified, only
	// this snapshot are described.
	SnapshotName *string `type:"string"`

	// If set to system, the output shows snapshots that were automatically created
	// by ElastiCache. If set to user the output shows snapshots that were manually
	// created. If omitted, the output shows both automatically and manually created
	// snapshots.
	SnapshotSource *string `type:"string"`
}

// String returns the string representation
func (s DescribeSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeSnapshots operation.
type DescribeSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of snapshots. Each item in the list contains detailed information
	// about one snapshot.
	Snapshots []Snapshot `locationNameList:"Snapshot" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}
