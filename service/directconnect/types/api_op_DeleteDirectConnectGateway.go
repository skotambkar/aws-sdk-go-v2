// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDirectConnectGatewayInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Direct Connect gateway.
	//
	// DirectConnectGatewayId is a required field
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDirectConnectGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDirectConnectGatewayInput"}

	if s.DirectConnectGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectConnectGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDirectConnectGatewayOutput struct {
	_ struct{} `type:"structure"`

	// The Direct Connect gateway.
	DirectConnectGateway *DirectConnectGateway `locationName:"directConnectGateway" type:"structure"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayOutput) String() string {
	return awsutil.Prettify(s)
}