// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for RemoveTags.
type RemoveTagsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the pipeline.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`

	// The keys of the tags to remove.
	//
	// TagKeys is a required field
	TagKeys []string `locationName:"tagKeys" type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveTagsInput"}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}
	if s.PipelineId != nil && len(*s.PipelineId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineId", 1))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of RemoveTags.
type RemoveTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsOutput) String() string {
	return awsutil.Prettify(s)
}
