// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDatasetImportJobsInput struct {
	_ struct{} `type:"structure"`

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether
	// to include or exclude, respectively, from the list, the predictors that match
	// the statement. The match statement consists of a key and a value. In this
	// release, Name is the only valid key, which filters on the DatasetImportJobName
	// property.
	//
	//    * Condition - IS or IS_NOT
	//
	//    * Key - Name
	//
	//    * Value - the value to match
	//
	// For example, to list all dataset import jobs named my_dataset_import_job,
	// you would specify:
	//
	// "Filters": [ { "Condition": "IS", "Key": "Name", "Value": "my_dataset_import_job"
	// } ]
	Filters []Filter `type:"list"`

	// The number of items to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the result of the previous request was truncated, the response includes
	// a NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDatasetImportJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatasetImportJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatasetImportJobsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDatasetImportJobsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that summarize each dataset import job's properties.
	DatasetImportJobs []DatasetImportJobSummary `type:"list"`

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDatasetImportJobsOutput) String() string {
	return awsutil.Prettify(s)
}
