// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateStateMachineInput struct {
	_ struct{} `type:"structure"`

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	Definition *string `locationName:"definition" min:"1" type:"string" sensitive:"true"`

	// The Amazon Resource Name (ARN) of the IAM role of the state machine.
	RoleArn *string `locationName:"roleArn" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the state machine.
	//
	// StateMachineArn is a required field
	StateMachineArn *string `locationName:"stateMachineArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateStateMachineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateStateMachineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateStateMachineInput"}
	if s.Definition != nil && len(*s.Definition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Definition", 1))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 1))
	}

	if s.StateMachineArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("StateMachineArn"))
	}
	if s.StateMachineArn != nil && len(*s.StateMachineArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StateMachineArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateStateMachineOutput struct {
	_ struct{} `type:"structure"`

	// The date and time the state machine was updated.
	//
	// UpdateDate is a required field
	UpdateDate *time.Time `locationName:"updateDate" type:"timestamp" required:"true"`
}

// String returns the string representation
func (s UpdateStateMachineOutput) String() string {
	return awsutil.Prettify(s)
}