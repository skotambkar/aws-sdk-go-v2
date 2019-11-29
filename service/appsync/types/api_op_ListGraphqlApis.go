// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListGraphqlApisInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results you want the request to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListGraphqlApisInput) String() string {
	return awsutil.Prettify(s)
}

type ListGraphqlApisOutput struct {
	_ struct{} `type:"structure"`

	// The GraphqlApi objects.
	GraphqlApis []GraphqlApi `locationName:"graphqlApis" type:"list"`

	// An identifier to be passed in the next request to this operation to return
	// the next set of items in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListGraphqlApisOutput) String() string {
	return awsutil.Prettify(s)
}