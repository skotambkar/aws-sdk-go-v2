// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

type AttachVolumeInput struct {
	_ struct{} `type:"structure"`

	// The device name (for example, /dev/sdh or xvdh).
	//
	// Device is a required field
	Device *string `type:"string" required:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the instance.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The ID of the EBS volume. The volume and instance must be within the same
	// Availability Zone.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AttachVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachVolumeInput"}

	if s.Device == nil {
		invalidParams.Add(aws.NewErrParamRequired("Device"))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes volume attachment details.
type AttachVolumeOutput struct {
	_ struct{} `type:"structure"`

	// The time stamp when the attachment initiated.
	AttachTime *time.Time `locationName:"attachTime" type:"timestamp"`

	// Indicates whether the EBS volume is deleted on instance termination.
	DeleteOnTermination *bool `locationName:"deleteOnTermination" type:"boolean"`

	// The device name.
	Device *string `locationName:"device" type:"string"`

	// The ID of the instance.
	InstanceId *string `locationName:"instanceId" type:"string"`

	// The attachment state of the volume.
	State enums.VolumeAttachmentState `locationName:"status" type:"string" enum:"true"`

	// The ID of the volume.
	VolumeId *string `locationName:"volumeId" type:"string"`
}

// String returns the string representation
func (s AttachVolumeOutput) String() string {
	return awsutil.Prettify(s)
}
