// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PeerVpcInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PeerVpcInput) String() string {
	return awsutil.Prettify(s)
}

type PeerVpcOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the request operation.
	Operation *Operation `locationName:"operation" type:"structure"`
}

// String returns the string representation
func (s PeerVpcOutput) String() string {
	return awsutil.Prettify(s)
}
