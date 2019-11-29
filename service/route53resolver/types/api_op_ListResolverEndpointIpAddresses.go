// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListResolverEndpointIpAddressesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of IP addresses that you want to return in the response
	// to a ListResolverEndpointIpAddresses request. If you don't specify a value
	// for MaxResults, Resolver returns up to 100 IP addresses.
	MaxResults *int64 `min:"1" type:"integer"`

	// For the first ListResolverEndpointIpAddresses request, omit this value.
	//
	// If the specified resolver endpoint has more than MaxResults IP addresses,
	// you can submit another ListResolverEndpointIpAddresses request to get the
	// next group of IP addresses. In the next request, specify the value of NextToken
	// from the previous response.
	NextToken *string `type:"string"`

	// The ID of the resolver endpoint that you want to get IP addresses for.
	//
	// ResolverEndpointId is a required field
	ResolverEndpointId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListResolverEndpointIpAddressesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResolverEndpointIpAddressesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResolverEndpointIpAddressesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ResolverEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverEndpointId"))
	}
	if s.ResolverEndpointId != nil && len(*s.ResolverEndpointId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverEndpointId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListResolverEndpointIpAddressesOutput struct {
	_ struct{} `type:"structure"`

	// The IP addresses that DNS queries pass through on their way to your network
	// (outbound endpoint) or on the way to Resolver (inbound endpoint).
	IpAddresses []IpAddressResponse `type:"list"`

	// The value that you specified for MaxResults in the request.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the specified endpoint has more than MaxResults IP addresses, you can
	// submit another ListResolverEndpointIpAddresses request to get the next group
	// of IP addresses. In the next request, specify the value of NextToken from
	// the previous response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListResolverEndpointIpAddressesOutput) String() string {
	return awsutil.Prettify(s)
}