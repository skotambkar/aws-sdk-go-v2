// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopQueryExecutionInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the query execution to stop.
	//
	// QueryExecutionId is a required field
	QueryExecutionId *string `type:"string" required:"true" idempotencyToken:"true"`
}

// String returns the string representation
func (s StopQueryExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopQueryExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopQueryExecutionInput"}

	if s.QueryExecutionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryExecutionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopQueryExecutionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopQueryExecutionOutput) String() string {
	return awsutil.Prettify(s)
}
