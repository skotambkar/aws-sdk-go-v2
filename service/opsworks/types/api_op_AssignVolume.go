// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssignVolumeInput struct {
	_ struct{} `type:"structure"`

	// The instance ID.
	InstanceId *string `type:"string"`

	// The volume ID.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssignVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssignVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssignVolumeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssignVolumeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssignVolumeOutput) String() string {
	return awsutil.Prettify(s)
}