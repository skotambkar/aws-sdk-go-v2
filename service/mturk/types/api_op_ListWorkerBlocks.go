// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListWorkerBlocksInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `min:"1" type:"integer"`

	// Pagination token
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListWorkerBlocksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListWorkerBlocksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListWorkerBlocksInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListWorkerBlocksOutput struct {
	_ struct{} `type:"structure"`

	// If the previous response was incomplete (because there is more data to retrieve),
	// Amazon Mechanical Turk returns a pagination token in the response. You can
	// use this pagination token to retrieve the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The number of assignments on the page in the filtered results list, equivalent
	// to the number of assignments returned by this call.
	NumResults *int64 `type:"integer"`

	// The list of WorkerBlocks, containing the collection of Worker IDs and reasons
	// for blocking.
	WorkerBlocks []WorkerBlock `type:"list"`
}

// String returns the string representation
func (s ListWorkerBlocksOutput) String() string {
	return awsutil.Prettify(s)
}