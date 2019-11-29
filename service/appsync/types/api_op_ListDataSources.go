// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDataSourcesInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The maximum number of results you want the request to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDataSourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDataSourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDataSourcesInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListDataSourcesOutput struct {
	_ struct{} `type:"structure"`

	// The DataSource objects.
	DataSources []DataSource `locationName:"dataSources" type:"list"`

	// An identifier to be passed in the next request to this operation to return
	// the next set of items in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDataSourcesOutput) String() string {
	return awsutil.Prettify(s)
}