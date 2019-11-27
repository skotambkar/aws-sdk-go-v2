// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTrafficMirrorTargetInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// The description of the Traffic Mirror target.
	Description *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The network interface ID that is associated with the target.
	NetworkInterfaceId *string `type:"string"`

	// The Amazon Resource Name (ARN) of the Network Load Balancer that is associated
	// with the target.
	NetworkLoadBalancerArn *string `type:"string"`

	// The tags to assign to the Traffic Mirror target.
	TagSpecifications []TagSpecification `locationName:"TagSpecification" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s CreateTrafficMirrorTargetInput) String() string {
	return awsutil.Prettify(s)
}

type CreateTrafficMirrorTargetOutput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. For more information, see How to Ensure Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the Traffic Mirror target.
	TrafficMirrorTarget *TrafficMirrorTarget `locationName:"trafficMirrorTarget" type:"structure"`
}

// String returns the string representation
func (s CreateTrafficMirrorTargetOutput) String() string {
	return awsutil.Prettify(s)
}
