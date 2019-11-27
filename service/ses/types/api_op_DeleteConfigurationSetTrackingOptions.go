// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete open and click tracking options in a configuration
// set.
type DeleteConfigurationSetTrackingOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set from which you want to delete the tracking
	// options.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationSetTrackingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationSetTrackingOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationSetTrackingOptionsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type DeleteConfigurationSetTrackingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationSetTrackingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}
