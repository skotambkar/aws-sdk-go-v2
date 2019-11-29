// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListBillingGroupsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return per request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Limit the results to billing groups whose names have the given prefix.
	NamePrefixFilter *string `location:"querystring" locationName:"namePrefixFilter" min:"1" type:"string"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBillingGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBillingGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBillingGroupsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NamePrefixFilter != nil && len(*s.NamePrefixFilter) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamePrefixFilter", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListBillingGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The list of billing groups.
	BillingGroups []GroupNameAndArn `locationName:"billingGroups" type:"list"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBillingGroupsOutput) String() string {
	return awsutil.Prettify(s)
}
