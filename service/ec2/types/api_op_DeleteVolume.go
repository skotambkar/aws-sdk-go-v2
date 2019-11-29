// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteVolumeInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the volume.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVolumeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteVolumeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteVolumeOutput) String() string {
	return awsutil.Prettify(s)
}
