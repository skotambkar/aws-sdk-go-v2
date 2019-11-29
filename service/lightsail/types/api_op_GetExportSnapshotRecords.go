// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetExportSnapshotRecordsInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to a specific page of results for your get export
	// snapshot records request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetExportSnapshotRecordsInput) String() string {
	return awsutil.Prettify(s)
}

type GetExportSnapshotRecordsOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects describing the export snapshot records.
	ExportSnapshotRecords []ExportSnapshotRecord `locationName:"exportSnapshotRecords" type:"list"`

	// A token used for advancing to the next page of results of your get relational
	// database bundles request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetExportSnapshotRecordsOutput) String() string {
	return awsutil.Prettify(s)
}