// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetTraceGraphInput struct {
	_ struct{} `type:"structure"`

	// Pagination token. Not used.
	NextToken *string `type:"string"`

	// Trace IDs of requests for which to generate a service graph.
	//
	// TraceIds is a required field
	TraceIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s GetTraceGraphInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTraceGraphInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTraceGraphInput"}

	if s.TraceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("TraceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTraceGraphOutput struct {
	_ struct{} `type:"structure"`

	// Pagination token. Not used.
	NextToken *string `type:"string"`

	// The services that have processed one of the specified requests.
	Services []Service `type:"list"`
}

// String returns the string representation
func (s GetTraceGraphOutput) String() string {
	return awsutil.Prettify(s)
}