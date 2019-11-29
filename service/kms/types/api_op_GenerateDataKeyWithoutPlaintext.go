// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/kms/enums"
)

type GenerateDataKeyWithoutPlaintextInput struct {
	_ struct{} `type:"structure"`

	// A set of key-value pairs that represents additional authenticated data.
	//
	// For more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the AWS Key Management Service Developer Guide.
	EncryptionContext map[string]string `type:"map"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`

	// The identifier of the customer master key (CMK) that encrypts the data key.
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

	// The length of the data key. Use AES_128 to generate a 128-bit symmetric key,
	// or AES_256 to generate a 256-bit symmetric key.
	KeySpec enums.DataKeySpec `type:"string" enum:"true"`

	// The length of the data key in bytes. For example, use the value 64 to generate
	// a 512-bit data key (64 bytes is 512 bits). For common key lengths (128-bit
	// and 256-bit symmetric keys), we recommend that you use the KeySpec field
	// instead of this one.
	NumberOfBytes *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s GenerateDataKeyWithoutPlaintextInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateDataKeyWithoutPlaintextInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateDataKeyWithoutPlaintextInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.NumberOfBytes != nil && *s.NumberOfBytes < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("NumberOfBytes", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GenerateDataKeyWithoutPlaintextOutput struct {
	_ struct{} `type:"structure"`

	// The encrypted data key. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not encoded.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	CiphertextBlob []byte `min:"1" type:"blob"`

	// The identifier of the CMK that encrypted the data key.
	KeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GenerateDataKeyWithoutPlaintextOutput) String() string {
	return awsutil.Prettify(s)
}
