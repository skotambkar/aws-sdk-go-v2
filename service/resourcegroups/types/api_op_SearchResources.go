// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SearchResourcesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of group member ARNs returned by SearchResources in paginated
	// output. By default, this number is 50.
	MaxResults *int64 `min:"1" type:"integer"`

	// The NextToken value that is returned in a paginated SearchResources request.
	// To get the next page of results, run the call again, add the NextToken parameter,
	// and specify the NextToken value.
	NextToken *string `type:"string"`

	// The search query, using the same formats that are supported for resource
	// group definition.
	//
	// ResourceQuery is a required field
	ResourceQuery *ResourceQuery `type:"structure" required:"true"`
}

// String returns the string representation
func (s SearchResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchResourcesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ResourceQuery == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceQuery"))
	}
	if s.ResourceQuery != nil {
		if err := s.ResourceQuery.Validate(); err != nil {
			invalidParams.AddNested("ResourceQuery", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SearchResourcesOutput struct {
	_ struct{} `type:"structure"`

	// The NextToken value to include in a subsequent SearchResources request, to
	// get more results.
	NextToken *string `type:"string"`

	// A list of QueryError objects. Each error is an object that contains ErrorCode
	// and Message structures. Possible values for ErrorCode are CLOUDFORMATION_STACK_INACTIVE
	// and CLOUDFORMATION_STACK_NOT_EXISTING.
	QueryErrors []QueryError `type:"list"`

	// The ARNs and resource types of resources that are members of the group that
	// you specified.
	ResourceIdentifiers []ResourceIdentifier `type:"list"`
}

// String returns the string representation
func (s SearchResourcesOutput) String() string {
	return awsutil.Prettify(s)
}
