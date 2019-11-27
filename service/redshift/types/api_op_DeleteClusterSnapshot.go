// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteClusterSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the cluster the snapshot was created from. This
	// parameter is required if your IAM user has a policy containing a snapshot
	// resource element that specifies anything other than * for the cluster name.
	//
	// Constraints: Must be the name of valid cluster.
	SnapshotClusterIdentifier *string `type:"string"`

	// The unique identifier of the manual snapshot to be deleted.
	//
	// Constraints: Must be the name of an existing snapshot that is in the available,
	// failed, or cancelled state.
	//
	// SnapshotIdentifier is a required field
	SnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteClusterSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteClusterSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteClusterSnapshotInput"}

	if s.SnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteClusterSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Describes a snapshot.
	Snapshot *Snapshot `type:"structure"`
}

// String returns the string representation
func (s DeleteClusterSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}
