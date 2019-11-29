// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePortfolioInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The description of the portfolio.
	Description *string `type:"string"`

	// The name to use for display purposes.
	//
	// DisplayName is a required field
	DisplayName *string `min:"1" type:"string" required:"true"`

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The name of the portfolio provider.
	//
	// ProviderName is a required field
	ProviderName *string `min:"1" type:"string" required:"true"`

	// One or more tags.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreatePortfolioInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePortfolioInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePortfolioInput"}

	if s.DisplayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DisplayName"))
	}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.ProviderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProviderName"))
	}
	if s.ProviderName != nil && len(*s.ProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProviderName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePortfolioOutput struct {
	_ struct{} `type:"structure"`

	// Information about the portfolio.
	PortfolioDetail *PortfolioDetail `type:"structure"`

	// Information about the tags associated with the portfolio.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreatePortfolioOutput) String() string {
	return awsutil.Prettify(s)
}