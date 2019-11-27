// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetStatisticsInput struct {
	_ struct{} `type:"structure"`

	// The aggregation field name. Currently not supported.
	AggregationField *string `locationName:"aggregationField" min:"1" type:"string"`

	// The name of the index to search. The default value is AWS_Things.
	IndexName *string `locationName:"indexName" min:"1" type:"string"`

	// The query used to search. You can specify "*" for the query string to get
	// the count of all indexed things in your AWS account.
	//
	// QueryString is a required field
	QueryString *string `locationName:"queryString" min:"1" type:"string" required:"true"`

	// The version of the query used to search.
	QueryVersion *string `locationName:"queryVersion" type:"string"`
}

// String returns the string representation
func (s GetStatisticsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetStatisticsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetStatisticsInput"}
	if s.AggregationField != nil && len(*s.AggregationField) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AggregationField", 1))
	}
	if s.IndexName != nil && len(*s.IndexName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IndexName", 1))
	}

	if s.QueryString == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryString"))
	}
	if s.QueryString != nil && len(*s.QueryString) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QueryString", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetStatisticsOutput struct {
	_ struct{} `type:"structure"`

	// The statistics returned by the Fleet Indexing service based on the query
	// and aggregation field.
	Statistics *Statistics `locationName:"statistics" type:"structure"`
}

// String returns the string representation
func (s GetStatisticsOutput) String() string {
	return awsutil.Prettify(s)
}
