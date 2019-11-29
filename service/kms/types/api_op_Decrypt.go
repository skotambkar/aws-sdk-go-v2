// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DecryptInput struct {
	_ struct{} `type:"structure"`

	// Ciphertext to be decrypted. The blob includes metadata.
	//
	// CiphertextBlob is automatically base64 encoded/decoded by the SDK.
	//
	// CiphertextBlob is a required field
	CiphertextBlob []byte `min:"1" type:"blob" required:"true"`

	// The encryption context. If this was specified in the Encrypt function, it
	// must be specified here or the decryption operation will fail. For more information,
	// see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context).
	EncryptionContext map[string]string `type:"map"`

	// A list of grant tokens.
	//
	// For more information, see Grant Tokens (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []string `type:"list"`
}

// String returns the string representation
func (s DecryptInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecryptInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DecryptInput"}

	if s.CiphertextBlob == nil {
		invalidParams.Add(aws.NewErrParamRequired("CiphertextBlob"))
	}
	if s.CiphertextBlob != nil && len(s.CiphertextBlob) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CiphertextBlob", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DecryptOutput struct {
	_ struct{} `type:"structure"`

	// ARN of the key used to perform the decryption. This value is returned if
	// no errors are encountered during the operation.
	KeyId *string `min:"1" type:"string"`

	// Decrypted plaintext data. When you use the HTTP API or the AWS CLI, the value
	// is Base64-encoded. Otherwise, it is not encoded.
	//
	// Plaintext is automatically base64 encoded/decoded by the SDK.
	Plaintext []byte `min:"1" type:"blob" sensitive:"true"`
}

// String returns the string representation
func (s DecryptOutput) String() string {
	return awsutil.Prettify(s)
}