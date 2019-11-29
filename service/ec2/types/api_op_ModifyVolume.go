// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

type ModifyVolumeInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The target IOPS rate of the volume.
	//
	// This is only valid for Provisioned IOPS SSD (io1) volumes. For more information,
	// see Provisioned IOPS SSD (io1) Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html#EBSVolumeTypes_piops).
	//
	// Default: If no IOPS value is specified, the existing value is retained.
	Iops *int64 `type:"integer"`

	// The target size of the volume, in GiB. The target volume size must be greater
	// than or equal to than the existing size of the volume. For information about
	// available EBS volume sizes, see Amazon EBS Volume Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html).
	//
	// Default: If no size is specified, the existing size is retained.
	Size *int64 `type:"integer"`

	// The ID of the volume.
	//
	// VolumeId is a required field
	VolumeId *string `type:"string" required:"true"`

	// The target EBS volume type of the volume.
	//
	// Default: If no type is specified, the existing type is retained.
	VolumeType enums.VolumeType `type:"string" enum:"true"`
}

// String returns the string representation
func (s ModifyVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVolumeInput"}

	if s.VolumeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVolumeOutput struct {
	_ struct{} `type:"structure"`

	// Information about the volume modification.
	VolumeModification *VolumeModification `locationName:"volumeModification" type:"structure"`
}

// String returns the string representation
func (s ModifyVolumeOutput) String() string {
	return awsutil.Prettify(s)
}