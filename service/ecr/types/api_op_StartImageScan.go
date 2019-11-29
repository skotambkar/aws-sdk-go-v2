// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartImageScanInput struct {
	_ struct{} `type:"structure"`

	// An object with identifying information for an Amazon ECR image.
	//
	// ImageId is a required field
	ImageId *ImageIdentifier `locationName:"imageId" type:"structure" required:"true"`

	// The AWS account ID associated with the registry that contains the repository
	// in which to start an image scan request. If you do not specify a registry,
	// the default registry is assumed.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The name of the repository that contains the images to scan.
	//
	// RepositoryName is a required field
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s StartImageScanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartImageScanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartImageScanInput"}

	if s.ImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageId"))
	}

	if s.RepositoryName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RepositoryName"))
	}
	if s.RepositoryName != nil && len(*s.RepositoryName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("RepositoryName", 2))
	}
	if s.ImageId != nil {
		if err := s.ImageId.Validate(); err != nil {
			invalidParams.AddNested("ImageId", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartImageScanOutput struct {
	_ struct{} `type:"structure"`

	// An object with identifying information for an Amazon ECR image.
	ImageId *ImageIdentifier `locationName:"imageId" type:"structure"`

	// The current state of the scan.
	ImageScanStatus *ImageScanStatus `locationName:"imageScanStatus" type:"structure"`

	// The registry ID associated with the request.
	RegistryId *string `locationName:"registryId" type:"string"`

	// The repository name associated with the request.
	RepositoryName *string `locationName:"repositoryName" min:"2" type:"string"`
}

// String returns the string representation
func (s StartImageScanOutput) String() string {
	return awsutil.Prettify(s)
}
