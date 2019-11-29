// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/enums"
)

// The updates that you want to make to an existing output of an existing flow.
type UpdateFlowOutputInput struct {
	_ struct{} `type:"structure"`

	// The range of IP addresses that should be allowed to initiate output requests
	// to this flow. These IP addresses should be in the form of a Classless Inter-Domain
	// Routing (CIDR) block; for example, 10.0.0.0/16.
	CidrAllowList []string `locationName:"cidrAllowList" type:"list"`

	// A description of the output. This description appears only on the AWS Elemental
	// MediaConnect console and will not be seen by the end user.
	Description *string `locationName:"description" type:"string"`

	// The IP address where you want to send the output.
	Destination *string `locationName:"destination" type:"string"`

	// The type of key used for the encryption. If no keyType is provided, the service
	// will use the default setting (static-key).
	Encryption *UpdateEncryption `locationName:"encryption" type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// The maximum latency in milliseconds for Zixi-based streams.
	MaxLatency *int64 `locationName:"maxLatency" type:"integer"`

	// OutputArn is a required field
	OutputArn *string `location:"uri" locationName:"outputArn" type:"string" required:"true"`

	// The port to use when content is distributed to this output.
	Port *int64 `locationName:"port" type:"integer"`

	// The protocol to use for the output.
	Protocol enums.Protocol `locationName:"protocol" type:"string" enum:"true"`

	// The remote ID for the Zixi-pull stream.
	RemoteId *string `locationName:"remoteId" type:"string"`

	// The smoothing latency in milliseconds for RIST, RTP, and RTP-FEC streams.
	SmoothingLatency *int64 `locationName:"smoothingLatency" type:"integer"`

	// The stream ID that you want to use for this transport. This parameter applies
	// only to Zixi-based streams.
	StreamId *string `locationName:"streamId" type:"string"`
}

// String returns the string representation
func (s UpdateFlowOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFlowOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFlowOutputInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if s.OutputArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a successful UpdateFlowOutput request including the flow ARN
// and the updated output.
type UpdateFlowOutputOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that is associated with the updated output.
	FlowArn *string `locationName:"flowArn" type:"string"`

	// The settings for an output.
	Output *Output `locationName:"output" type:"structure"`
}

// String returns the string representation
func (s UpdateFlowOutputOutput) String() string {
	return awsutil.Prettify(s)
}
