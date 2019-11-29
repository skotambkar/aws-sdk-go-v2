// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request body for CreateConfiguration.
type CreateConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The description of the configuration.
	Description *string `locationName:"description" type:"string"`

	// The versions of Apache Kafka with which you can use this MSK configuration.
	//
	// KafkaVersions is a required field
	KafkaVersions []string `locationName:"kafkaVersions" type:"list" required:"true"`

	// The name of the configuration. Configuration names are strings that match
	// the regex "^[0-9A-Za-z-]+$".
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// ServerProperties is automatically base64 encoded/decoded by the SDK.
	//
	// ServerProperties is a required field
	ServerProperties []byte `locationName:"serverProperties" type:"blob" required:"true"`
}

// String returns the string representation
func (s CreateConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConfigurationInput"}

	if s.KafkaVersions == nil {
		invalidParams.Add(aws.NewErrParamRequired("KafkaVersions"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.ServerProperties == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerProperties"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response body for CreateConfiguration
type CreateConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the configuration.
	Arn *string `locationName:"arn" type:"string"`

	// The time when the configuration was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp" timestampFormat:"iso8601"`

	// Latest revision of the configuration.
	LatestRevision *ConfigurationRevision `locationName:"latestRevision" type:"structure"`

	// The name of the configuration. Configuration names are strings that match
	// the regex "^[0-9A-Za-z-]+$".
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s CreateConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}