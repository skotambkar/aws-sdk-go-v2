// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RevokeGrantInput struct {
	_ struct{} `type:"structure"`

	// Identifier of the grant to be revoked.
	//
	// GrantId is a required field
	GrantId *string `min:"1" type:"string" required:"true"`

	// A unique identifier for the customer master key associated with the grant.
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK. To specify
	// a CMK in a different AWS account, you must use the key ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RevokeGrantInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeGrantInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RevokeGrantInput"}

	if s.GrantId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GrantId"))
	}
	if s.GrantId != nil && len(*s.GrantId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GrantId", 1))
	}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RevokeGrantOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RevokeGrantOutput) String() string {
	return awsutil.Prettify(s)
}