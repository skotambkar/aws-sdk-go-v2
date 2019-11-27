// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Provides options for deleting an archive from an Amazon S3 Glacier vault.
type DeleteArchiveInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The ID of the archive to delete.
	//
	// ArchiveId is a required field
	ArchiveId *string `location:"uri" locationName:"archiveId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteArchiveInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteArchiveInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteArchiveInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.ArchiveId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ArchiveId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteArchiveOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteArchiveOutput) String() string {
	return awsutil.Prettify(s)
}
