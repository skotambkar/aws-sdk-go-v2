// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateKeyPairInput struct {
	_ struct{} `type:"structure"`

	// The name for your new key pair.
	//
	// KeyPairName is a required field
	KeyPairName *string `locationName:"keyPairName" type:"string" required:"true"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateKeyPairInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateKeyPairInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateKeyPairInput"}

	if s.KeyPairName == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyPairName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateKeyPairOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the new key pair
	// you just created.
	KeyPair *KeyPair `locationName:"keyPair" type:"structure"`

	// An array of key-value pairs containing information about the results of your
	// create key pair request.
	Operation *Operation `locationName:"operation" type:"structure"`

	// A base64-encoded RSA private key.
	PrivateKeyBase64 *string `locationName:"privateKeyBase64" type:"string"`

	// A base64-encoded public key of the ssh-rsa type.
	PublicKeyBase64 *string `locationName:"publicKeyBase64" type:"string"`
}

// String returns the string representation
func (s CreateKeyPairOutput) String() string {
	return awsutil.Prettify(s)
}