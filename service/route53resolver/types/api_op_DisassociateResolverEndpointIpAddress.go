// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateResolverEndpointIpAddressInput struct {
	_ struct{} `type:"structure"`

	// The IPv4 address that you want to remove from a resolver endpoint.
	//
	// IpAddress is a required field
	IpAddress *IpAddressUpdate `type:"structure" required:"true"`

	// The ID of the resolver endpoint that you want to disassociate an IP address
	// from.
	//
	// ResolverEndpointId is a required field
	ResolverEndpointId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateResolverEndpointIpAddressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateResolverEndpointIpAddressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateResolverEndpointIpAddressInput"}

	if s.IpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpAddress"))
	}

	if s.ResolverEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResolverEndpointId"))
	}
	if s.ResolverEndpointId != nil && len(*s.ResolverEndpointId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResolverEndpointId", 1))
	}
	if s.IpAddress != nil {
		if err := s.IpAddress.Validate(); err != nil {
			invalidParams.AddNested("IpAddress", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateResolverEndpointIpAddressOutput struct {
	_ struct{} `type:"structure"`

	// The response to an DisassociateResolverEndpointIpAddress request.
	ResolverEndpoint *ResolverEndpoint `type:"structure"`
}

// String returns the string representation
func (s DisassociateResolverEndpointIpAddressOutput) String() string {
	return awsutil.Prettify(s)
}
