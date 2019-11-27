// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DeleteAgentRequest
type DeleteAgentInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent to delete. Use the ListAgents
	// operation to return a list of agents for your account and AWS Region.
	//
	// AgentArn is a required field
	AgentArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAgentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAgentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAgentInput"}

	if s.AgentArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AgentArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAgentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAgentOutput) String() string {
	return awsutil.Prettify(s)
}
