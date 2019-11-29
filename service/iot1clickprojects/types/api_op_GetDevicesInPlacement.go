// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDevicesInPlacementInput struct {
	_ struct{} `type:"structure"`

	// The name of the placement to get the devices from.
	//
	// PlacementName is a required field
	PlacementName *string `location:"uri" locationName:"placementName" min:"1" type:"string" required:"true"`

	// The name of the project containing the placement.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDevicesInPlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDevicesInPlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDevicesInPlacementInput"}

	if s.PlacementName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlacementName"))
	}
	if s.PlacementName != nil && len(*s.PlacementName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementName", 1))
	}

	if s.ProjectName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectName"))
	}
	if s.ProjectName != nil && len(*s.ProjectName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDevicesInPlacementOutput struct {
	_ struct{} `type:"structure"`

	// An object containing the devices (zero or more) within the placement.
	//
	// Devices is a required field
	Devices map[string]string `locationName:"devices" type:"map" required:"true"`
}

// String returns the string representation
func (s GetDevicesInPlacementOutput) String() string {
	return awsutil.Prettify(s)
}
