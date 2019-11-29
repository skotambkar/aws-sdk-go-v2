// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateNodeInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time.
	// This identifier is required only if you make a service request directly using
	// an HTTP client. It is generated automatically if you use an AWS SDK or the
	// AWS CLI.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The unique identifier of the member that owns this node.
	//
	// MemberId is a required field
	MemberId *string `location:"uri" locationName:"memberId" min:"1" type:"string" required:"true"`

	// The unique identifier of the network in which this node runs.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The properties of a node configuration.
	//
	// NodeConfiguration is a required field
	NodeConfiguration *NodeConfiguration `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNodeInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}
	if s.MemberId != nil && len(*s.MemberId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberId", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if s.NodeConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeConfiguration"))
	}
	if s.NodeConfiguration != nil {
		if err := s.NodeConfiguration.Validate(); err != nil {
			invalidParams.AddNested("NodeConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateNodeOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the node.
	NodeId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateNodeOutput) String() string {
	return awsutil.Prettify(s)
}