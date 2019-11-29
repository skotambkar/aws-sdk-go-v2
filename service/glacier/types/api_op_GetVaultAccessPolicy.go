// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for GetVaultAccessPolicy.
type GetVaultAccessPolicyInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVaultAccessPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVaultAccessPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVaultAccessPolicyInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Output for GetVaultAccessPolicy.
type GetVaultAccessPolicyOutput struct {
	_ struct{} `type:"structure" payload:"Policy"`

	// Contains the returned vault access policy as a JSON string.
	Policy *VaultAccessPolicy `locationName:"policy" type:"structure"`
}

// String returns the string representation
func (s GetVaultAccessPolicyOutput) String() string {
	return awsutil.Prettify(s)
}