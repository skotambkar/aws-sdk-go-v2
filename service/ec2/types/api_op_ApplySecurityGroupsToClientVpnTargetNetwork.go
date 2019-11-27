// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ApplySecurityGroupsToClientVpnTargetNetworkInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Client VPN endpoint.
	//
	// ClientVpnEndpointId is a required field
	ClientVpnEndpointId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The IDs of the security groups to apply to the associated target network.
	// Up to 5 security groups can be applied to an associated target network.
	//
	// SecurityGroupIds is a required field
	SecurityGroupIds []string `locationName:"SecurityGroupId" locationNameList:"item" type:"list" required:"true"`

	// The ID of the VPC in which the associated target network is located.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ApplySecurityGroupsToClientVpnTargetNetworkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ApplySecurityGroupsToClientVpnTargetNetworkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ApplySecurityGroupsToClientVpnTargetNetworkInput"}

	if s.ClientVpnEndpointId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientVpnEndpointId"))
	}

	if s.SecurityGroupIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityGroupIds"))
	}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ApplySecurityGroupsToClientVpnTargetNetworkOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the applied security groups.
	SecurityGroupIds []string `locationName:"securityGroupIds" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s ApplySecurityGroupsToClientVpnTargetNetworkOutput) String() string {
	return awsutil.Prettify(s)
}
