// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to update the tracking options for a configuration set.
type UpdateConfigurationSetTrackingOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set for which you want to update the custom
	// tracking domain.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// A domain that is used to redirect email recipients to an Amazon SES-operated
	// domain. This domain captures open and click events generated by Amazon SES
	// emails.
	//
	// For more information, see Configuring Custom Domains to Handle Open and Click
	// Tracking (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/configure-custom-open-click-domains.html)
	// in the Amazon SES Developer Guide.
	//
	// TrackingOptions is a required field
	TrackingOptions *TrackingOptions `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateConfigurationSetTrackingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConfigurationSetTrackingOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConfigurationSetTrackingOptionsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if s.TrackingOptions == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrackingOptions"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type UpdateConfigurationSetTrackingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConfigurationSetTrackingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}
