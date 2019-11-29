// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListOriginEndpointsInput struct {
	_ struct{} `type:"structure"`

	ChannelId *string `location:"querystring" locationName:"channelId" type:"string"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListOriginEndpointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOriginEndpointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOriginEndpointsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListOriginEndpointsOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	OriginEndpoints []OriginEndpoint `locationName:"originEndpoints" type:"list"`
}

// String returns the string representation
func (s ListOriginEndpointsOutput) String() string {
	return awsutil.Prettify(s)
}