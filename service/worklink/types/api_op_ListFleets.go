// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListFleetsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be included in the next page.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token used to retrieve the next page of results for this operation.
	// If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFleetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFleetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFleetsInput"}
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

type ListFleetsOutput struct {
	_ struct{} `type:"structure"`

	// The summary list of the fleets.
	FleetSummaryList []FleetSummary `type:"list"`

	// The pagination token used to retrieve the next page of results for this operation.
	// If there are no more pages, this value is null.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFleetsOutput) String() string {
	return awsutil.Prettify(s)
}
