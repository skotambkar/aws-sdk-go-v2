// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateMaintenanceWindowTargetInput struct {
	_ struct{} `type:"structure"`

	// An optional description for the update.
	Description *string `min:"1" type:"string" sensitive:"true"`

	// A name for the update.
	Name *string `min:"3" type:"string"`

	// User-provided value that will be included in any CloudWatch events raised
	// while running tasks for these targets in this maintenance window.
	OwnerInformation *string `min:"1" type:"string" sensitive:"true"`

	// If True, then all fields that are required by the RegisterTargetWithMaintenanceWindow
	// action are also required for this API request. Optional fields that are not
	// specified are set to null.
	Replace *bool `type:"boolean"`

	// The targets to add or replace.
	Targets []Target `type:"list"`

	// The maintenance window ID with which to modify the target.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`

	// The target ID to modify.
	//
	// WindowTargetId is a required field
	WindowTargetId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateMaintenanceWindowTargetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMaintenanceWindowTargetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMaintenanceWindowTargetInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.Name != nil && len(*s.Name) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 3))
	}
	if s.OwnerInformation != nil && len(*s.OwnerInformation) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OwnerInformation", 1))
	}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}

	if s.WindowTargetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowTargetId"))
	}
	if s.WindowTargetId != nil && len(*s.WindowTargetId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowTargetId", 36))
	}
	if s.Targets != nil {
		for i, v := range s.Targets {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Targets", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateMaintenanceWindowTargetOutput struct {
	_ struct{} `type:"structure"`

	// The updated description.
	Description *string `min:"1" type:"string" sensitive:"true"`

	// The updated name.
	Name *string `min:"3" type:"string"`

	// The updated owner.
	OwnerInformation *string `min:"1" type:"string" sensitive:"true"`

	// The updated targets.
	Targets []Target `type:"list"`

	// The maintenance window ID specified in the update request.
	WindowId *string `min:"20" type:"string"`

	// The target ID specified in the update request.
	WindowTargetId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s UpdateMaintenanceWindowTargetOutput) String() string {
	return awsutil.Prettify(s)
}
