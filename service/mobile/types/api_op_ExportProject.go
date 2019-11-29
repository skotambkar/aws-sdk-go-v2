// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure used in requests to export project configuration details.
type ExportProjectInput struct {
	_ struct{} `type:"structure"`

	// Unique project identifier.
	//
	// ProjectId is a required field
	ProjectId *string `location:"uri" locationName:"projectId" type:"string" required:"true"`
}

// String returns the string representation
func (s ExportProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExportProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExportProjectInput"}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result structure used for requests to export project configuration details.
type ExportProjectOutput struct {
	_ struct{} `type:"structure"`

	// URL which can be used to download the exported project configuation file(s).
	DownloadUrl *string `locationName:"downloadUrl" type:"string"`

	// URL which can be shared to allow other AWS users to create their own project
	// in AWS Mobile Hub with the same configuration as the specified project. This
	// URL pertains to a snapshot in time of the project configuration that is created
	// when this API is called. If you want to share additional changes to your
	// project configuration, then you will need to create and share a new snapshot
	// by calling this method again.
	ShareUrl *string `locationName:"shareUrl" type:"string"`

	// Unique identifier for the exported snapshot of the project configuration.
	// This snapshot identifier is included in the share URL.
	SnapshotId *string `locationName:"snapshotId" type:"string"`
}

// String returns the string representation
func (s ExportProjectOutput) String() string {
	return awsutil.Prettify(s)
}
