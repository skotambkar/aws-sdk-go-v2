// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/organizations/enums"
)

type ListChildrenInput struct {
	_ struct{} `type:"structure"`

	// Filters the output to include only the specified child type.
	//
	// ChildType is a required field
	ChildType enums.ChildType `type:"string" required:"true" enum:"true"`

	// (Optional) Use this to limit the number of results you want included per
	// page in the response. If you do not include this parameter, it defaults to
	// a value that is specific to the operation. If additional items exist beyond
	// the maximum you specify, the NextToken response element is present and has
	// a value (is not null). Include that value as the NextToken request parameter
	// in the next call to the operation to get the next part of the results. Note
	// that Organizations might return fewer results than the maximum even when
	// there are more results available. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// Use this parameter if you receive a NextToken response in a previous request
	// that indicates that there is more output available. Set it to the value of
	// the previous call's NextToken response to indicate where the output should
	// continue from.
	NextToken *string `type:"string"`

	// The unique identifier (ID) for the parent root or OU whose children you want
	// to list.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) for a parent ID string
	// requires one of the following:
	//
	//    * Root - A string that begins with "r-" followed by from 4 to 32 lower-case
	//    letters or digits.
	//
	//    * Organizational unit (OU) - A string that begins with "ou-" followed
	//    by from 4 to 32 lower-case letters or digits (the ID of the root that
	//    the OU is in) followed by a second "-" dash and from 8 to 32 additional
	//    lower-case letters or digits.
	//
	// ParentId is a required field
	ParentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListChildrenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListChildrenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListChildrenInput"}
	if len(s.ChildType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ChildType"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ParentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListChildrenOutput struct {
	_ struct{} `type:"structure"`

	// The list of children of the specified parent container.
	Children []Child `type:"list"`

	// If present, this value indicates that there is more output available than
	// is included in the current response. Use this value in the NextToken request
	// parameter in a subsequent call to the operation to get the next part of the
	// output. You should repeat this until the NextToken response element comes
	// back as null.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListChildrenOutput) String() string {
	return awsutil.Prettify(s)
}