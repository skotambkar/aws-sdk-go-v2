// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/enums"
)

// Send an request with an empty body to the regional API endpoint to get your
// account API endpoint.
type DescribeEndpointsInput struct {
	_ struct{} `type:"structure"`

	// Optional. Max number of endpoints, up to twenty, that will be returned at
	// one time.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// Optional field, defaults to DEFAULT. Specify DEFAULT for this operation to
	// return your endpoints if any exist, or to create an endpoint for you and
	// return it if one doesn't already exist. Specify GET_ONLY to return your endpoints
	// if any exist, or an empty list if none exist.
	Mode enums.DescribeEndpointsMode `locationName:"mode" type:"string" enum:"true"`

	// Use this string, provided with the response to a previous request, to request
	// the next batch of endpoints.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Successful describe endpoints requests will return your account API endpoint.
type DescribeEndpointsOutput struct {
	_ struct{} `type:"structure"`

	// List of endpoints
	Endpoints []Endpoint `locationName:"endpoints" type:"list"`

	// Use this string to request the next batch of endpoints.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}