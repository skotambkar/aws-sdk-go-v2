// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchGetTracesInput struct {
	_ struct{} `type:"structure"`

	// Pagination token. Not used.
	NextToken *string `type:"string"`

	// Specify the trace IDs of requests for which to retrieve segments.
	//
	// TraceIds is a required field
	TraceIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetTracesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetTracesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetTracesInput"}

	if s.TraceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("TraceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchGetTracesOutput struct {
	_ struct{} `type:"structure"`

	// Pagination token. Not used.
	NextToken *string `type:"string"`

	// Full traces for the specified requests.
	Traces []Trace `type:"list"`

	// Trace IDs of requests that haven't been processed.
	UnprocessedTraceIds []string `type:"list"`
}

// String returns the string representation
func (s BatchGetTracesOutput) String() string {
	return awsutil.Prettify(s)
}