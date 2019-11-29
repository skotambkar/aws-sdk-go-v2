// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConformancePackStatusInput struct {
	_ struct{} `type:"structure"`

	// Comma-separated list of conformance pack names.
	ConformancePackNames []string `type:"list"`

	// The maximum number of conformance packs returned on each page.
	Limit *int64 `type:"integer"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConformancePackStatusInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeConformancePackStatusOutput struct {
	_ struct{} `type:"structure"`

	// A list of ConformancePackStatusDetail objects.
	ConformancePackStatusDetails []ConformancePackStatusDetail `type:"list"`

	// The nextToken string returned in a previous request that you use to request
	// the next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConformancePackStatusOutput) String() string {
	return awsutil.Prettify(s)
}
