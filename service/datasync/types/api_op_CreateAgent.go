// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// CreateAgentRequest
type CreateAgentInput struct {
	_ struct{} `type:"structure"`

	// Your agent activation key. You can get the activation key either by sending
	// an HTTP GET request with redirects that enable you to get the agent IP address
	// (port 80). Alternatively, you can get it from the AWS DataSync console.
	//
	// The redirect URL returned in the response provides you the activation key
	// for your agent in the query string parameter activationKey. It might also
	// include other activation-related parameters; however, these are merely defaults.
	// The arguments you pass to this API call determine the actual configuration
	// of your agent.
	//
	// For more information, see Activating an Agent in the AWS DataSync User Guide.
	//
	// ActivationKey is a required field
	ActivationKey *string `type:"string" required:"true"`

	// The name you configured for your agent. This value is a text reference that
	// is used to identify the agent in the console.
	AgentName *string `min:"1" type:"string"`

	// The ARNs of the security groups used to protect your data transfer task subnets.
	// See CreateAgentRequest$SubnetArns.
	SecurityGroupArns []string `min:"1" type:"list"`

	// The Amazon Resource Names (ARNs) of the subnets in which DataSync will create
	// elastic network interfaces for each data transfer task. The agent that runs
	// a task must be private. When you start a task that is associated with an
	// agent created in a VPC, or one that has access to an IP address in a VPC,
	// then the task is also private. In this case, DataSync creates four network
	// interfaces for each task in your subnet. For a data transfer to work, the
	// agent must be able to route to all these four network interfaces.
	SubnetArns []string `min:"1" type:"list"`

	// The key-value pair that represents the tag that you want to associate with
	// the agent. The value can be an empty string. This value helps you manage,
	// filter, and search for your agents.
	//
	// Valid characters for key and value are letters, spaces, and numbers representable
	// in UTF-8 format, and the following special characters: + - = . _ : / @.
	Tags []TagListEntry `type:"list"`

	// The ID of the VPC (Virtual Private Cloud) endpoint that the agent has access
	// to. This is the client-side VPC endpoint, also called a PrivateLink. If you
	// don't have a PrivateLink VPC endpoint, see Creating a VPC Endpoint Service
	// Configuration (https://docs.aws.amazon.com/vpc/latest/userguide/endpoint-service.html#create-endpoint-service)
	// in the AWS VPC User Guide.
	//
	// VPC endpoint ID looks like this: vpce-01234d5aff67890e1.
	VpcEndpointId *string `type:"string"`
}

// String returns the string representation
func (s CreateAgentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAgentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAgentInput"}

	if s.ActivationKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivationKey"))
	}
	if s.AgentName != nil && len(*s.AgentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AgentName", 1))
	}
	if s.SecurityGroupArns != nil && len(s.SecurityGroupArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityGroupArns", 1))
	}
	if s.SubnetArns != nil && len(s.SubnetArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SubnetArns", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// CreateAgentResponse
type CreateAgentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent. Use the ListAgents operation
	// to return a list of agents for your account and AWS Region.
	AgentArn *string `type:"string"`
}

// String returns the string representation
func (s CreateAgentOutput) String() string {
	return awsutil.Prettify(s)
}
