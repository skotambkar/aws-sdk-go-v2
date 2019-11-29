// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateHsmClientCertificateInput struct {
	_ struct{} `type:"structure"`

	// The identifier to be assigned to the new HSM client certificate that the
	// cluster will use to connect to the HSM to use the database encryption keys.
	//
	// HsmClientCertificateIdentifier is a required field
	HsmClientCertificateIdentifier *string `type:"string" required:"true"`

	// A list of tag instances.
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateHsmClientCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHsmClientCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHsmClientCertificateInput"}

	if s.HsmClientCertificateIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmClientCertificateIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateHsmClientCertificateOutput struct {
	_ struct{} `type:"structure"`

	// Returns information about an HSM client certificate. The certificate is stored
	// in a secure Hardware Storage Module (HSM), and used by the Amazon Redshift
	// cluster to encrypt data files.
	HsmClientCertificate *HsmClientCertificate `type:"structure"`
}

// String returns the string representation
func (s CreateHsmClientCertificateOutput) String() string {
	return awsutil.Prettify(s)
}
