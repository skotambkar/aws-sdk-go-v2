// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDiskInput struct {
	_ struct{} `type:"structure"`

	// An array of objects that represent the add-ons to enable for the new disk.
	AddOns []AddOnRequest `locationName:"addOns" type:"list"`

	// The Availability Zone where you want to create the disk (e.g., us-east-2a).
	// Use the same Availability Zone as the Lightsail instance to which you want
	// to attach the disk.
	//
	// Use the get regions operation to list the Availability Zones where Lightsail
	// is currently available.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `locationName:"availabilityZone" type:"string" required:"true"`

	// The unique Lightsail disk name (e.g., my-disk).
	//
	// DiskName is a required field
	DiskName *string `locationName:"diskName" type:"string" required:"true"`

	// The size of the disk in GB (e.g., 32).
	//
	// SizeInGb is a required field
	SizeInGb *int64 `locationName:"sizeInGb" type:"integer" required:"true"`

	// The tag keys and optional values to add to the resource during create.
	//
	// To tag a resource after it has been created, see the tag resource operation.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateDiskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDiskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDiskInput"}

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

type CreateDiskOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the API operations.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s CreateDiskOutput) String() string {
	return awsutil.Prettify(s)
}