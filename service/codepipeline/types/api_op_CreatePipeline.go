// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreatePipeline action.
type CreatePipelineInput struct {
	_ struct{} `type:"structure"`

	// Represents the structure of actions and stages to be performed in the pipeline.
	//
	// Pipeline is a required field
	Pipeline *PipelineDeclaration `locationName:"pipeline" type:"structure" required:"true"`

	// The tags for the pipeline.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreatePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePipelineInput"}

	if s.Pipeline == nil {
		invalidParams.Add(aws.NewErrParamRequired("Pipeline"))
	}
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			invalidParams.AddNested("Pipeline", err.(aws.ErrInvalidParams))
		}
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

// Represents the output of a CreatePipeline action.
type CreatePipelineOutput struct {
	_ struct{} `type:"structure"`

	// Represents the structure of actions and stages to be performed in the pipeline.
	Pipeline *PipelineDeclaration `locationName:"pipeline" type:"structure"`

	// Specifies the tags applied to the pipeline.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreatePipelineOutput) String() string {
	return awsutil.Prettify(s)
}