// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePlacementInput struct {
	_ struct{} `type:"structure"`

	// Optional user-defined key/value pairs providing contextual data (such as
	// location or function) for the placement.
	Attributes map[string]string `locationName:"attributes" type:"map"`

	// The name of the placement to be created.
	//
	// PlacementName is a required field
	PlacementName *string `locationName:"placementName" min:"1" type:"string" required:"true"`

	// The name of the project in which to create the placement.
	//
	// ProjectName is a required field
	ProjectName *string `location:"uri" locationName:"projectName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePlacementInput"}

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

type CreatePlacementOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreatePlacementOutput) String() string {
	return awsutil.Prettify(s)
}