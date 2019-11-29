// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeProductsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return.
	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	// The token that is required for pagination.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s DescribeProductsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProductsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProductsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeProductsOutput struct {
	_ struct{} `type:"structure"`

	// The token that is required for pagination.
	NextToken *string `type:"string"`

	// A list of products, including details for each product.
	//
	// Products is a required field
	Products []Product `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeProductsOutput) String() string {
	return awsutil.Prettify(s)
}