// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/enums"
)

type ListIncomingTypedLinksInput struct {
	_ struct{} `type:"structure"`

	// The consistency level to execute the request at.
	ConsistencyLevel enums.ConsistencyLevel `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the directory where you want to list the
	// typed links.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Provides range filters for multiple attributes. When providing ranges to
	// typed link selection, any inexact ranges must be specified at the end. Any
	// attributes that do not have a range specified are presumed to match the entire
	// range.
	FilterAttributeRanges []TypedLinkAttributeRange `type:"list"`

	// Filters are interpreted in the order of the attributes on the typed link
	// facet, not the order in which they are supplied to any API calls.
	FilterTypedLink *TypedLinkSchemaAndFacetName `type:"structure"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// Reference that identifies the object whose attributes will be listed.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListIncomingTypedLinksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListIncomingTypedLinksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListIncomingTypedLinksInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}
	if s.FilterAttributeRanges != nil {
		for i, v := range s.FilterAttributeRanges {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "FilterAttributeRanges", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.FilterTypedLink != nil {
		if err := s.FilterTypedLink.Validate(); err != nil {
			invalidParams.AddNested("FilterTypedLink", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListIncomingTypedLinksOutput struct {
	_ struct{} `type:"structure"`

	// Returns one or more typed link specifiers as output.
	LinkSpecifiers []TypedLinkSpecifier `type:"list"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListIncomingTypedLinksOutput) String() string {
	return awsutil.Prettify(s)
}