// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/swf/enums"
)

type RegisterWorkflowTypeInput struct {
	_ struct{} `type:"structure"`

	// If set, specifies the default policy to use for the child workflow executions
	// when a workflow execution of this type is terminated, by calling the TerminateWorkflowExecution
	// action explicitly or due to an expired timeout. This default can be overridden
	// when starting a workflow execution using the StartWorkflowExecution action
	// or the StartChildWorkflowExecution Decision.
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
	DefaultChildPolicy enums.ChildPolicy `locationName:"defaultChildPolicy" type:"string" enum:"true"`

	// If set, specifies the default maximum duration for executions of this workflow
	// type. You can override this default when starting an execution through the
	// StartWorkflowExecution Action or StartChildWorkflowExecution Decision.
	//
	// The duration is specified in seconds; an integer greater than or equal to
	// 0. Unlike some of the other timeout parameters in Amazon SWF, you cannot
	// specify a value of "NONE" for defaultExecutionStartToCloseTimeout; there
	// is a one-year max limit on the time that a workflow execution can run. Exceeding
	// this limit always causes the workflow execution to time out.
	DefaultExecutionStartToCloseTimeout *string `locationName:"defaultExecutionStartToCloseTimeout" type:"string"`

	// The default IAM role attached to this workflow type.
	//
	// Executions of this workflow type need IAM roles to invoke Lambda functions.
	// If you don't specify an IAM role when you start this workflow type, the default
	// Lambda role is attached to the execution. For more information, see https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html)
	// in the Amazon SWF Developer Guide.
	DefaultLambdaRole *string `locationName:"defaultLambdaRole" min:"1" type:"string"`

	// If set, specifies the default task list to use for scheduling decision tasks
	// for executions of this workflow type. This default is used only if a task
	// list isn't provided when starting the execution through the StartWorkflowExecution
	// Action or StartChildWorkflowExecution Decision.
	DefaultTaskList *TaskList `locationName:"defaultTaskList" type:"structure"`

	// The default task priority to assign to the workflow type. If not assigned,
	// then 0 is used. Valid values are integers that range from Java's Integer.MIN_VALUE
	// (-2147483648) to Integer.MAX_VALUE (2147483647). Higher numbers indicate
	// higher priority.
	//
	// For more information about setting task priority, see Setting Task Priority
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the Amazon SWF Developer Guide.
	DefaultTaskPriority *string `locationName:"defaultTaskPriority" type:"string"`

	// If set, specifies the default maximum duration of decision tasks for this
	// workflow type. This default can be overridden when starting a workflow execution
	// using the StartWorkflowExecution action or the StartChildWorkflowExecution
	// Decision.
	//
	// The duration is specified in seconds, an integer greater than or equal to
	// 0. You can use NONE to specify unlimited duration.
	DefaultTaskStartToCloseTimeout *string `locationName:"defaultTaskStartToCloseTimeout" type:"string"`

	// Textual description of the workflow type.
	Description *string `locationName:"description" type:"string"`

	// The name of the domain in which to register the workflow type.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The name of the workflow type.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not be the literal string arn.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The version of the workflow type.
	//
	// The workflow type consists of the name and version, the combination of which
	// must be unique within the domain. To get a list of all currently registered
	// workflow types, use the ListWorkflowTypes action.
	//
	// The specified string must not start or end with whitespace. It must not contain
	// a : (colon), / (slash), | (vertical bar), or any control characters (\u0000-\u001f
	// | \u007f-\u009f). Also, it must not be the literal string arn.
	//
	// Version is a required field
	Version *string `locationName:"version" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterWorkflowTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterWorkflowTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterWorkflowTypeInput"}
	if s.DefaultLambdaRole != nil && len(*s.DefaultLambdaRole) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultLambdaRole", 1))
	}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Version == nil {
		invalidParams.Add(aws.NewErrParamRequired("Version"))
	}
	if s.Version != nil && len(*s.Version) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Version", 1))
	}
	if s.DefaultTaskList != nil {
		if err := s.DefaultTaskList.Validate(); err != nil {
			invalidParams.AddNested("DefaultTaskList", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterWorkflowTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterWorkflowTypeOutput) String() string {
	return awsutil.Prettify(s)
}
