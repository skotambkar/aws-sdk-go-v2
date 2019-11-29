// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateInternetGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s CreateInternetGatewayInput) String() string {
	return awsutil.Prettify(s)
}

type CreateInternetGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Information about the internet gateway.
	InternetGateway *InternetGateway `locationName:"internetGateway" type:"structure"`
}

// String returns the string representation
func (s CreateInternetGatewayOutput) String() string {
	return awsutil.Prettify(s)
}