// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure used to list tags for resource.
type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// Resource arn used to list tags.
	//
	// ResourceArn is a required field
	ResourceArn *string `location:"uri" locationName:"resourceArn" type:"string" required:"true"`
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

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for list tags.
type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// Tags result for response.
	Tags map[string]string `locationName:"tags" min:"1" type:"map"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}