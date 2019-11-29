// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetCrawlersInput struct {
	_ struct{} `type:"structure"`

	// The number of crawlers to return on each call.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCrawlersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCrawlersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCrawlersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCrawlersOutput struct {
	_ struct{} `type:"structure"`

	// A list of crawler metadata.
	Crawlers []Crawler `type:"list"`

	// A continuation token, if the returned list has not reached the end of those
	// defined in this customer account.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCrawlersOutput) String() string {
	return awsutil.Prettify(s)
}
