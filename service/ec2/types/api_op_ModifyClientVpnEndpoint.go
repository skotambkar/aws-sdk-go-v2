// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyClientVpnEndpointInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint to modify.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Information about the client connection logging options.
	//
	// If you enable client connection logging, data about client connections is
	// sent to a Cloudwatch Logs log stream. The following information is logged:
	//
	//    * Client connection requests
	//
	//    * Client connection results (successful and unsuccessful)
	//
	//    * Reasons for unsuccessful client connection requests
	//
	//    * Client connection termination time
	ConnectionLogOptions *ConnectionLogOptions `type:"structure"`

	// A brief description of the Client VPN endpoint.
	Description *string `type:"string"`

	// Information about the DNS servers to be used by Client VPN connections. A
	// Client VPN endpoint can have up to two DNS servers.
	DnsServers *DnsServersOptionsModifyStructure `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ARN of the server certificate to be used. The server certificate must
	// be provisioned in AWS Certificate Manager (ACM).
	ServerCertificateArn *string `type:"string"`

	// Indicates whether the VPN is split-tunnel.
	//
	// For information about split-tunnel VPN endpoints, see Split-Tunnel AWS Client
	// VPN Endpoint (https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/split-tunnel-vpn.html)
	// in the AWS Client VPN Administrator Guide.
	SplitTunnel *bool `type:"boolean"`
}

// String returns the string representation
func (s ModifyClientVpnEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyClientVpnEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyClientVpnEndpointInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyClientVpnEndpointOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyClientVpnEndpointOutput) String() string {
	return awsutil.Prettify(s)
}
