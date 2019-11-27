// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeBackupInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) associated with the backup.
	//
	// BackupArn is a required field
	BackupArn *string `min:"37" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBackupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBackupInput"}

	if s.BackupArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupArn"))
	}
	if s.BackupArn != nil && len(*s.BackupArn) < 37 {
		invalidParams.Add(aws.NewErrParamMinLen("BackupArn", 37))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeBackupOutput struct {
	_ struct{} `type:"structure"`

	// Contains the description of the backup created for the table.
	BackupDescription *BackupDescription `type:"structure"`
}

// String returns the string representation
func (s DescribeBackupOutput) String() string {
	return awsutil.Prettify(s)
}
