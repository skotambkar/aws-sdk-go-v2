// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyClusterSnapshotInput struct {
	_ struct{} `type:"structure"`

	// A Boolean option to override an exception if the retention period has already
	// passed.
	Force *bool `type:"boolean"`

	// The number of days that a manual snapshot is retained. If the value is -1,
	// the manual snapshot is retained indefinitely.
	//
	// If the manual snapshot falls outside of the new retention period, you can
	// specify the force option to immediately delete the snapshot.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	ManualSnapshotRetentionPeriod *int64 `type:"integer"`

	// The identifier of the snapshot whose setting you want to modify.
	//
	// SnapshotIdentifier is a required field
	SnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyClusterSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyClusterSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyClusterSnapshotInput"}

	if s.SnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyClusterSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Describes a snapshot.
	Snapshot *Snapshot `type:"structure"`
}

// String returns the string representation
func (s ModifyClusterSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}