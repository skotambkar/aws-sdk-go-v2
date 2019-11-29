// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateKeyPairInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// A unique name for the key pair.
	//
	// Constraints: Up to 255 ASCII characters
	//
	// KeyName is a required field
	KeyName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateKeyPairInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateKeyPairInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateKeyPairInput"}

	if s.KeyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes a key pair.
type CreateKeyPairOutput struct {
	_ struct{} `type:"structure"`

	// The SHA-1 digest of the DER encoded private key.
	KeyFingerprint *string `locationName:"keyFingerprint" type:"string"`

	// An unencrypted PEM encoded RSA private key.
	KeyMaterial *string `locationName:"keyMaterial" type:"string" sensitive:"true"`

	// The name of the key pair.
	KeyName *string `locationName:"keyName" type:"string"`
}

// String returns the string representation
func (s CreateKeyPairOutput) String() string {
	return awsutil.Prettify(s)
}
