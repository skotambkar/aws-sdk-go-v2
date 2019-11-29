// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddClientIDToOpenIDConnectProviderInput struct {
	_ struct{} `type:"structure"`

	// The client ID (also known as audience) to add to the IAM OpenID Connect provider
	// resource.
	//
	// ClientID is a required field
	ClientID *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM OpenID Connect (OIDC) provider
	// resource to add the client ID to. You can get a list of OIDC provider ARNs
	// by using the ListOpenIDConnectProviders operation.
	//
	// OpenIDConnectProviderArn is a required field
	OpenIDConnectProviderArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s AddClientIDToOpenIDConnectProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddClientIDToOpenIDConnectProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddClientIDToOpenIDConnectProviderInput"}

	if s.ClientID == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientID"))
	}
	if s.ClientID != nil && len(*s.ClientID) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientID", 1))
	}

	if s.OpenIDConnectProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("OpenIDConnectProviderArn"))
	}
	if s.OpenIDConnectProviderArn != nil && len(*s.OpenIDConnectProviderArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("OpenIDConnectProviderArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddClientIDToOpenIDConnectProviderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddClientIDToOpenIDConnectProviderOutput) String() string {
	return awsutil.Prettify(s)
}