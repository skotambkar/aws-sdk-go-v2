// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CopyClusterSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The number of days that a manual snapshot is retained. If the value is -1,
	// the manual snapshot is retained indefinitely.
	//
	// The value must be either -1 or an integer between 1 and 3,653.
	//
	// The default value is -1.
	ManualSnapshotRetentionPeriod *int64 `type:"integer"`

	// The identifier of the cluster the source snapshot was created from. This
	// parameter is required if your IAM user has a policy containing a snapshot
	// resource element that specifies anything other than * for the cluster name.
	//
	// Constraints:
	//
	//    * Must be the identifier for a valid cluster.
	SourceSnapshotClusterIdentifier *string `type:"string"`

	// The identifier for the source snapshot.
	//
	// Constraints:
	//
	//    * Must be the identifier for a valid automated snapshot whose state is
	//    available.
	//
	// SourceSnapshotIdentifier is a required field
	SourceSnapshotIdentifier *string `type:"string" required:"true"`

	// The identifier given to the new manual snapshot.
	//
	// Constraints:
	//
	//    * Cannot be null, empty, or blank.
	//
	//    * Must contain from 1 to 255 alphanumeric characters or hyphens.
	//
	//    * First character must be a letter.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	//
	//    * Must be unique for the AWS account that is making the request.
	//
	// TargetSnapshotIdentifier is a required field
	TargetSnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopyClusterSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyClusterSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyClusterSnapshotInput"}

	if s.SourceSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceSnapshotIdentifier"))
	}

	if s.TargetSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CopyClusterSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Describes a snapshot.
	Snapshot *Snapshot `type:"structure"`
}

// String returns the string representation
func (s CopyClusterSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}
