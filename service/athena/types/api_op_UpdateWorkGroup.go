// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/athena/enums"
)

type UpdateWorkGroupInput struct {
	_ struct{} `type:"structure"`

	// The workgroup configuration that will be updated for the given workgroup.
	ConfigurationUpdates *WorkGroupConfigurationUpdates `type:"structure"`

	// The workgroup description.
	Description *string `type:"string"`

	// The workgroup state that will be updated for the given workgroup.
	State enums.WorkGroupState `type:"string" enum:"true"`

	// The specified workgroup that will be updated.
	//
	// WorkGroup is a required field
	WorkGroup *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateWorkGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateWorkGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateWorkGroupInput"}

	if s.WorkGroup == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkGroup"))
	}
	if s.ConfigurationUpdates != nil {
		if err := s.ConfigurationUpdates.Validate(); err != nil {
			invalidParams.AddNested("ConfigurationUpdates", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateWorkGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateWorkGroupOutput) String() string {
	return awsutil.Prettify(s)
}
