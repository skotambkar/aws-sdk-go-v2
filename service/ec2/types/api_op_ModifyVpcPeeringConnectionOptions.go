// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyVpcPeeringConnectionOptionsInput struct {
	_ struct{} `type:"structure"`

	// The VPC peering connection options for the accepter VPC.
	AccepterPeeringConnectionOptions *PeeringConnectionOptionsRequest `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The VPC peering connection options for the requester VPC.
	RequesterPeeringConnectionOptions *PeeringConnectionOptionsRequest `type:"structure"`

	// The ID of the VPC peering connection.
	//
	// VpcPeeringConnectionId is a required field
	VpcPeeringConnectionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpcPeeringConnectionOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcPeeringConnectionOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpcPeeringConnectionOptionsInput"}

	if s.VpcPeeringConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcPeeringConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpcPeeringConnectionOptionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the VPC peering connection options for the accepter VPC.
	AccepterPeeringConnectionOptions *PeeringConnectionOptions `locationName:"accepterPeeringConnectionOptions" type:"structure"`

	// Information about the VPC peering connection options for the requester VPC.
	RequesterPeeringConnectionOptions *PeeringConnectionOptions `locationName:"requesterPeeringConnectionOptions" type:"structure"`
}

// String returns the string representation
func (s ModifyVpcPeeringConnectionOptionsOutput) String() string {
	return awsutil.Prettify(s)
}