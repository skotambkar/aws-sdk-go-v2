// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateTopicRuleDestinationInput struct {
	_ struct{} `type:"structure"`

	// The topic rule destination configuration.
	//
	// DestinationConfiguration is a required field
	DestinationConfiguration *TopicRuleDestinationConfiguration `locationName:"destinationConfiguration" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateTopicRuleDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTopicRuleDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTopicRuleDestinationInput"}

	if s.DestinationConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationConfiguration"))
	}
	if s.DestinationConfiguration != nil {
		if err := s.DestinationConfiguration.Validate(); err != nil {
			invalidParams.AddNested("DestinationConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateTopicRuleDestinationOutput struct {
	_ struct{} `type:"structure"`

	// The topic rule destination.
	TopicRuleDestination *TopicRuleDestination `locationName:"topicRuleDestination" type:"structure"`
}

// String returns the string representation
func (s CreateTopicRuleDestinationOutput) String() string {
	return awsutil.Prettify(s)
}
