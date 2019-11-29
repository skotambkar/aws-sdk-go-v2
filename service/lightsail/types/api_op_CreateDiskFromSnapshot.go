// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDiskFromSnapshotInput struct {
	_ struct{} `type:"structure"`

	// An array of objects that represent the add-ons to enable for the new disk.
	AddOns []AddOnRequest `locationName:"addOns" type:"list"`

	// The Availability Zone where you want to create the disk (e.g., us-east-2a).
	// Choose the same Availability Zone as the Lightsail instance where you want
	// to create the disk.
	//
	// Use the GetRegions operation to list the Availability Zones where Lightsail
	// is currently available.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `locationName:"availabilityZone" type:"string" required:"true"`

	// The unique Lightsail disk name (e.g., my-disk).
	//
	// DiskName is a required field
	DiskName *string `locationName:"diskName" type:"string" required:"true"`

	// The name of the disk snapshot (e.g., my-snapshot) from which to create the
	// new storage disk.
	//
	// This parameter cannot be defined together with the source disk name parameter.
	// The disk snapshot name and source disk name parameters are mutually exclusive.
	DiskSnapshotName *string `locationName:"diskSnapshotName" type:"string"`

	// The date of the automatic snapshot to use for the new disk.
	//
	// Use the get auto snapshots operation to identify the dates of the available
	// automatic snapshots.
	//
	// Constraints:
	//
	//    * Must be specified in YYYY-MM-DD format.
	//
	//    * This parameter cannot be defined together with the use latest restorable
	//    auto snapshot parameter. The restore date and use latest restorable auto
	//    snapshot parameters are mutually exclusive.
	//
	// Define this parameter only when creating a new disk from an automatic snapshot.
	// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	RestoreDate *string `locationName:"restoreDate" type:"string"`

	// The size of the disk in GB (e.g., 32).
	//
	// SizeInGb is a required field
	SizeInGb *int64 `locationName:"sizeInGb" type:"integer" required:"true"`

	// The name of the source disk from which the source automatic snapshot was
	// created.
	//
	// This parameter cannot be defined together with the disk snapshot name parameter.
	// The source disk name and disk snapshot name parameters are mutually exclusive.
	//
	// Define this parameter only when creating a new disk from an automatic snapshot.
	// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	SourceDiskName *string `locationName:"sourceDiskName" type:"string"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`

	// A Boolean value to indicate whether to use the latest available automatic
	// snapshot.
	//
	// This parameter cannot be defined together with the restore date parameter.
	// The use latest restorable auto snapshot and restore date parameters are mutually
	// exclusive.
	//
	// Define this parameter only when creating a new disk from an automatic snapshot.
	// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	UseLatestRestorableAutoSnapshot *bool `locationName:"useLatestRestorableAutoSnapshot" type:"boolean"`
}

// String returns the string representation
func (s CreateDiskFromSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDiskFromSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDiskFromSnapshotInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}

	if s.DiskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskName"))
	}

	if s.SizeInGb == nil {
		invalidParams.Add(aws.NewErrParamRequired("SizeInGb"))
	}
	if s.AddOns != nil {
		for i, v := range s.AddOns {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AddOns", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDiskFromSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the API operations.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s CreateDiskFromSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}
