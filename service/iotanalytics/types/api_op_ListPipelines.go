// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListPipelinesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in this request.
	//
	// The default value is 100.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListPipelinesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPipelinesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPipelinesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListPipelinesOutput struct {
	_ struct{} `type:"structure"`

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of "PipelineSummary" objects.
	PipelineSummaries []PipelineSummary `locationName:"pipelineSummaries" type:"list"`
}

// String returns the string representation
func (s ListPipelinesOutput) String() string {
	return awsutil.Prettify(s)
}
