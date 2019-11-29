// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TagResourceInput struct {
	_ struct{} `type:"structure"`

	// The cluster identifier (ID) for the cluster that you are tagging. To find
	// the cluster ID, use DescribeClusters.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`

	// A list of one or more tags.
	//
	// TagList is a required field
	TagList []Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s TagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagResourceInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.TagList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagList"))
	}
	if s.TagList != nil && len(s.TagList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TagList", 1))
	}
	if s.TagList != nil {
		for i, v := range s.TagList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagResourceOutput) String() string {
	return awsutil.Prettify(s)
}
