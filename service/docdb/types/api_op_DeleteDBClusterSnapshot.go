// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to DeleteDBClusterSnapshot.
type DeleteDBClusterSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the DB cluster snapshot to delete.
	//
	// Constraints: Must be the name of an existing DB cluster snapshot in the available
	// state.
	//
	// DBClusterSnapshotIdentifier is a required field
	DBClusterSnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBClusterSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBClusterSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBClusterSnapshotInput"}

	if s.DBClusterSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBClusterSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about a DB cluster snapshot.
	DBClusterSnapshot *DBClusterSnapshot `type:"structure"`
}

// String returns the string representation
func (s DeleteDBClusterSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}
