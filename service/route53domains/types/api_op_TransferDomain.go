// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The TransferDomain request includes the following elements.
type TransferDomainInput struct {
	_ struct{} `type:"structure"`

	// Provides detailed contact information.
	//
	// AdminContact is a required field
	AdminContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`

	// The authorization code for the domain. You get this value from the current
	// registrar.
	AuthCode *string `type:"string" sensitive:"true"`

	// Indicates whether the domain will be automatically renewed (true) or not
	// (false). Autorenewal only takes effect after the account is charged.
	//
	// Default: true
	AutoRenew *bool `type:"boolean"`

	// The name of the domain that you want to transfer to Amazon Route 53.
	//
	// Constraints: The domain name can contain only the letters a through z, the
	// numbers 0 through 9, and hyphen (-). Internationalized Domain Names are not
	// supported.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// The number of years that you want to register the domain for. Domains are
	// registered for a minimum of one year. The maximum period depends on the top-level
	// domain.
	//
	// Default: 1
	//
	// DurationInYears is a required field
	DurationInYears *int64 `min:"1" type:"integer" required:"true"`

	// Reserved for future use.
	IdnLangCode *string `type:"string"`

	// Contains details for the host and glue IP addresses.
	Nameservers []Nameserver `type:"list"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the admin contact.
	//
	// Default: true
	PrivacyProtectAdminContact *bool `type:"boolean"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the registrant contact (domain
	// owner).
	//
	// Default: true
	PrivacyProtectRegistrantContact *bool `type:"boolean"`

	// Whether you want to conceal contact information from WHOIS queries. If you
	// specify true, WHOIS ("who is") queries return contact information either
	// for Amazon Registrar (for .com, .net, and .org domains) or for our registrar
	// associate, Gandi (for all other TLDs). If you specify false, WHOIS queries
	// return the information that you entered for the technical contact.
	//
	// Default: true
	PrivacyProtectTechContact *bool `type:"boolean"`

	// Provides detailed contact information.
	//
	// RegistrantContact is a required field
	RegistrantContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`

	// Provides detailed contact information.
	//
	// TechContact is a required field
	TechContact *ContactDetail `type:"structure" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s TransferDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TransferDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TransferDomainInput"}

	if s.AdminContact == nil {
		invalidParams.Add(aws.NewErrParamRequired("AdminContact"))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.DurationInYears == nil {
		invalidParams.Add(aws.NewErrParamRequired("DurationInYears"))
	}
	if s.DurationInYears != nil && *s.DurationInYears < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("DurationInYears", 1))
	}

	if s.RegistrantContact == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegistrantContact"))
	}

	if s.TechContact == nil {
		invalidParams.Add(aws.NewErrParamRequired("TechContact"))
	}
	if s.AdminContact != nil {
		if err := s.AdminContact.Validate(); err != nil {
			invalidParams.AddNested("AdminContact", err.(aws.ErrInvalidParams))
		}
	}
	if s.Nameservers != nil {
		for i, v := range s.Nameservers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Nameservers", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RegistrantContact != nil {
		if err := s.RegistrantContact.Validate(); err != nil {
			invalidParams.AddNested("RegistrantContact", err.(aws.ErrInvalidParams))
		}
	}
	if s.TechContact != nil {
		if err := s.TechContact.Validate(); err != nil {
			invalidParams.AddNested("TechContact", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The TranserDomain response includes the following element.
type TransferDomainOutput struct {
	_ struct{} `type:"structure"`

	// Identifier for tracking the progress of the request. To use this ID to query
	// the operation status, use GetOperationDetail.
	//
	// OperationId is a required field
	OperationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TransferDomainOutput) String() string {
	return awsutil.Prettify(s)
}
