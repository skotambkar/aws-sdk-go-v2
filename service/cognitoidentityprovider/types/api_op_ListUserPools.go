// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to list user pools.
type ListUserPoolsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results you want the request to return when listing
	// the user pools.
	//
	// MaxResults is a required field
	MaxResults *int64 `min:"1" type:"integer" required:"true"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListUserPoolsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUserPoolsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUserPoolsInput"}

	if s.MaxResults == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaxResults"))
	}
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

// Represents the response to list user pools.
type ListUserPoolsOutput struct {
	_ struct{} `type:"structure"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `min:"1" type:"string"`

	// The user pools from the response to list users.
	UserPools []UserPoolDescriptionType `type:"list"`
}

// String returns the string representation
func (s ListUserPoolsOutput) String() string {
	return awsutil.Prettify(s)
}