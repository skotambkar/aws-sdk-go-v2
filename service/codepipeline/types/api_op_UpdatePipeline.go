// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of an UpdatePipeline action.
type UpdatePipelineInput struct {
	_ struct{} `type:"structure"`

	// The name of the pipeline to be updated.
	//
	// Pipeline is a required field
	Pipeline *PipelineDeclaration `locationName:"pipeline" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdatePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePipelineInput"}

	if s.Pipeline == nil {
		invalidParams.Add(aws.NewErrParamRequired("Pipeline"))
	}
	if s.Pipeline != nil {
		if err := s.Pipeline.Validate(); err != nil {
			invalidParams.AddNested("Pipeline", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an UpdatePipeline action.
type UpdatePipelineOutput struct {
	_ struct{} `type:"structure"`

	// The structure of the updated pipeline.
	Pipeline *PipelineDeclaration `locationName:"pipeline" type:"structure"`
}

// String returns the string representation
func (s UpdatePipelineOutput) String() string {
	return awsutil.Prettify(s)
}
