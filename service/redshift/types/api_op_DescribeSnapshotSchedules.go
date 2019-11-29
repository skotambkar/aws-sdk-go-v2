// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSnapshotSchedulesInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the cluster whose snapshot schedules you want to
	// view.
	ClusterIdentifier *string `type:"string"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// The maximum number or response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	MaxRecords *int64 `type:"integer"`

	// A unique identifier for a snapshot schedule.
	ScheduleIdentifier *string `type:"string"`

	// The key value for a snapshot schedule tag.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// The value corresponding to the key of the snapshot schedule tag.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotSchedulesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeSnapshotSchedulesOutput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// A list of SnapshotSchedules.
	SnapshotSchedules []SnapshotSchedule `locationNameList:"SnapshotSchedule" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotSchedulesOutput) String() string {
	return awsutil.Prettify(s)
}