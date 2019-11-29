// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetConferenceProviderInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the newly created conference provider.
	//
	// ConferenceProviderArn is a required field
	ConferenceProviderArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetConferenceProviderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConferenceProviderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConferenceProviderInput"}

	if s.ConferenceProviderArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConferenceProviderArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetConferenceProviderOutput struct {
	_ struct{} `type:"structure"`

	// The conference provider.
	ConferenceProvider *ConferenceProvider `type:"structure"`
}

// String returns the string representation
func (s GetConferenceProviderOutput) String() string {
	return awsutil.Prettify(s)
}