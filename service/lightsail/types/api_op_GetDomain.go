// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDomainInput struct {
	_ struct{} `type:"structure"`

	// The domain name for which your want to return information about.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDomainOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about your get domain
	// request.
	Domain *Domain `locationName:"domain" type:"structure"`
}

// String returns the string representation
func (s GetDomainOutput) String() string {
	return awsutil.Prettify(s)
}