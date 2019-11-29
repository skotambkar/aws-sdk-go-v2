// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateClientVpnTargetNetworkInput struct {
	_ struct{} `type:"structure"`

	// The ID of the target network association.
	//
	// AssociationId is a required field
	AssociationId *string `type:"string" required:"true"`

	// The ID of the Client VPN endpoint from which to disassociate the target network.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`
}

// String returns the string representation
func (s DisassociateClientVpnTargetNetworkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateClientVpnTargetNetworkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateClientVpnTargetNetworkInput"}

	if s.AssociationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationId"))
	}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateClientVpnTargetNetworkOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the target network association.
	AssociationId *string `locationName:"associationId" type:"string"`

	// The current state of the target network association.
	Status *AssociationStatus `locationName:"status" type:"structure"`
}

// String returns the string representation
func (s DisassociateClientVpnTargetNetworkOutput) String() string {
	return awsutil.Prettify(s)
}
