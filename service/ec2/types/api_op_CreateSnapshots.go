// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

type CreateSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// Copies the tags from the specified volume to corresponding snapshot.
	CopyTagsFromSource enums.CopyTagsFromSource `type:"string" enum:"true"`

	// A description propagated to every snapshot specified by the instance.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The instance to specify which volumes should be included in the snapshots.
	//
	// InstanceSpecification is a required field
	InstanceSpecification *InstanceSpecification `type:"structure" required:"true"`

	// Tags to apply to every snapshot specified by the instance.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSnapshotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSnapshotsInput"}

	if s.InstanceSpecification == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceSpecification"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// List of snapshots.
	Snapshots []SnapshotInfo `locationName:"snapshotSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}