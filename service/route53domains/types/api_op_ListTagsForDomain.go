// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The ListTagsForDomainRequest includes the following elements.
type ListTagsForDomainInput struct {
	_ struct{} `type:"structure"`

	// The domain for which you want to get a list of tags.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The ListTagsForDomain response includes the following elements.
type ListTagsForDomainOutput struct {
	_ struct{} `type:"structure"`

	// A list of the tags that are associated with the specified domain.
	//
	// TagList is a required field
	TagList []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s ListTagsForDomainOutput) String() string {
	return awsutil.Prettify(s)
}