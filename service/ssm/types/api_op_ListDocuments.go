// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDocumentsInput struct {
	_ struct{} `type:"structure"`

	// One or more filters. Use a filter to return a more specific list of results.
	DocumentFilterList []DocumentFilter `min:"1" type:"list"`

	// One or more filters. Use a filter to return a more specific list of results.
	Filters []DocumentKeyValuesFilter `type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDocumentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDocumentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDocumentsInput"}
	if s.DocumentFilterList != nil && len(s.DocumentFilterList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentFilterList", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.DocumentFilterList != nil {
		for i, v := range s.DocumentFilterList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DocumentFilterList", i), err.(aws.ErrInvalidParams))
			}
		}
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

type ListDocumentsOutput struct {
	_ struct{} `type:"structure"`

	// The names of the Systems Manager documents.
	DocumentIdentifiers []DocumentIdentifier `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDocumentsOutput) String() string {
	return awsutil.Prettify(s)
}
