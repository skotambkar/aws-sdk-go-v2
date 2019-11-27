// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteSubnetInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the subnet.
	//
	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSubnetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSubnetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSubnetInput"}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteSubnetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteSubnetOutput) String() string {
	return awsutil.Prettify(s)
}
