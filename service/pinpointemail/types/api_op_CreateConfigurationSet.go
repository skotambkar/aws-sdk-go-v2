// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to create a configuration set.
type CreateConfigurationSetInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `type:"string" required:"true"`

	// An object that defines the dedicated IP pool that is used to send emails
	// that you send using the configuration set.
	DeliveryOptions *DeliveryOptions `type:"structure"`

	// An object that defines whether or not Amazon Pinpoint collects reputation
	// metrics for the emails that you send that use the configuration set.
	ReputationOptions *ReputationOptions `type:"structure"`

	// An object that defines whether or not Amazon Pinpoint can send email that
	// you send using the configuration set.
	SendingOptions *SendingOptions `type:"structure"`

	// An array of objects that define the tags (keys and values) that you want
	// to associate with the configuration set.
	Tags []Tag `type:"list"`

	// An object that defines the open and click tracking options for emails that
	// you send using the configuration set.
	TrackingOptions *TrackingOptions `type:"structure"`
}

// String returns the string representation
func (s CreateConfigurationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConfigurationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConfigurationSetInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.TrackingOptions != nil {
		if err := s.TrackingOptions.Validate(); err != nil {
			invalidParams.AddNested("TrackingOptions", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type CreateConfigurationSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateConfigurationSetOutput) String() string {
	return awsutil.Prettify(s)
}
