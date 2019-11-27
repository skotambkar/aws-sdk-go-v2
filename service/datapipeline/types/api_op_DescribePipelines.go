// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribePipelines.
type DescribePipelinesInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the pipelines to describe. You can pass as many as 25 identifiers
	// in a single call. To obtain pipeline IDs, call ListPipelines.
	//
	// PipelineIds is a required field
	PipelineIds []string `locationName:"pipelineIds" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribePipelinesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePipelinesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePipelinesInput"}

	if s.PipelineIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of DescribePipelines.
type DescribePipelinesOutput struct {
	_ struct{} `type:"structure"`

	// An array of descriptions for the specified pipelines.
	//
	// PipelineDescriptionList is a required field
	PipelineDescriptionList []PipelineDescription `locationName:"pipelineDescriptionList" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribePipelinesOutput) String() string {
	return awsutil.Prettify(s)
}
