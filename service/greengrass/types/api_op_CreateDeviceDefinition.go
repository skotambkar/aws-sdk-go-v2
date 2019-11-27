// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDeviceDefinitionInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// Information about a device definition version.
	InitialVersion *DeviceDefinitionVersion `type:"structure"`

	Name *string `type:"string"`

	// The key-value pair for the resource tag.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateDeviceDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDeviceDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDeviceDefinitionInput"}
	if s.InitialVersion != nil {
		if err := s.InitialVersion.Validate(); err != nil {
			invalidParams.AddNested("InitialVersion", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDeviceDefinitionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	LastUpdatedTimestamp *string `type:"string"`

	LatestVersion *string `type:"string"`

	LatestVersionArn *string `type:"string"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s CreateDeviceDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}
