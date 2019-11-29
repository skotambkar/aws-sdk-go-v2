// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListEntitiesInput struct {
	_ struct{} `type:"structure"`

	// The catalog related to the request. Fixed value: AWSMarketplace
	//
	// Catalog is a required field
	Catalog *string `min:"1" type:"string" required:"true"`

	// The type of entities to retrieve.
	//
	// EntityType is a required field
	EntityType *string `min:"1" type:"string" required:"true"`

	// An array of filter objects. Each filter object contains two attributes, filterName
	// and filterValues.
	FilterList []Filter `min:"1" type:"list"`

	// Specifies the upper limit of the elements on a single page. If a value isn't
	// provided, the default value is 20.
	MaxResults *int64 `min:"1" type:"integer"`

	// The value of the next token, if it exists. Null if there are no more results.
	NextToken *string `min:"1" type:"string"`

	// An object that contains two attributes, sortBy and sortOrder.
	Sort *Sort `type:"structure"`
}

// String returns the string representation
func (s ListEntitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEntitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListEntitiesInput"}

	if s.Catalog == nil {
		invalidParams.Add(aws.NewErrParamRequired("Catalog"))
	}
	if s.Catalog != nil && len(*s.Catalog) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Catalog", 1))
	}

	if s.EntityType == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityType"))
	}
	if s.EntityType != nil && len(*s.EntityType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityType", 1))
	}
	if s.FilterList != nil && len(s.FilterList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FilterList", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.FilterList != nil {
		for i, v := range s.FilterList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "FilterList", i), err.(aws.ErrInvalidParams))
			}
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

type ListEntitiesOutput struct {
	_ struct{} `type:"structure"`

	// Array of EntitySummary object.
	EntitySummaryList []EntitySummary `type:"list"`

	// The value of the next token if it exists. Null if there is no more result.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListEntitiesOutput) String() string {
	return awsutil.Prettify(s)
}
