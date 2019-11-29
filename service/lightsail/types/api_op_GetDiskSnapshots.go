// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDiskSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to the next page of results from your GetDiskSnapshots
	// request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetDiskSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

type GetDiskSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects containing information about all block storage disk snapshots.
	DiskSnapshots []DiskSnapshot `locationName:"diskSnapshots" type:"list"`

	// A token used for advancing to the next page of results from your GetDiskSnapshots
	// request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetDiskSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}