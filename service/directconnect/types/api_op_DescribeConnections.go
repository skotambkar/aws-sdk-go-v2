// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConnectionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection.
	ConnectionId *string `locationName:"connectionId" type:"string"`
}

// String returns the string representation
func (s DescribeConnectionsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeConnectionsOutput struct {
	_ struct{} `type:"structure"`

	// The connections.
	Connections []Connection `locationName:"connections" type:"list"`
}

// String returns the string representation
func (s DescribeConnectionsOutput) String() string {
	return awsutil.Prettify(s)
}
