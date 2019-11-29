// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input to the RegisterCACertificate operation.
type RegisterCACertificateInput struct {
	_ struct{} `type:"structure"`

	// Allows this CA certificate to be used for auto registration of device certificates.
	AllowAutoRegistration *bool `location:"querystring" locationName:"allowAutoRegistration" type:"boolean"`

	// The CA certificate.
	//
	// CaCertificate is a required field
	CaCertificate *string `locationName:"caCertificate" min:"1" type:"string" required:"true"`

	// Information about the registration configuration.
	RegistrationConfig *RegistrationConfig `locationName:"registrationConfig" type:"structure"`

	// A boolean value that specifies if the CA certificate is set to active.
	SetAsActive *bool `location:"querystring" locationName:"setAsActive" type:"boolean"`

	// The private key verification certificate.
	//
	// VerificationCertificate is a required field
	VerificationCertificate *string `locationName:"verificationCertificate" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterCACertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterCACertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterCACertificateInput"}

	if s.CaCertificate == nil {
		invalidParams.Add(aws.NewErrParamRequired("CaCertificate"))
	}
	if s.CaCertificate != nil && len(*s.CaCertificate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CaCertificate", 1))
	}

	if s.VerificationCertificate == nil {
		invalidParams.Add(aws.NewErrParamRequired("VerificationCertificate"))
	}
	if s.VerificationCertificate != nil && len(*s.VerificationCertificate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VerificationCertificate", 1))
	}
	if s.RegistrationConfig != nil {
		if err := s.RegistrationConfig.Validate(); err != nil {
			invalidParams.AddNested("RegistrationConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output from the RegisterCACertificateResponse operation.
type RegisterCACertificateOutput struct {
	_ struct{} `type:"structure"`

	// The CA certificate ARN.
	CertificateArn *string `locationName:"certificateArn" type:"string"`

	// The CA certificate identifier.
	CertificateId *string `locationName:"certificateId" min:"64" type:"string"`
}

// String returns the string representation
func (s RegisterCACertificateOutput) String() string {
	return awsutil.Prettify(s)
}