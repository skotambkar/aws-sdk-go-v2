// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/enums"
)

type AddTagsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the ML object to tag. For example, exampleModelId.
	//
	// ResourceId is a required field
	ResourceId *string `min:"1" type:"string" required:"true"`

	// The type of the ML object to tag.
	//
	// ResourceType is a required field
	ResourceType enums.TaggableResourceType `type:"string" required:"true" enum:"true"`

	// The key-value pairs to use to create tags. If you specify a key without specifying
	// a value, Amazon ML creates a tag with the specified key and a value of null.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}
	if s.ResourceId != nil && len(*s.ResourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceId", 1))
	}
	if len(s.ResourceType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Amazon ML returns the following elements.
type AddTagsOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the ML object that was tagged.
	ResourceId *string `min:"1" type:"string"`

	// The type of the ML object that was tagged.
	ResourceType enums.TaggableResourceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s AddTagsOutput) String() string {
	return awsutil.Prettify(s)
}
