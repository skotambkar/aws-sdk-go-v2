// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetResourcePoliciesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The principal.
	Principal *string `locationName:"principal" type:"string"`

	// The Amazon Resource Names (ARN) of the resources.
	//
	// ResourceArns is a required field
	ResourceArns []string `locationName:"resourceArns" type:"list" required:"true"`
}

// String returns the string representation
func (s GetResourcePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourcePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourcePoliciesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ResourceArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArns"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetResourcePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A key policy document, in JSON format.
	Policies []string `locationName:"policies" type:"list"`
}

// String returns the string representation
func (s GetResourcePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}
