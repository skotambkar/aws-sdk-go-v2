// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ssm/enums"
)

type AddTagsToResourceInput struct {
	_ struct{} `type:"structure"`

	// The resource ID you want to tag.
	//
	// Use the ID of the resource. Here are some examples:
	//
	// ManagedInstance: mi-012345abcde
	//
	// MaintenanceWindow: mw-012345abcde
	//
	// PatchBaseline: pb-012345abcde
	//
	// For the Document and Parameter values, use the name of the resource.
	//
	// The ManagedInstance type for this API action is only for on-premises managed
	// instances. You must specify the name of the managed instance in the following
	// format: mi-ID_number. For example, mi-1a2b3c4d5e6f.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`

	// Specifies the type of resource you are tagging.
	//
	// The ManagedInstance type for this API action is for on-premises managed instances.
	// You must specify the name of the managed instance in the following format:
	// mi-ID_number. For example, mi-1a2b3c4d5e6f.
	//
	// ResourceType is a required field
	ResourceType enums.ResourceTypeForTagging `type:"string" required:"true" enum:"true"`

	// One or more tags. The value parameter is required, but if you don't want
	// the tag to have a value, specify the parameter with no value, and we set
	// the value to an empty string.
	//
	// Do not enter personally identifiable information in this field.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsToResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsToResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsToResourceInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
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

type AddTagsToResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsToResourceOutput) String() string {
	return awsutil.Prettify(s)
}
