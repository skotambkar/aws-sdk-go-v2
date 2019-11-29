// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DeregisterImage.
type DeregisterImageInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the AMI.
	//
	// ImageId is a required field
	ImageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterImageInput"}

	if s.ImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterImageOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterImageOutput) String() string {
	return awsutil.Prettify(s)
}