// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateWorkflowInput struct {
	_ struct{} `type:"structure"`

	// A collection of properties to be used as part of each execution of the workflow.
	DefaultRunProperties map[string]string `type:"map"`

	// The description of the workflow.
	Description *string `type:"string"`

	// Name of the workflow to be updated.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateWorkflowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateWorkflowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateWorkflowInput"}

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

type UpdateWorkflowOutput struct {
	_ struct{} `type:"structure"`

	// The name of the workflow which was specified in input.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateWorkflowOutput) String() string {
	return awsutil.Prettify(s)
}
