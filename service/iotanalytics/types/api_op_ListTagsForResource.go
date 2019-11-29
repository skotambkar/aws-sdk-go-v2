// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource whose tags you want to list.
	//
	// ResourceArn is a required field
	ResourceArn *string `location:"querystring" locationName:"resourceArn" min:"20" type:"string" required:"true"`
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
	if s.ResourceArn != nil && len(*s.ResourceArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// The tags (metadata) which you have assigned to the resource.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}
