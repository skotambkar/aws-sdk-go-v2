// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/backup/enums"
)

type ListBackupJobsInput struct {
	_ struct{} `type:"structure"`

	// Returns only backup jobs that will be stored in the specified backup vault.
	// Backup vaults are identified by names that are unique to the account used
	// to create them and the AWS Region where they are created. They consist of
	// lowercase letters, numbers, and hyphens.
	ByBackupVaultName *string `location:"querystring" locationName:"backupVaultName" type:"string"`

	// Returns only backup jobs that were created after the specified date.
	ByCreatedAfter *time.Time `location:"querystring" locationName:"createdAfter" type:"timestamp"`

	// Returns only backup jobs that were created before the specified date.
	ByCreatedBefore *time.Time `location:"querystring" locationName:"createdBefore" type:"timestamp"`

	// Returns only backup jobs that match the specified resource Amazon Resource
	// Name (ARN).
	ByResourceArn *string `location:"querystring" locationName:"resourceArn" type:"string"`

	// Returns only backup jobs for the specified resources:
	//
	//    * EBS for Amazon Elastic Block Store
	//
	//    * SGW for AWS Storage Gateway
	//
	//    * RDS for Amazon Relational Database Service
	//
	//    * DDB for Amazon DynamoDB
	//
	//    * EFS for Amazon Elastic File System
	ByResourceType *string `location:"querystring" locationName:"resourceType" type:"string"`

	// Returns only backup jobs that are in the specified state.
	ByState enums.BackupJobState `location:"querystring" locationName:"state" type:"string" enum:"true"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBackupJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBackupJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBackupJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListBackupJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of structures containing metadata about your backup jobs returned
	// in JSON format.
	BackupJobs []BackupJob `type:"list"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListBackupJobsOutput) String() string {
	return awsutil.Prettify(s)
}
