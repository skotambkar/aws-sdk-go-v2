// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iam/enums"
)

type UpdateSigningCertificateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the signing certificate you want to update.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// CertificateId is a required field
	CertificateId *string `min:"24" type:"string" required:"true"`

	// The status you want to assign to the certificate. Active means that the certificate
	// can be used for API calls to AWS Inactive means that the certificate cannot
	// be used.
	//
	// Status is a required field
	Status enums.StatusType `type:"string" required:"true" enum:"true"`

	// The name of the IAM user the signing certificate belongs to.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateSigningCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSigningCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSigningCertificateInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}
	if s.CertificateId != nil && len(*s.CertificateId) < 24 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateId", 24))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSigningCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateSigningCertificateOutput) String() string {
	return awsutil.Prettify(s)
}