// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The GET request to get all the usage plans of the caller's account.
type GetUsagePlansInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the API key associated with the usage plans.
	KeyId *string `location:"querystring" locationName:"keyId" type:"string"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetUsagePlansInput) String() string {
	return awsutil.Prettify(s)
}

// Represents a collection of usage plans for an AWS account.
//
// Create and Use Usage Plans (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlansOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []UsagePlan `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetUsagePlansOutput) String() string {
	return awsutil.Prettify(s)
}
