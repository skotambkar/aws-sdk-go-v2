// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/xray/enums"
)

type PutEncryptionConfigInput struct {
	_ struct{} `type:"structure"`

	// An AWS KMS customer master key (CMK) in one of the following formats:
	//
	//    * Alias - The name of the key. For example, alias/MyKey.
	//
	//    * Key ID - The KMS key ID of the key. For example, ae4aa6d49-a4d8-9df9-a475-4ff6d7898456.
	//
	//    * ARN - The full Amazon Resource Name of the key ID or alias. For example,
	//    arn:aws:kms:us-east-2:123456789012:key/ae4aa6d49-a4d8-9df9-a475-4ff6d7898456.
	//    Use this format to specify a key in a different account.
	//
	// Omit this key if you set Type to NONE.
	KeyId *string `min:"1" type:"string"`

	// The type of encryption. Set to KMS to use your own key for encryption. Set
	// to NONE for default encryption.
	//
	// Type is a required field
	Type enums.EncryptionType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s PutEncryptionConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEncryptionConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutEncryptionConfigInput"}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutEncryptionConfigOutput struct {
	_ struct{} `type:"structure"`

	// The new encryption configuration.
	EncryptionConfig *EncryptionConfig `type:"structure"`
}

// String returns the string representation
func (s PutEncryptionConfigOutput) String() string {
	return awsutil.Prettify(s)
}
