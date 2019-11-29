// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/kms/enums"
)

type GetParametersForImportInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the CMK into which you will import key material. The CMK's
	// Origin must be EXTERNAL.
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
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

	// The algorithm you will use to encrypt the key material before importing it
	// with ImportKeyMaterial. For more information, see Encrypt the Key Material
	// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys-encrypt-key-material.html)
	// in the AWS Key Management Service Developer Guide.
	//
	// WrappingAlgorithm is a required field
	WrappingAlgorithm enums.AlgorithmSpec `type:"string" required:"true" enum:"true"`

	// The type of wrapping key (public key) to return in the response. Only 2048-bit
	// RSA public keys are supported.
	//
	// WrappingKeySpec is a required field
	WrappingKeySpec enums.WrappingKeySpec `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetParametersForImportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetParametersForImportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetParametersForImportInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if len(s.WrappingAlgorithm) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("WrappingAlgorithm"))
	}
	if len(s.WrappingKeySpec) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("WrappingKeySpec"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetParametersForImportOutput struct {
	_ struct{} `type:"structure"`

	// The import token to send in a subsequent ImportKeyMaterial request.
	//
	// ImportToken is automatically base64 encoded/decoded by the SDK.
	ImportToken []byte `min:"1" type:"blob"`

	// The identifier of the CMK to use in a subsequent ImportKeyMaterial request.
	// This is the same CMK specified in the GetParametersForImport request.
	KeyId *string `min:"1" type:"string"`

	// The time at which the import token and public key are no longer valid. After
	// this time, you cannot use them to make an ImportKeyMaterial request and you
	// must send another GetParametersForImport request to get new ones.
	ParametersValidTo *time.Time `type:"timestamp"`

	// The public key to use to encrypt the key material before importing it with
	// ImportKeyMaterial.
	//
	// PublicKey is automatically base64 encoded/decoded by the SDK.
	PublicKey []byte `min:"1" type:"blob" sensitive:"true"`
}

// String returns the string representation
func (s GetParametersForImportOutput) String() string {
	return awsutil.Prettify(s)
}