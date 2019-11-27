// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CopyImageInput struct {
	_ struct{} `type:"structure"`

	// The description that the image will have when it is copied to the destination.
	DestinationImageDescription *string `type:"string"`

	// The name that the image will have when it is copied to the destination.
	//
	// DestinationImageName is a required field
	DestinationImageName *string `type:"string" required:"true"`

	// The destination region to which the image will be copied. This parameter
	// is required, even if you are copying an image within the same region.
	//
	// DestinationRegion is a required field
	DestinationRegion *string `min:"1" type:"string" required:"true"`

	// The name of the image to copy.
	//
	// SourceImageName is a required field
	SourceImageName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CopyImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CopyImageInput"}

	if s.DestinationImageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationImageName"))
	}

	if s.DestinationRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationRegion"))
	}
	if s.DestinationRegion != nil && len(*s.DestinationRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DestinationRegion", 1))
	}

	if s.SourceImageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceImageName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CopyImageOutput struct {
	_ struct{} `type:"structure"`

	// The name of the destination image.
	DestinationImageName *string `type:"string"`
}

// String returns the string representation
func (s CopyImageOutput) String() string {
	return awsutil.Prettify(s)
}
