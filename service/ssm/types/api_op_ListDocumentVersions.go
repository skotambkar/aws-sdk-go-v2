// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDocumentVersionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The name of the document about which you want version information.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDocumentVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDocumentVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDocumentVersionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDocumentVersionsOutput struct {
	_ struct{} `type:"structure"`

	// The document versions.
	DocumentVersions []DocumentVersionInfo `min:"1" type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDocumentVersionsOutput) String() string {
	return awsutil.Prettify(s)
}