// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMLTransformsInput struct {
	_ struct{} `type:"structure"`

	// The filter transformation criteria.
	Filter *TransformFilterCriteria `type:"structure"`

	// The maximum number of results to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A paginated token to offset the results.
	NextToken *string `type:"string"`

	// The sorting criteria.
	Sort *TransformSortCriteria `type:"structure"`
}

// String returns the string representation
func (s GetMLTransformsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMLTransformsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMLTransformsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}
	if s.Sort != nil {
		if err := s.Sort.Validate(); err != nil {
			invalidParams.AddNested("Sort", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMLTransformsOutput struct {
	_ struct{} `type:"structure"`

	// A pagination token, if more results are available.
	NextToken *string `type:"string"`

	// A list of machine learning transforms.
	//
	// Transforms is a required field
	Transforms []MLTransform `type:"list" required:"true"`
}

// String returns the string representation
func (s GetMLTransformsOutput) String() string {
	return awsutil.Prettify(s)
}
