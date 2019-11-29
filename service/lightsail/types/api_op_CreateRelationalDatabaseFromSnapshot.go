// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRelationalDatabaseFromSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone in which to create your new database. Use the us-east-2a
	// case-sensitive format.
	//
	// You can get a list of Availability Zones by using the get regions operation.
	// Be sure to add the include relational database Availability Zones parameter
	// to your request.
	AvailabilityZone *string `locationName:"availabilityZone" type:"string"`

	// Specifies the accessibility options for your new database. A value of true
	// specifies a database that is available to resources outside of your Lightsail
	// account. A value of false specifies a database that is available only to
	// your Lightsail resources in the same region as your database.
	PubliclyAccessible *bool `locationName:"publiclyAccessible" type:"boolean"`

	// The bundle ID for your new database. A bundle describes the performance specifications
	// for your database.
	//
	// You can get a list of database bundle IDs by using the get relational database
	// bundles operation.
	//
	// When creating a new database from a snapshot, you cannot choose a bundle
	// that is smaller than the bundle of the source database.
	RelationalDatabaseBundleId *string `locationName:"relationalDatabaseBundleId" type:"string"`

	// The name to use for your new database.
	//
	// Constraints:
	//
	//    * Must contain from 2 to 255 alphanumeric characters, or hyphens.
	//
	//    * The first and last character must be a letter or number.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`

	// The name of the database snapshot from which to create your new database.
	RelationalDatabaseSnapshotName *string `locationName:"relationalDatabaseSnapshotName" type:"string"`

	// The date and time to restore your database from.
	//
	// Constraints:
	//
	//    * Must be before the latest restorable time for the database.
	//
	//    * Cannot be specified if the use latest restorable time parameter is true.
	//
	//    * Specified in Coordinated Universal Time (UTC).
	//
	//    * Specified in the Unix time format. For example, if you wish to use a
	//    restore time of October 1, 2018, at 8 PM UTC, then you input 1538424000
	//    as the restore time.
	RestoreTime *time.Time `locationName:"restoreTime" type:"timestamp"`

	// The name of the source database.
	SourceRelationalDatabaseName *string `locationName:"sourceRelationalDatabaseName" type:"string"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`

	// Specifies whether your database is restored from the latest backup time.
	// A value of true restores from the latest backup time.
	//
	// Default: false
	//
	// Constraints: Cannot be specified if the restore time parameter is provided.
	UseLatestRestorableTime *bool `locationName:"useLatestRestorableTime" type:"boolean"`
}

// String returns the string representation
func (s CreateRelationalDatabaseFromSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRelationalDatabaseFromSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRelationalDatabaseFromSnapshotInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateRelationalDatabaseFromSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your create relational database from snapshot
	// request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s CreateRelationalDatabaseFromSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}