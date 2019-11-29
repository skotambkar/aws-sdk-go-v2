// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// List the tags for your AWS Elemental MediaConvert resource by sending a request
// with the Amazon Resource Name (ARN) of the resource. To get the ARN, send
// a GET request with the resource name.
type ListTagsForResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource that you want to list tags
	// for. To get the ARN, send a GET request with the resource name.
	//
	// Arn is a required field
	Arn *string `location:"uri" locationName:"arn" type:"string" required:"true"`
}

// String returns the string representation
func (s ListTagsForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTagsForResourceInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A successful request to list the tags for a resource returns a JSON map of
// tags.
type ListTagsForResourceOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) and tags for an AWS Elemental MediaConvert
	// resource.
	ResourceTags *ResourceTags `locationName:"resourceTags" type:"structure"`
}

// String returns the string representation
func (s ListTagsForResourceOutput) String() string {
	return awsutil.Prettify(s)
}