// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A complex type that contains information about the request to authorize associating
// a VPC with your private hosted zone. Authorization is only required when
// a private hosted zone and a VPC were created by using different accounts.
type CreateVPCAssociationAuthorizationInput struct {
	_ struct{} `locationName:"CreateVPCAssociationAuthorizationRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// The ID of the private hosted zone that you want to authorize associating
	// a VPC with.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// A complex type that contains the VPC ID and region for the VPC that you want
	// to authorize associating with your hosted zone.
	//
	// VPC is a required field
	VPC *VPC `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVPCAssociationAuthorizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVPCAssociationAuthorizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVPCAssociationAuthorizationInput"}

	if s.HostedZoneId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostedZoneId"))
	}

	if s.VPC == nil {
		invalidParams.Add(aws.NewErrParamRequired("VPC"))
	}
	if s.VPC != nil {
		if err := s.VPC.Validate(); err != nil {
			invalidParams.AddNested("VPC", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A complex type that contains the response information from a CreateVPCAssociationAuthorization
// request.
type CreateVPCAssociationAuthorizationOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the hosted zone that you authorized associating a VPC with.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `type:"string" required:"true"`

	// The VPC that you authorized associating with a hosted zone.
	//
	// VPC is a required field
	VPC *VPC `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVPCAssociationAuthorizationOutput) String() string {
	return awsutil.Prettify(s)
}