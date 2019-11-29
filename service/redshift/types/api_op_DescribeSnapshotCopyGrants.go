// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The result of the DescribeSnapshotCopyGrants action.
type DescribeSnapshotCopyGrantsInput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeSnapshotCopyGrant request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	//
	// Constraints: You can specify either the SnapshotCopyGrantName parameter or
	// the Marker parameter, but not both.
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

	// The name of the snapshot copy grant.
	SnapshotCopyGrantName *string `type:"string"`

	// A tag key or keys for which you want to return all matching resources that
	// are associated with the specified key or keys. For example, suppose that
	// you have resources tagged with keys called owner and environment. If you
	// specify both of these tag keys in the request, Amazon Redshift returns a
	// response with all resources that have either or both of these tag keys associated
	// with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching resources
	// that are associated with the specified value or values. For example, suppose
	// that you have resources tagged with values called admin and test. If you
	// specify both of these tag values in the request, Amazon Redshift returns
	// a response with all resources that have either or both of these tag values
	// associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotCopyGrantsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeSnapshotCopyGrantsOutput struct {
	_ struct{} `type:"structure"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeSnapshotCopyGrant request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	//
	// Constraints: You can specify either the SnapshotCopyGrantName parameter or
	// the Marker parameter, but not both.
	Marker *string `type:"string"`

	// The list of SnapshotCopyGrant objects.
	SnapshotCopyGrants []SnapshotCopyGrant `locationNameList:"SnapshotCopyGrant" type:"list"`
}

// String returns the string representation
func (s DescribeSnapshotCopyGrantsOutput) String() string {
	return awsutil.Prettify(s)
}