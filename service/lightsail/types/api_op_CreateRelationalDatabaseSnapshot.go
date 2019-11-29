// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRelationalDatabaseSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of the database on which to base your new snapshot.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`

	// The name for your new database snapshot.
	//
	// Constraints:
	//
	//    * Must contain from 2 to 255 alphanumeric characters, or hyphens.
	//
	//    * The first and last character must be a letter or number.
	//
	// RelationalDatabaseSnapshotName is a required field
	RelationalDatabaseSnapshotName *string `locationName:"relationalDatabaseSnapshotName" type:"string" required:"true"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateRelationalDatabaseSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRelationalDatabaseSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRelationalDatabaseSnapshotInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if s.RelationalDatabaseSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRelationalDatabaseSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your create relational database snapshot
	// request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s CreateRelationalDatabaseSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}