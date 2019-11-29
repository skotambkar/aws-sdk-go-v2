// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeSavingsPlanRatesInput struct {
	_ struct{} `type:"structure"`

	// The filters.
	Filters []SavingsPlanRateFilter `locationName:"filters" type:"list"`

	// The maximum number of results to return with a single call. To retrieve additional
	// results, make another call with the returned token value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The ID of the Savings Plan.
	//
	// SavingsPlanId is a required field
	SavingsPlanId *string `locationName:"savingsPlanId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSavingsPlanRatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSavingsPlanRatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSavingsPlanRatesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.SavingsPlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SavingsPlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSavingsPlanRatesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The ID of the Savings Plan.
	SavingsPlanId *string `locationName:"savingsPlanId" type:"string"`

	// Information about the Savings Plans rates.
	SearchResults []SavingsPlanRate `locationName:"searchResults" type:"list"`
}

// String returns the string representation
func (s DescribeSavingsPlanRatesOutput) String() string {
	return awsutil.Prettify(s)
}