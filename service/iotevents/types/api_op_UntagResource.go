// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UntagResourceInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the resource.
	//
	// ResourceArn is a required field
	ResourceArn *string `location:"querystring" locationName:"resourceArn" min:"1" type:"string" required:"true"`

	// A list of the keys of the tags to be removed from the resource.
	//
	// TagKeys is a required field
	TagKeys []string `location:"querystring" locationName:"tagKeys" type:"list" required:"true"`
}

// String returns the string representation
func (s UntagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UntagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagResourceOutput) String() string {
	return awsutil.Prettify(s)
}
