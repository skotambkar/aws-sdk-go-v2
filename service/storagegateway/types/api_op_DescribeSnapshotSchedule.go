// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the DescribeSnapshotScheduleInput$VolumeARN of the
// volume.
type DescribeSnapshotScheduleInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the volume. Use the ListVolumes operation
	// to return a list of gateway volumes.
	//
	// VolumeARN is a required field
	VolumeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSnapshotScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSnapshotScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSnapshotScheduleInput"}

	if s.VolumeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeARN"))
	}
	if s.VolumeARN != nil && len(*s.VolumeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("VolumeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeSnapshotScheduleOutput struct {
	_ struct{} `type:"structure"`

	// The snapshot description.
	Description *string `min:"1" type:"string"`

	// The number of hours between snapshots.
	RecurrenceInHours *int64 `min:"1" type:"integer"`

	// The hour of the day at which the snapshot schedule begins represented as
	// hh, where hh is the hour (0 to 23). The hour of the day is in the time zone
	// of the gateway.
	StartAt *int64 `type:"integer"`

	// A list of up to 50 tags assigned to the snapshot schedule, sorted alphabetically
	// by key name. Each tag is a key-value pair. For a gateway with more than 10
	// tags assigned, you can view all tags using the ListTagsForResource API operation.
	Tags []Tag `type:"list"`

	// A value that indicates the time zone of the gateway.
	Timezone *string `min:"3" type:"string"`

	// The Amazon Resource Name (ARN) of the volume that was specified in the request.
	VolumeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DescribeSnapshotScheduleOutput) String() string {
	return awsutil.Prettify(s)
}