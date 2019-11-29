// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to obtain information about the event destinations for a configuration
// set.
type GetConfigurationSetEventDestinationsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that contains the event destination.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetConfigurationSetEventDestinationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetConfigurationSetEventDestinationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetConfigurationSetEventDestinationsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about an event destination for a configuration set.
type GetConfigurationSetEventDestinationsOutput struct {
	_ struct{} `type:"structure"`

	// An array that includes all of the events destinations that have been configured
	// for the configuration set.
	EventDestinations []EventDestination `type:"list"`
}

// String returns the string representation
func (s GetConfigurationSetEventDestinationsOutput) String() string {
	return awsutil.Prettify(s)
}
