// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateApplicationSettingsInput struct {
	_ struct{} `type:"structure" payload:"WriteApplicationSettingsRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the default settings for an application.
	//
	// WriteApplicationSettingsRequest is a required field
	WriteApplicationSettingsRequest *WriteApplicationSettingsRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApplicationSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApplicationSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApplicationSettingsInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.WriteApplicationSettingsRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("WriteApplicationSettingsRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateApplicationSettingsOutput struct {
	_ struct{} `type:"structure" payload:"ApplicationSettingsResource"`

	// Provides information about an application, including the default settings
	// for an application.
	//
	// ApplicationSettingsResource is a required field
	ApplicationSettingsResource *ApplicationSettingsResource `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateApplicationSettingsOutput) String() string {
	return awsutil.Prettify(s)
}
