// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A complex type that contains information about the VPC that you want to disassociate
// from a specified private hosted zone.
type DisassociateVPCFromHostedZoneInput struct {
	_ struct{} `locationName:"DisassociateVPCFromHostedZoneRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// Optional: A comment about the disassociation request.
	Comment *string `type:"string"`

	// The ID of the private hosted zone that you want to disassociate a VPC from.
	//
	// HostedZoneId is a required field
	HostedZoneId *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// A complex type that contains information about the VPC that you're disassociating
	// from the specified hosted zone.
	//
	// VPC is a required field
	VPC *VPC `type:"structure" required:"true"`
}

// String returns the string representation
func (s DisassociateVPCFromHostedZoneInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateVPCFromHostedZoneInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateVPCFromHostedZoneInput"}

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

// A complex type that contains the response information for the disassociate
// request.
type DisassociateVPCFromHostedZoneOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that describes the changes made to the specified private hosted
	// zone.
	//
	// ChangeInfo is a required field
	ChangeInfo *ChangeInfo `type:"structure" required:"true"`
}

// String returns the string representation
func (s DisassociateVPCFromHostedZoneOutput) String() string {
	return awsutil.Prettify(s)
}