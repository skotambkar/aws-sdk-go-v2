// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetWorkflowRunInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether to include the workflow graph in response or not.
	IncludeGraph *bool `type:"boolean"`

	// Name of the workflow being run.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The ID of the workflow run.
	//
	// RunId is a required field
	RunId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetWorkflowRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWorkflowRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWorkflowRunInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.RunId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RunId"))
	}
	if s.RunId != nil && len(*s.RunId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RunId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetWorkflowRunOutput struct {
	_ struct{} `type:"structure"`

	// The requested workflow run metadata.
	Run *WorkflowRun `type:"structure"`
}

// String returns the string representation
func (s GetWorkflowRunOutput) String() string {
	return awsutil.Prettify(s)
}