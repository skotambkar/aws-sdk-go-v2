// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/datasync/enums"
)

// DescribeAgent
type DescribeAgentInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent to describe.
	//
	// AgentArn is a required field
	AgentArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAgentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAgentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAgentInput"}

	if s.AgentArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AgentArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeAgentResponse
type DescribeAgentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent.
	AgentArn *string `type:"string"`

	// The time that the agent was activated (that is, created in your account).
	CreationTime *time.Time `type:"timestamp"`

	// The type of endpoint that your agent is connected to. If the endpoint is
	// a VPC endpoint, the agent is not accessible over the public Internet.
	EndpointType enums.EndpointType `type:"string" enum:"true"`

	// The time that the agent last connected to DataSyc.
	LastConnectionTime *time.Time `type:"timestamp"`

	// The name of the agent.
	Name *string `min:"1" type:"string"`

	// The subnet and the security group that DataSync used to access a VPC endpoint.
	PrivateLinkConfig *PrivateLinkConfig `type:"structure"`

	// The status of the agent. If the status is ONLINE, then the agent is configured
	// properly and is available to use. The Running status is the normal running
	// status for an agent. If the status is OFFLINE, the agent's VM is turned off
	// or the agent is in an unhealthy state. When the issue that caused the unhealthy
	// state is resolved, the agent returns to ONLINE status.
	Status enums.AgentStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeAgentOutput) String() string {
	return awsutil.Prettify(s)
}
