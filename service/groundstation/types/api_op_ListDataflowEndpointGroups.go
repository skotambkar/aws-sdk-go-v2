// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDataflowEndpointGroupsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDataflowEndpointGroupsInput) String() string {
	return awsutil.Prettify(s)
}

type ListDataflowEndpointGroupsOutput struct {
	_ struct{} `type:"structure"`

	DataflowEndpointGroupList []DataflowEndpointListItem `locationName:"dataflowEndpointGroupList" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDataflowEndpointGroupsOutput) String() string {
	return awsutil.Prettify(s)
}
