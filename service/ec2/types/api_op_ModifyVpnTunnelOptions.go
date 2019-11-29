// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyVpnTunnelOptionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The tunnel options to modify.
	//
	// TunnelOptions is a required field
	TunnelOptions *ModifyVpnTunnelOptionsSpecification `type:"structure" required:"true"`

	// The ID of the AWS Site-to-Site VPN connection.
	//
	// VpnConnectionId is a required field
	VpnConnectionId *string `type:"string" required:"true"`

	// The external IP address of the VPN tunnel.
	//
	// VpnTunnelOutsideIpAddress is a required field
	VpnTunnelOutsideIpAddress *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpnTunnelOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpnTunnelOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpnTunnelOptionsInput"}

	if s.TunnelOptions == nil {
		invalidParams.Add(aws.NewErrParamRequired("TunnelOptions"))
	}

	if s.VpnConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpnConnectionId"))
	}

	if s.VpnTunnelOutsideIpAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpnTunnelOutsideIpAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpnTunnelOptionsOutput struct {
	_ struct{} `type:"structure"`

	// Describes a VPN connection.
	VpnConnection *VpnConnection `locationName:"vpnConnection" type:"structure"`
}

// String returns the string representation
func (s ModifyVpnTunnelOptionsOutput) String() string {
	return awsutil.Prettify(s)
}