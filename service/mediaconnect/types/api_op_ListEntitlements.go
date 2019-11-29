// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListEntitlementsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListEntitlementsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEntitlementsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEntitlementsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a successful ListEntitlements request. The response includes
// the ARN of each entitlement, the name of the associated flow, and the NextToken
// to use in a subsequent ListEntitlements request.
type ListEntitlementsOutput struct {
	_ struct{} `type:"structure"`

	// A list of entitlements that have been granted to you from other AWS accounts.
	Entitlements []ListedEntitlement `locationName:"entitlements" type:"list"`

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListEntitlements request with MaxResults set at 5.
	// The service returns the first batch of results (up to 5) and a NextToken
	// value. To see the next batch of results, you can submit the ListEntitlements
	// request a second time and specify the NextToken value.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListEntitlementsOutput) String() string {
	return awsutil.Prettify(s)
}