// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTargetsInput struct {
	_ struct{} `type:"structure"`

	// The filters to use to return information by service or resource type. Valid
	// filters include target type, target address, and target status.
	//
	// A filter with the same name can appear more than once when used with OR statements.
	// Filters with different names should be applied with AND statements.
	Filters []ListTargetsFilter `type:"list"`

	// A non-negative integer used to limit the number of returned results. The
	// maximum number of results that can be returned is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTargetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
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

type ListTargetsOutput struct {
	_ struct{} `type:"structure"`

	// An enumeration token that can be used in a request to return the next batch
	// of results.
	NextToken *string `type:"string"`

	// The list of notification rule targets.
	Targets []TargetSummary `type:"list"`
}

// String returns the string representation
func (s ListTargetsOutput) String() string {
	return awsutil.Prettify(s)
}
