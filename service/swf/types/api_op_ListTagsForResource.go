// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the Amazon SWF domain.
	//
	// ResourceArn is a required field
	ResourceArn *string `locationName:"resourceArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// An array of tags associated with the domain.
	Tags []ResourceTag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}
