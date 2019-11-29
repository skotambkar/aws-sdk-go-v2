// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to enable or disable the email sending capabilities
// for a specific configuration set.
type UpdateConfigurationSetSendingEnabledInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to update.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// Describes whether email sending is enabled or disabled for the configuration
	// set.
	//
	// Enabled is a required field
	Enabled *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s UpdateConfigurationSetSendingEnabledInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigurationSetSendingEnabledInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigurationSetSendingEnabledInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if s.Enabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("Enabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateConfigurationSetSendingEnabledOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConfigurationSetSendingEnabledOutput) String() string {
	return awsutil.Prettify(s)
}