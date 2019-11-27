// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RespondActivityTaskCanceledInput struct {
	_ struct{} `type:"structure"`

	// Information about the cancellation.
	Details *string `locationName:"details" type:"string"`

	// The taskToken of the ActivityTask.
	//
	// taskToken is generated by the service and should be treated as an opaque
	// value. If the task is passed to another process, its taskToken must also
	// be passed. This enables it to provide its progress and respond with results.
	//
	// TaskToken is a required field
	TaskToken *string `locationName:"taskToken" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RespondActivityTaskCanceledInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RespondActivityTaskCanceledInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RespondActivityTaskCanceledInput"}

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

type RespondActivityTaskCanceledOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RespondActivityTaskCanceledOutput) String() string {
	return awsutil.Prettify(s)
}
