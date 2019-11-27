// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure for update Domain Association request.
type UpdateDomainAssociationInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name of the domain.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`

	// Enables automated creation of Subdomains for branches.
	EnableAutoSubDomain *bool `locationName:"enableAutoSubDomain" type:"boolean"`

	// Setting structure for the Subdomain.
	//
	// SubDomainSettings is a required field
	SubDomainSettings []SubDomainSetting `locationName:"subDomainSettings" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateDomainAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDomainAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDomainAssociationInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.SubDomainSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubDomainSettings"))
	}
	if s.SubDomainSettings != nil {
		for i, v := range s.SubDomainSettings {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SubDomainSettings", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result structure for the update Domain Association request.
type UpdateDomainAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Domain Association structure.
	//
	// DomainAssociation is a required field
	DomainAssociation *DomainAssociation `locationName:"domainAssociation" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateDomainAssociationOutput) String() string {
	return awsutil.Prettify(s)
}
