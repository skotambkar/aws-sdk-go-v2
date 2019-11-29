// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure for the delete Domain Association request.
type DeleteDomainAssociationInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name of the domain.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDomainAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDomainAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDomainAssociationInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDomainAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Structure for Domain Association, which associates a custom domain with an
	// Amplify App.
	//
	// DomainAssociation is a required field
	DomainAssociation *DomainAssociation `locationName:"domainAssociation" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteDomainAssociationOutput) String() string {
	return awsutil.Prettify(s)
}