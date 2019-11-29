// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListFindingsInput struct {
	_ struct{} `type:"structure"`

	// The ARNs of the assessment runs that generate the findings that you want
	// to list.
	AssessmentRunArns []string `locationName:"assessmentRunArns" type:"list"`

	// You can use this parameter to specify a subset of data to be included in
	// the action's response.
	//
	// For a record to match a filter, all specified filter attributes must match.
	// When multiple values are specified for a filter attribute, any of the values
	// can match.
	Filter *FindingFilter `locationName:"filter" type:"structure"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 10. The maximum value is 500.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListFindings action. Subsequent
	// calls to the action fill nextToken in the request with the value of NextToken
	// from the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListFindingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFindingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFindingsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			invalidParams.AddNested("Filter", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListFindingsOutput struct {
	_ struct{} `type:"structure"`

	// A list of ARNs that specifies the findings returned by the action.
	//
	// FindingArns is a required field
	FindingArns []string `locationName:"findingArns" type:"list" required:"true"`

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to
	// be listed, this parameter is set to null.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListFindingsOutput) String() string {
	return awsutil.Prettify(s)
}
