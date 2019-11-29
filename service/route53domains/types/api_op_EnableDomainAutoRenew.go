// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableDomainAutoRenewInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that you want to enable automatic renewal for.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableDomainAutoRenewInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableDomainAutoRenewInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableDomainAutoRenewInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableDomainAutoRenewOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableDomainAutoRenewOutput) String() string {
	return awsutil.Prettify(s)
}
