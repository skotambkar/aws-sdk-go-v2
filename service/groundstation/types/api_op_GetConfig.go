// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/enums"
)

type GetConfigInput struct {
	_ struct{} `type:"structure"`

	// ConfigId is a required field
	ConfigId *string `location:"uri" locationName:"configId" type:"string" required:"true"`

	// ConfigType is a required field
	ConfigType enums.ConfigCapabilityType `location:"uri" locationName:"configType" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConfigInput"}

	if s.ConfigId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigId"))
	}
	if len(s.ConfigType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ConfigType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetConfigOutput struct {
	_ struct{} `type:"structure"`

	// ConfigArn is a required field
	ConfigArn *string `locationName:"configArn" type:"string" required:"true"`

	// Object containing the parameters for a Config.
	//
	// See the subtype definitions for what each type of Config contains.
	//
	// ConfigData is a required field
	ConfigData *ConfigTypeData `locationName:"configData" type:"structure" required:"true"`

	// ConfigId is a required field
	ConfigId *string `locationName:"configId" type:"string" required:"true"`

	ConfigType enums.ConfigCapabilityType `locationName:"configType" type:"string" enum:"true"`

	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetConfigOutput) String() string {
	return awsutil.Prettify(s)
}