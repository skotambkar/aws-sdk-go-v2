// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RequestCancelWorkflowExecutionInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain containing the workflow execution to cancel.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The runId of the workflow execution to cancel.
	RunId *string `locationName:"runId" type:"string"`

	// The workflowId of the workflow execution to cancel.
	//
	// WorkflowId is a required field
	WorkflowId *string `locationName:"workflowId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RequestCancelWorkflowExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestCancelWorkflowExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestCancelWorkflowExecutionInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.WorkflowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkflowId"))
	}
	if s.WorkflowId != nil && len(*s.WorkflowId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkflowId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RequestCancelWorkflowExecutionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RequestCancelWorkflowExecutionOutput) String() string {
	return awsutil.Prettify(s)
}