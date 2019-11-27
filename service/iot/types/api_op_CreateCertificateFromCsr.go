// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the CreateCertificateFromCsr operation.
type CreateCertificateFromCsrInput struct {
	_ struct{} `type:"structure"`

	// The certificate signing request (CSR).
	//
	// CertificateSigningRequest is a required field
	CertificateSigningRequest *string `locationName:"certificateSigningRequest" min:"1" type:"string" required:"true"`

	// Specifies whether the certificate is active.
	SetAsActive *bool `location:"querystring" locationName:"setAsActive" type:"boolean"`
}

// String returns the string representation
func (s CreateCertificateFromCsrInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCertificateFromCsrInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCertificateFromCsrInput"}

	if s.CertificateSigningRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateSigningRequest"))
	}
	if s.CertificateSigningRequest != nil && len(*s.CertificateSigningRequest) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateSigningRequest", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output from the CreateCertificateFromCsr operation.
type CreateCertificateFromCsrOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the certificate. You can use the ARN as
	// a principal for policy operations.
	CertificateArn *string `locationName:"certificateArn" type:"string"`

	// The ID of the certificate. Certificate management operations only take a
	// certificateId.
	CertificateId *string `locationName:"certificateId" min:"64" type:"string"`

	// The certificate data, in PEM format.
	CertificatePem *string `locationName:"certificatePem" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateCertificateFromCsrOutput) String() string {
	return awsutil.Prettify(s)
}
