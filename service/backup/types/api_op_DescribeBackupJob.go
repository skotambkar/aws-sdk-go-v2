// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/backup/enums"
)

type DescribeBackupJobInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a request to AWS Backup to back up a resource.
	//
	// BackupJobId is a required field
	BackupJobId *string `location:"uri" locationName:"backupJobId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBackupJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBackupJobInput"}

	if s.BackupJobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupJobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeBackupJobOutput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a request to AWS Backup to back up a resource.
	BackupJobId *string `type:"string"`

	// The size, in bytes, of a backup.
	BackupSizeInBytes *int64 `type:"long"`

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string `type:"string"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	BackupVaultName *string `type:"string"`

	// The size in bytes transferred to a backup vault at the time that the job
	// status was queried.
	BytesTransferred *int64 `type:"long"`

	// The date and time that a job to create a backup job is completed, in Unix
	// format and Coordinated Universal Time (UTC). The value of CreationDate is
	// accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time `type:"timestamp"`

	// Contains identifying information about the creation of a backup job, including
	// the BackupPlanArn, BackupPlanId, BackupPlanVersion, and BackupRuleId of the
	// backup plan that is used to create it.
	CreatedBy *RecoveryPointCreator `type:"structure"`

	// The date and time that a backup job is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp"`

	// The date and time that a job to back up resources is expected to be completed,
	// in Unix format and Coordinated Universal Time (UTC). The value of ExpectedCompletionDate
	// is accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	ExpectedCompletionDate *time.Time `type:"timestamp"`

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string `type:"string"`

	// Contains an estimated percentage that is complete of a job at the time the
	// job status was queried.
	PercentDone *string `type:"string"`

	// An ARN that uniquely identifies a recovery point; for example, arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string `type:"string"`

	// An ARN that uniquely identifies a saved resource. The format of the ARN depends
	// on the resource type.
	ResourceArn *string `type:"string"`

	// The type of AWS resource to be backed-up; for example, an Amazon Elastic
	// Block Store (Amazon EBS) volume or an Amazon Relational Database Service
	// (Amazon RDS) database.
	ResourceType *string `type:"string"`

	// Specifies the time in Unix format and Coordinated Universal Time (UTC) when
	// a backup job must be started before it is canceled. The value is calculated
	// by adding the start window to the scheduled time. So if the scheduled time
	// were 6:00 PM and the start window is 2 hours, the StartBy time would be 8:00
	// PM on the date specified. The value of StartBy is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	StartBy *time.Time `type:"timestamp"`

	// The current state of a resource recovery point.
	State enums.BackupJobState `type:"string" enum:"true"`

	// A detailed message explaining the status of the job to back up a resource.
	StatusMessage *string `type:"string"`
}

// String returns the string representation
func (s DescribeBackupJobOutput) String() string {
	return awsutil.Prettify(s)
}
