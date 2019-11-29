// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteVpcPeeringConnectionInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC peering connection.
	//
	// VpcPeeringConnectionId is a required field
	VpcPeeringConnectionId *string `locationName:"vpcPeeringConnectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpcPeeringConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcPeeringConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVpcPeeringConnectionInput"}

	if s.VpcPeeringConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcPeeringConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteVpcPeeringConnectionOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteVpcPeeringConnectionOutput) String() string {
	return awsutil.Prettify(s)
}
