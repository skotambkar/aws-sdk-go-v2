// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateDomainInput struct {
	_ struct{} `type:"structure"`

	// The ARN of an issued ACM certificate that is valid for the domain being associated.
	//
	// AcmCertificateArn is a required field
	AcmCertificateArn *string `type:"string" required:"true"`

	// The name to display.
	DisplayName *string `type:"string"`

	// The fully qualified domain name (FQDN).
	//
	// DomainName is a required field
	DomainName *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateDomainInput"}

	if s.AcmCertificateArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcmCertificateArn"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 1))
	}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateDomainOutput) String() string {
	return awsutil.Prettify(s)
}
