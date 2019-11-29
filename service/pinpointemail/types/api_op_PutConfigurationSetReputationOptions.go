// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to enable or disable tracking of reputation metrics for a configuration
// set.
type PutConfigurationSetReputationOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to enable or disable reputation
	// metric tracking for.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`

	// If true, tracking of reputation metrics is enabled for the configuration
	// set. If false, tracking of reputation metrics is disabled for the configuration
	// set.
	ReputationMetricsEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s PutConfigurationSetReputationOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutConfigurationSetReputationOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutConfigurationSetReputationOptionsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutConfigurationSetReputationOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationSetReputationOptionsOutput) String() string {
	return awsutil.Prettify(s)
}