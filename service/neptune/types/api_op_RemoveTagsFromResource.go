// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RemoveTagsFromResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Neptune resource that the tags are removed from. This value is
	// an Amazon Resource Name (ARN). For information about creating an ARN, see
	// Constructing an Amazon Resource Name (ARN) (https://docs.aws.amazon.com/neptune/latest/UserGuide/tagging.ARN.html#tagging.ARN.Constructing).
	//
	// ResourceName is a required field
	ResourceName *string `type:"string" required:"true"`

	// The tag key (name) of the tag to be removed.
	//
	// TagKeys is a required field
	TagKeys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsFromResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveTagsFromResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveTagsFromResourceInput"}

	if s.ResourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceName"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RemoveTagsFromResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsFromResourceOutput) String() string {
	return awsutil.Prettify(s)
}