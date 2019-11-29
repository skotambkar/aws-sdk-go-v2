// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Arn is a required field
	Arn *string `location:"uri" locationName:"arn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConfigurationInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response body for DescribeConfiguration.
type DescribeConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the configuration.
	Arn *string `locationName:"arn" type:"string"`

	// The time when the configuration was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp" timestampFormat:"iso8601"`

	// The description of the configuration.
	Description *string `locationName:"description" type:"string"`

	// The versions of Apache Kafka with which you can use this MSK configuration.
	KafkaVersions []string `locationName:"kafkaVersions" type:"list"`

	// Latest revision of the configuration.
	LatestRevision *ConfigurationRevision `locationName:"latestRevision" type:"structure"`

	// The name of the configuration. Configuration names are strings that match
	// the regex "^[0-9A-Za-z-]+$".
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s DescribeConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}