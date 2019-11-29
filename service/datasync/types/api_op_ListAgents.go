// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// ListAgentsRequest
type ListAgentsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of agents to list.
	MaxResults *int64 `type:"integer"`

	// An opaque string that indicates the position at which to begin the next list
	// of agents.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAgentsInput) String() string {
	return awsutil.Prettify(s)
}

// ListAgentsResponse
type ListAgentsOutput struct {
	_ struct{} `type:"structure"`

	// A list of agents in your account.
	Agents []AgentListEntry `type:"list"`

	// An opaque string that indicates the position at which to begin returning
	// the next list of agents.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAgentsOutput) String() string {
	return awsutil.Prettify(s)
}