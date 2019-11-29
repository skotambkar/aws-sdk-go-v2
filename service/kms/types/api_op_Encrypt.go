// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EncryptInput struct {
	_ struct{} `type:"structure"`

	// Name-value pair that specifies the encryption context to be used for authenticated
	// encryption. If used here, the same value must be supplied to the Decrypt
	// API or decryption will fail. For more information, see Encryption Context
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context).
	EncryptionContext map[string]string `type:"map"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`

	// A unique identifier for the customer master key (CMK).
	//
	// To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name,
	// or alias ARN. When using an alias name, prefix it with "alias/". To specify
	// a CMK in a different AWS account, you must use the key ARN or alias ARN.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Alias name: alias/ExampleAlias
	//
	//    * Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey. To
	// get the alias name and alias ARN, use ListAliases.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// Data to be encrypted.
	//
	// Plaintext is automatically base64 encoded/decoded by the SDK.
	//
	// Plaintext is a required field
	Plaintext []byte `min:"1" type:"blob" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s EncryptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EncryptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EncryptInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}

	if s.Plaintext == nil {
		invalidParams.Add(aws.NewErrParamRequired("Plaintext"))
	}
	if s.Plaintext != nil && len(s.Plaintext) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Plaintext", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EncryptOutput struct {
	_ struct{} `type:"structure"`

	// The encrypted plaintext. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not encoded.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	CiphertextBlob []byte `min:"1" type:"blob"`

	// The ID of the key used during encryption.
	KeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s EncryptOutput) String() string {
	return awsutil.Prettify(s)
}
