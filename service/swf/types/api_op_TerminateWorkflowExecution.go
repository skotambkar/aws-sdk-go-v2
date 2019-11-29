// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/swf/enums"
)

type TerminateWorkflowExecutionInput struct {
	_ struct{} `type:"structure"`

	// If set, specifies the policy to use for the child workflow executions of
	// the workflow execution being terminated. This policy overrides the child
	// policy specified for the workflow execution at registration time or when
	// starting the execution.
	//
	// The supported child policies are:
	//
	//    * TERMINATE – The child executions are terminated.
	//
	//    * REQUEST_CANCEL – A request to cancel is attempted for each child execution
	//    by recording a WorkflowExecutionCancelRequested event in its history.
	//    It is up to the decider to take appropriate actions when it receives an
	//    execution history with this event.
	//
	//    * ABANDON – No action is taken. The child executions continue to run.
	//
	// A child policy for this workflow execution must be specified either as a
	// default for the workflow type or through this parameter. If neither this
	// parameter is set nor a default child policy was specified at registration
	// time then a fault is returned.
	ChildPolicy enums.ChildPolicy `locationName:"childPolicy" type:"string" enum:"true"`

	// Details for terminating the workflow execution.
	Details *string `locationName:"details" type:"string"`

	// The domain of the workflow execution to terminate.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// A descriptive reason for terminating the workflow execution.
	Reason *string `locationName:"reason" type:"string"`

	// The runId of the workflow execution to terminate.
	RunId *string `locationName:"runId" type:"string"`

	// The workflowId of the workflow execution to terminate.
	//
	// WorkflowId is a required field
	WorkflowId *string `locationName:"workflowId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TerminateWorkflowExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateWorkflowExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateWorkflowExecutionInput"}

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

type TerminateWorkflowExecutionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TerminateWorkflowExecutionOutput) String() string {
	return awsutil.Prettify(s)
}