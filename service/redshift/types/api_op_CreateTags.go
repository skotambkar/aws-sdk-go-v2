// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the output from the CreateTags action.
type CreateTagsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) to which you want to add the tag or tags.
	// For example, arn:aws:redshift:us-east-1:123456789:cluster:t1.
	//
	// ResourceName is a required field
	ResourceName *string `type:"string" required:"true"`

	// One or more name/value pairs to add as tags to the specified resource. Each
	// tag name is passed in with the parameter Key and the corresponding value
	// is passed in with the parameter Value. The Key and Value parameters are separated
	// by a comma (,). Separate multiple tags with a space. For example, --tags
	// "Key"="owner","Value"="admin" "Key"="environment","Value"="test" "Key"="version","Value"="1.0".
	//
	// Tags is a required field
	Tags []Tag `locationNameList:"Tag" type:"list" required:"true"`
}

// String returns the string representation
func (s CreateTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTagsInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTagsOutput) String() string {
	return awsutil.Prettify(s)
}