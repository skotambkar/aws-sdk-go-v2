// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeregisterVolumeInput struct {
	_ struct{} `type:"structure"`

	// The AWS OpsWorks Stacks volume ID, which is the GUID that AWS OpsWorks Stacks
	// assigned to the instance when you registered the volume with the stack, not
	// the Amazon EC2 volume ID.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeregisterVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterVolumeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterVolumeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeregisterVolumeOutput) String() string {
	return awsutil.Prettify(s)
}
