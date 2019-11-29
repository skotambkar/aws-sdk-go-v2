// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateGroupCertificateConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The amount of time remaining before the certificate expires, in milliseconds.
	CertificateExpiryInMilliseconds *string `type:"string"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateGroupCertificateConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGroupCertificateConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGroupCertificateConfigurationInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateGroupCertificateConfigurationOutput struct {
	_ struct{} `type:"structure"`

	CertificateAuthorityExpiryInMilliseconds *string `type:"string"`

	CertificateExpiryInMilliseconds *string `type:"string"`

	GroupId *string `type:"string"`
}

// String returns the string representation
func (s UpdateGroupCertificateConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
