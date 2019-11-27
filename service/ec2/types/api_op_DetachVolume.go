// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

// Contains the parameters for DetachVolume.
type DetachVolumeInput struct {
	_ struct{} `type:"structure"`

	// The device name.
	Device *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// Forces detachment if the previous detachment attempt did not occur cleanly
	// (for example, logging into an instance, unmounting the volume, and detaching
	// normally). This option can lead to data loss or a corrupted file system.
	// Use this option only as a last resort to detach a volume from a failed instance.
	// The instance won't have an opportunity to flush file system caches or file
	// system metadata. If you use this option, you must perform file system check
	// and repair procedures.
	Force *bool `type:"boolean"`

	// The ID of the instance.
	InstanceId *string `type:"string"`

	// The ID of the volume.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachVolumeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes volume attachment details.
type DetachVolumeOutput struct {
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
func (s DetachVolumeOutput) String() string {
	return awsutil.Prettify(s)
}
