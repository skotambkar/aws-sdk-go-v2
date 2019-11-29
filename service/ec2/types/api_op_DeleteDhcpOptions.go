// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDhcpOptionsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the DHCP options set.
	//
	// DhcpOptionsId is a required field
	DhcpOptionsId *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s DeleteDhcpOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDhcpOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDhcpOptionsInput"}

	if s.DhcpOptionsId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DhcpOptionsId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDhcpOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDhcpOptionsOutput) String() string {
	return awsutil.Prettify(s)
}
