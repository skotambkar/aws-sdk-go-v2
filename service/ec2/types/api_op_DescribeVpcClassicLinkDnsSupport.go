// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVpcClassicLinkDnsSupportInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"5" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// One or more VPC IDs.
	VpcIds []string `locationNameList:"VpcId" type:"list"`
}

// String returns the string representation
func (s DescribeVpcClassicLinkDnsSupportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcClassicLinkDnsSupportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVpcClassicLinkDnsSupportInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeVpcClassicLinkDnsSupportOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// Information about the ClassicLink DNS support status of the VPCs.
	Vpcs []ClassicLinkDnsSupport `locationName:"vpcs" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpcClassicLinkDnsSupportOutput) String() string {
	return awsutil.Prettify(s)
}