// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// Reserved for future use.
	Limit *int64 `type:"integer"`

	// Reserved for future use.
	NextToken *string `type:"string"`

	// Identifier (ID) of the directory for which you want to retrieve tags.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// Reserved for future use.
	NextToken *string `type:"string"`

	// List of tags returned by the ListTagsForResource operation.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}
