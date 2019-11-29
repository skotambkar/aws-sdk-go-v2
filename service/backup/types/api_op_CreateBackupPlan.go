// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateBackupPlanInput struct {
	_ struct{} `type:"structure"`

	// Specifies the body of a backup plan. Includes a BackupPlanName and one or
	// more sets of Rules.
	//
	// BackupPlan is a required field
	BackupPlan *BackupPlanInput `type:"structure" required:"true"`

	// To help organize your resources, you can assign your own metadata to the
	// resources that you create. Each tag is a key-value pair. The specified tags
	// are assigned to all backups created with this plan.
	BackupPlanTags map[string]string `type:"map" sensitive:"true"`

	// Identifies the request and allows failed requests to be retried without the
	// risk of executing the operation twice. If the request includes a CreatorRequestId
	// that matches an existing backup plan, that plan is returned. This parameter
	// is optional.
	CreatorRequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateBackupPlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateBackupPlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateBackupPlanInput"}

	if s.BackupPlan == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlan"))
	}
	if s.BackupPlan != nil {
		if err := s.BackupPlan.Validate(); err != nil {
			invalidParams.AddNested("BackupPlan", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateBackupPlanOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example, arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string `type:"string"`

	// Uniquely identifies a backup plan.
	BackupPlanId *string `type:"string"`

	// The date and time that a backup plan is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp"`

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1024 bytes long. They cannot be edited.
	VersionId *string `type:"string"`
}

// String returns the string representation
func (s CreateBackupPlanOutput) String() string {
	return awsutil.Prettify(s)
}