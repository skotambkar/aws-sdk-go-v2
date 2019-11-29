// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetWorkflowInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether to include a graph when returning the workflow resource
	// metadata.
	IncludeGraph *bool `type:"boolean"`

	// The name of the workflow to retrieve.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetWorkflowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWorkflowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWorkflowInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetWorkflowOutput struct {
	_ struct{} `type:"structure"`

	// The resource metadata for the workflow.
	Workflow *Workflow `type:"structure"`
}

// String returns the string representation
func (s GetWorkflowOutput) String() string {
	return awsutil.Prettify(s)
}
