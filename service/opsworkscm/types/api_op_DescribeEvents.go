// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEventsInput struct {
	_ struct{} `type:"structure"`

	// To receive a paginated response, use this parameter to specify the maximum
	// number of results to be returned with a single call. If the number of available
	// results exceeds this maximum, the response includes a NextToken value that
	// you can assign to the NextToken request parameter to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// NextToken is a string that is returned in some command responses. It indicates
	// that not all entries have been returned, and that you must run at least one
	// more request to get remaining items. To get remaining results, call DescribeEvents
	// again, and assign the token from the previous results as the value of the
	// nextToken parameter. If there are no more results, the response object's
	// nextToken parameter value is null. Setting a nextToken value that was not
	// returned in your previous results causes an InvalidNextTokenException to
	// occur.
	NextToken *string `type:"string"`

	// The name of the server for which you want to view events.
	//
	// ServerName is a required field
	ServerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ServerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerName"))
	}
	if s.ServerName != nil && len(*s.ServerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEventsOutput struct {
	_ struct{} `type:"structure"`

	// NextToken is a string that is returned in some command responses. It indicates
	// that not all entries have been returned, and that you must run at least one
	// more request to get remaining items. To get remaining results, call DescribeEvents
	// again, and assign the token from the previous results as the value of the
	// nextToken parameter. If there are no more results, the response object's
	// nextToken parameter value is null. Setting a nextToken value that was not
	// returned in your previous results causes an InvalidNextTokenException to
	// occur.
	NextToken *string `type:"string"`

	// Contains the response to a DescribeEvents request.
	ServerEvents []ServerEvent `type:"list"`
}

// String returns the string representation
func (s DescribeEventsOutput) String() string {
	return awsutil.Prettify(s)
}