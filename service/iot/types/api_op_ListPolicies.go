// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the ListPolicies operation.
type ListPoliciesInput struct {
	_ struct{} `type:"structure"`

	// Specifies the order for results. If true, the results are returned in ascending
	// creation order.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// The marker for the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The result page size.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPoliciesInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output from the ListPolicies operation.
type ListPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string `locationName:"nextMarker" type:"string"`

	// The descriptions of the policies.
	Policies []Policy `locationName:"policies" type:"list"`
}

// String returns the string representation
func (s ListPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}