// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListWebsiteAuthorizationProvidersInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`

	// The maximum number of results to be included in the next page.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListWebsiteAuthorizationProvidersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListWebsiteAuthorizationProvidersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListWebsiteAuthorizationProvidersInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListWebsiteAuthorizationProvidersOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`

	// The website authorization providers.
	WebsiteAuthorizationProviders []WebsiteAuthorizationProviderSummary `type:"list"`
}

// String returns the string representation
func (s ListWebsiteAuthorizationProvidersOutput) String() string {
	return awsutil.Prettify(s)
}