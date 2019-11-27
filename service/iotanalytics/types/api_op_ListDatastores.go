// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDatastoresInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in this request.
	//
	// The default value is 100.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatastoresInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDatastoresInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDatastoresInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDatastoresOutput struct {
	_ struct{} `type:"structure"`

	// A list of "DatastoreSummary" objects.
	DatastoreSummaries []DatastoreSummary `locationName:"datastoreSummaries" type:"list"`

	// The token to retrieve the next set of results, or null if there are no more
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDatastoresOutput) String() string {
	return awsutil.Prettify(s)
}
