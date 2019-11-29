// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/enums"
)

type CreateConfigInput struct {
	_ struct{} `type:"structure"`

	// Object containing the parameters for a Config.
	//
	// See the subtype definitions for what each type of Config contains.
	//
	// ConfigData is a required field
	ConfigData *ConfigTypeData `locationName:"configData" type:"structure" required:"true"`

	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConfigInput"}

	if s.ConfigData == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigData"))
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

type CreateConfigOutput struct {
	_ struct{} `type:"structure"`

	ConfigArn *string `locationName:"configArn" type:"string"`

	ConfigId *string `locationName:"configId" type:"string"`

	ConfigType enums.ConfigCapabilityType `locationName:"configType" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateConfigOutput) String() string {
	return awsutil.Prettify(s)
}