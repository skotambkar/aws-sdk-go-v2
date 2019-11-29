// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStateMachineForExecutionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the execution you want state machine information
	// for.
	//
	// ExecutionArn is a required field
	ExecutionArn *string `locationName:"executionArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStateMachineForExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStateMachineForExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStateMachineForExecutionInput"}

	if s.ExecutionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExecutionArn"))
	}
	if s.ExecutionArn != nil && len(*s.ExecutionArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExecutionArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStateMachineForExecutionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html).
	//
	// Definition is a required field
	Definition *string `locationName:"definition" min:"1" type:"string" required:"true" sensitive:"true"`

	// The name of the state machine associated with the execution.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the IAM role of the State Machine for the
	// execution.
	//
	// RoleArn is a required field
	RoleArn *string `locationName:"roleArn" min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the state machine associated with the execution.
	//
	// StateMachineArn is a required field
	StateMachineArn *string `locationName:"stateMachineArn" min:"1" type:"string" required:"true"`

	// The date and time the state machine associated with an execution was updated.
	// For a newly created state machine, this is the creation date.
	//
	// UpdateDate is a required field
	UpdateDate *time.Time `locationName:"updateDate" type:"timestamp" required:"true"`
}

// String returns the string representation
func (s DescribeStateMachineForExecutionOutput) String() string {
	return awsutil.Prettify(s)
}
