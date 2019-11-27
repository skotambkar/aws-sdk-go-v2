// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/enums"
)

type CreateResolverEndpointInput struct {
	_ struct{} `type:"structure"`

	// A unique string that identifies the request and that allows failed requests
	// to be retried without the risk of executing the operation twice. CreatorRequestId
	// can be any unique string, for example, a date/time stamp.
	//
	// CreatorRequestId is a required field
	CreatorRequestId *string `min:"1" type:"string" required:"true"`

	// Specify the applicable value:
	//
	//    * INBOUND: Resolver forwards DNS queries to the DNS service for a VPC
	//    from your network or another VPC
	//
	//    * OUTBOUND: Resolver forwards DNS queries from the DNS service for a VPC
	//    to your network or another VPC
	//
	// Direction is a required field
	Direction enums.ResolverEndpointDirection `type:"string" required:"true" enum:"true"`

	// The subnets and IP addresses in your VPC that you want DNS queries to pass
	// through on the way from your VPCs to your network (for outbound endpoints)
	// or on the way from your network to your VPCs (for inbound resolver endpoints).
	//
	// IpAddresses is a required field
	IpAddresses []IpAddressRequest `min:"1" type:"list" required:"true"`

	// A friendly name that lets you easily find a configuration in the Resolver
	// dashboard in the Route 53 console.
	Name *string `type:"string"`

	// The ID of one or more security groups that you want to use to control access
	// to this VPC. The security group that you specify must include one or more
	// inbound rules (for inbound resolver endpoints) or outbound rules (for outbound
	// resolver endpoints).
	//
	// SecurityGroupIds is a required field
	SecurityGroupIds []string `type:"list" required:"true"`

	// A list of the tag keys and values that you want to associate with the endpoint.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateResolverEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResolverEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateResolverEndpointInput"}

	if s.CreatorRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CreatorRequestId"))
	}
	if s.CreatorRequestId != nil && len(*s.CreatorRequestId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CreatorRequestId", 1))
	}
	if len(s.Direction) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Direction"))
	}

	if s.IpAddresses == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpAddresses"))
	}
	if s.IpAddresses != nil && len(s.IpAddresses) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IpAddresses", 1))
	}

	if s.SecurityGroupIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityGroupIds"))
	}
	if s.IpAddresses != nil {
		for i, v := range s.IpAddresses {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "IpAddresses", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateResolverEndpointOutput struct {
	_ struct{} `type:"structure"`

	// Information about the CreateResolverEndpoint request, including the status
	// of the request.
	ResolverEndpoint *ResolverEndpoint `type:"structure"`
}

// String returns the string representation
func (s CreateResolverEndpointOutput) String() string {
	return awsutil.Prettify(s)
}
