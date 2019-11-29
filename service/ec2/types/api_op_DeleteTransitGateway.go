// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTransitGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the transit gateway.
	//
	// TransitGatewayId is a required field
	TransitGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTransitGatewayInput"}

	if s.TransitGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TransitGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTransitGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Information about the deleted transit gateway.
	TransitGateway *TransitGateway `locationName:"transitGateway" type:"structure"`
}

// String returns the string representation
func (s DeleteTransitGatewayOutput) String() string {
	return awsutil.Prettify(s)
}
