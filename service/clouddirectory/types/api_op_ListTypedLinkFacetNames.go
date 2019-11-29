// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTypedLinkFacetNamesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTypedLinkFacetNamesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTypedLinkFacetNamesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTypedLinkFacetNamesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTypedLinkFacetNamesOutput struct {
	_ struct{} `type:"structure"`

	// The names of typed link facets that exist within the schema.
	FacetNames []string `type:"list"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListTypedLinkFacetNamesOutput) String() string {
	return awsutil.Prettify(s)
}
