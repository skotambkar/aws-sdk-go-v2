// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RotateEncryptionKeyInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the cluster that you want to rotate the encryption
	// keys for.
	//
	// Constraints: Must be the name of valid cluster that has encryption enabled.
	//
	// ClusterIdentifier is a required field
	ClusterIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RotateEncryptionKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RotateEncryptionKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RotateEncryptionKeyInput"}

	if s.ClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RotateEncryptionKeyOutput struct {
	_ struct{} `type:"structure"`

	// Describes a cluster.
	Cluster *Cluster `type:"structure"`
}

// String returns the string representation
func (s RotateEncryptionKeyOutput) String() string {
	return awsutil.Prettify(s)
}