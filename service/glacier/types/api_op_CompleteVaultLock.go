// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input values for CompleteVaultLock.
type CompleteVaultLockInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You
	// can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you specify your account ID, do
	// not include any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The lockId value is the lock ID obtained from a InitiateVaultLock request.
	//
	// LockId is a required field
	LockId *string `location:"uri" locationName:"lockId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s CompleteVaultLockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CompleteVaultLockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CompleteVaultLockInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.LockId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LockId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CompleteVaultLockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CompleteVaultLockOutput) String() string {
	return awsutil.Prettify(s)
}
