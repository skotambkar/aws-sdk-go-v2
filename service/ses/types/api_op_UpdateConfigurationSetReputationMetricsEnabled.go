// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to modify the reputation metric publishing settings
// for a configuration set.
type UpdateConfigurationSetReputationMetricsEnabledInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to update.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// Describes whether or not Amazon SES will publish reputation metrics for the
	// configuration set, such as bounce and complaint rates, to Amazon CloudWatch.
	//
	// Enabled is a required field
	Enabled *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s UpdateConfigurationSetReputationMetricsEnabledInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigurationSetReputationMetricsEnabledInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigurationSetReputationMetricsEnabledInput"}

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

type UpdateConfigurationSetReputationMetricsEnabledOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConfigurationSetReputationMetricsEnabledOutput) String() string {
	return awsutil.Prettify(s)
}
