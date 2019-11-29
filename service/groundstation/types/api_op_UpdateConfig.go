// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/enums"
)

type UpdateConfigInput struct {
	_ struct{} `type:"structure"`

	// Object containing the parameters for a Config.
	//
	// See the subtype definitions for what each type of Config contains.
	//
	// ConfigData is a required field
	ConfigData *ConfigTypeData `locationName:"configData" type:"structure" required:"true"`

	// ConfigId is a required field
	ConfigId *string `location:"uri" locationName:"configId" type:"string" required:"true"`

	// ConfigType is a required field
	ConfigType enums.ConfigCapabilityType `location:"uri" locationName:"configType" type:"string" required:"true" enum:"true"`

	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigInput"}

	if s.ConfigData == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigData"))
	}

	if s.ConfigId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigId"))
	}
	if len(s.ConfigType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ConfigType"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.ConfigData != nil {
		if err := s.ConfigData.Validate(); err != nil {
			invalidParams.AddNested("ConfigData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateConfigOutput struct {
	_ struct{} `type:"structure"`

	ConfigArn *string `locationName:"configArn" type:"string"`

	ConfigId *string `locationName:"configId" type:"string"`

	ConfigType enums.ConfigCapabilityType `locationName:"configType" type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateConfigOutput) String() string {
	return awsutil.Prettify(s)
}