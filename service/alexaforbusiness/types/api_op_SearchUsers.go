// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SearchUsersInput struct {
	_ struct{} `type:"structure"`

	// The filters to use for listing a specific set of users. Required. Supported
	// filter keys are UserId, FirstName, LastName, Email, and EnrollmentStatus.
	Filters []Filter `type:"list"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved. Required.
	MaxResults *int64 `min:"1" type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	// Required.
	NextToken *string `min:"1" type:"string"`

	// The sort order to use in listing the filtered set of users. Required. Supported
	// sort keys are UserId, FirstName, LastName, Email, and EnrollmentStatus.
	SortCriteria []Sort `type:"list"`
}

// String returns the string representation
func (s SearchUsersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SearchUsersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SearchUsersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SortCriteria != nil {
		for i, v := range s.SortCriteria {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "SortCriteria", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SearchUsersOutput struct {
	_ struct{} `type:"structure"`

	// The token returned to indicate that there is more data available.
	NextToken *string `min:"1" type:"string"`

	// The total number of users returned.
	TotalCount *int64 `type:"integer"`

	// The users that meet the specified set of filter criteria, in sort order.
	Users []UserData `type:"list"`
}

// String returns the string representation
func (s SearchUsersOutput) String() string {
	return awsutil.Prettify(s)
}
