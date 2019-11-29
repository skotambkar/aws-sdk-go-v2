// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListMemberAccountsInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter to indicate the maximum number of items that you want
	// in the response. The default value is 250.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// Use this parameter when paginating results. Set the value of this parameter
	// to null on your first call to the ListMemberAccounts action. Subsequent calls
	// to the action fill nextToken in the request with the value of nextToken from
	// the previous response to continue listing data.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListMemberAccountsInput) String() string {
	return awsutil.Prettify(s)
}

type ListMemberAccountsOutput struct {
	_ struct{} `type:"structure"`

	// A list of the Amazon Macie member accounts returned by the action. The current
	// master account is also included in this list.
	MemberAccounts []MemberAccount `locationName:"memberAccounts" type:"list"`

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to
	// be listed, this parameter is set to null.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListMemberAccountsOutput) String() string {
	return awsutil.Prettify(s)
}