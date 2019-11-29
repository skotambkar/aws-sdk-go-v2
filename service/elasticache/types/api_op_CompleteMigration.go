// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CompleteMigrationInput struct {
	_ struct{} `type:"structure"`

	// Forces the migration to stop without ensuring that data is in sync. It is
	// recommended to use this option only to abort the migration and not recommended
	// when application wants to continue migration to ElastiCache.
	Force *bool `type:"boolean"`

	// The ID of the replication group to which data is being migrated.
	//
	// ReplicationGroupId is a required field
	ReplicationGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CompleteMigrationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CompleteMigrationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CompleteMigrationInput"}

	if s.ReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CompleteMigrationOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *ReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s CompleteMigrationOutput) String() string {
	return awsutil.Prettify(s)
}