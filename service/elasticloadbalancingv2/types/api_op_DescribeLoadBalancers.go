// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Names (ARN) of the load balancers. You can specify up
	// to 20 load balancers in a single call.
	LoadBalancerArns []string `type:"list"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `type:"string"`

	// The names of the load balancers.
	Names []string `type:"list"`

	// The maximum number of results to return with this call.
	PageSize *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s DescribeLoadBalancersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLoadBalancersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLoadBalancersInput"}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	// Information about the load balancers.
	LoadBalancers []LoadBalancer `type:"list"`

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s DescribeLoadBalancersOutput) String() string {
	return awsutil.Prettify(s)
}
