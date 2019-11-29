// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete a configuration set event destination. Configuration
// set event destinations are associated with configuration sets, which enable
// you to publish email sending events. For information about using configuration
// sets, see the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type DeleteConfigurationSetEventDestinationInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set from which to delete the event destination.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// The name of the event destination to delete.
	//
	// EventDestinationName is a required field
	EventDestinationName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationSetEventDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationSetEventDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationSetEventDestinationInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if s.EventDestinationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventDestinationName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
type DeleteConfigurationSetEventDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationSetEventDestinationOutput) String() string {
	return awsutil.Prettify(s)
}