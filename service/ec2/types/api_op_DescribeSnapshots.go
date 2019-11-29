// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * description - A description of the snapshot.
	//
	//    * encrypted - Indicates whether the snapshot is encrypted (true | false)
	//
	//    * owner-alias - Value from an Amazon-maintained list (amazon | self |
	//    all | aws-marketplace | microsoft) of snapshot owners. Not to be confused
	//    with the user-configured AWS account alias, which is set from the IAM
	//    console.
	//
	//    * owner-id - The ID of the AWS account that owns the snapshot.
	//
	//    * progress - The progress of the snapshot, as a percentage (for example,
	//    80%).
	//
	//    * snapshot-id - The snapshot ID.
	//
	//    * start-time - The time stamp when the snapshot was initiated.
	//
	//    * status - The status of the snapshot (pending | completed | error).
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
	//
	//    * volume-id - The ID of the volume the snapshot is for.
	//
	//    * volume-size - The size of the volume, in GiB.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of snapshot results returned by DescribeSnapshots in paginated
	// output. When this parameter is used, DescribeSnapshots only returns MaxResults
	// results in a single page along with a NextToken response element. The remaining
	// results of the initial request can be seen by sending another DescribeSnapshots
	// request with the returned NextToken value. This value can be between 5 and
	// 1000; if MaxResults is given a value larger than 1000, only 1000 results
	// are returned. If this parameter is not used, then DescribeSnapshots returns
	// all results. You cannot specify this parameter and the snapshot IDs parameter
	// in the same request.
	MaxResults *int64 `type:"integer"`

	// The NextToken value returned from a previous paginated DescribeSnapshots
	// request where MaxResults was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the NextToken value. This value is null when there are no more results
	// to return.
	NextToken *string `type:"string"`

	// Describes the snapshots owned by these owners.
	OwnerIds []string `locationName:"Owner" locationNameList:"Owner" type:"list"`

	// The IDs of the AWS accounts that can create volumes from the snapshot.
	RestorableByUserIds []string `locationName:"RestorableBy" type:"list"`

	// The snapshot IDs.
	//
	// Default: Describes the snapshots for which you have create volume permissions.
	SnapshotIds []string `locationName:"SnapshotId" locationNameList:"SnapshotId" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// The NextToken value to include in a future DescribeSnapshots request. When
	// the results of a DescribeSnapshots request exceed MaxResults, this value
	// can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the snapshots.
	Snapshots []Snapshot `locationName:"snapshotSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}