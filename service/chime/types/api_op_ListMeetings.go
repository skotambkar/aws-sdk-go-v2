// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListMeetingsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The token to use to retrieve the next page of results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`
}

// String returns the string representation
func (s ListMeetingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMeetingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMeetingsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListMeetingsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime SDK meeting information.
	Meetings []Meeting `type:"list"`

	// The token to use to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMeetingsOutput) String() string {
	return awsutil.Prettify(s)
}