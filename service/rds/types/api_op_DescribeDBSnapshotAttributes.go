// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDBSnapshotAttributesInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the DB snapshot to describe the attributes for.
	//
	// DBSnapshotIdentifier is a required field
	DBSnapshotIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDBSnapshotAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBSnapshotAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBSnapshotAttributesInput"}

	if s.DBSnapshotIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSnapshotIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDBSnapshotAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Contains the results of a successful call to the DescribeDBSnapshotAttributes
	// API action.
	//
	// Manual DB snapshot attributes are used to authorize other AWS accounts to
	// copy or restore a manual DB snapshot. For more information, see the ModifyDBSnapshotAttribute
	// API action.
	DBSnapshotAttributesResult *DBSnapshotAttributesResult `type:"structure"`
}

// String returns the string representation
func (s DescribeDBSnapshotAttributesOutput) String() string {
	return awsutil.Prettify(s)
}