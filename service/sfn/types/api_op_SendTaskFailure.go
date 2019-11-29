// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SendTaskFailureInput struct {
	_ struct{} `type:"structure"`

	// A more detailed explanation of the cause of the failure.
	Cause *string `locationName:"cause" type:"string" sensitive:"true"`

	// The error code of the failure.
	Error *string `locationName:"error" type:"string" sensitive:"true"`

	// The token that represents this task. Task tokens are generated by Step Functions
	// when tasks are assigned to a worker, or in the context object (https://docs.aws.amazon.com/step-functions/latest/dg/input-output-contextobject.html)
	// when a workflow enters a task state. See GetActivityTaskOutput$taskToken.
	//
	// TaskToken is a required field
	TaskToken *string `locationName:"taskToken" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s SendTaskFailureInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendTaskFailureInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendTaskFailureInput"}

	if s.TaskToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskToken"))
	}
	if s.TaskToken != nil && len(*s.TaskToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TaskToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SendTaskFailureOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SendTaskFailureOutput) String() string {
	return awsutil.Prettify(s)
}
