// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to add a custom domain for tracking open and click events to a
// configuration set.
type PutConfigurationSetTrackingOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to add a custom tracking
	// domain to.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`

	// The domain that you want to use to track open and click events.
	CustomRedirectDomain *string `type:"string"`
}

// String returns the string representation
func (s PutConfigurationSetTrackingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutConfigurationSetTrackingOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutConfigurationSetTrackingOptionsInput"}

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
type PutConfigurationSetTrackingOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationSetTrackingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}