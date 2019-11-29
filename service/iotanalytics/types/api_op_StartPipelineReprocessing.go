// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartPipelineReprocessingInput struct {
	_ struct{} `type:"structure"`

	// The end time (exclusive) of raw message data that is reprocessed.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The name of the pipeline on which to start reprocessing.
	//
	// PipelineName is a required field
	PipelineName *string `location:"uri" locationName:"pipelineName" min:"1" type:"string" required:"true"`

	// The start time (inclusive) of raw message data that is reprocessed.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`
}

// String returns the string representation
func (s StartPipelineReprocessingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartPipelineReprocessingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartPipelineReprocessingInput"}

	if s.PipelineName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineName"))
	}
	if s.PipelineName != nil && len(*s.PipelineName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PipelineName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartPipelineReprocessingOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the pipeline reprocessing activity that was started.
	ReprocessingId *string `locationName:"reprocessingId" type:"string"`
}

// String returns the string representation
func (s StartPipelineReprocessingOutput) String() string {
	return awsutil.Prettify(s)
}
