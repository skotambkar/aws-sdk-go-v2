// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteNetworkProfileInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the network profile associated with a device.
	//
	// NetworkProfileArn is a required field
	NetworkProfileArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNetworkProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNetworkProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNetworkProfileInput"}

	if s.NetworkProfileArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkProfileArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteNetworkProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNetworkProfileOutput) String() string {
	return awsutil.Prettify(s)
}