// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UploadSigningCertificateInput struct {
	_ struct{} `type:"structure"`

	// The contents of the signing certificate.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// CertificateBody is a required field
	CertificateBody *string `min:"1" type:"string" required:"true"`

	// The name of the user the signing certificate is for.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UploadSigningCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadSigningCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UploadSigningCertificateInput"}

	if s.CertificateBody == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateBody"))
	}
	if s.CertificateBody != nil && len(*s.CertificateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateBody", 1))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful UploadSigningCertificate request.
type UploadSigningCertificateOutput struct {
	_ struct{} `type:"structure"`

	// Information about the certificate.
	//
	// Certificate is a required field
	Certificate *SigningCertificate `type:"structure" required:"true"`
}

// String returns the string representation
func (s UploadSigningCertificateOutput) String() string {
	return awsutil.Prettify(s)
}
