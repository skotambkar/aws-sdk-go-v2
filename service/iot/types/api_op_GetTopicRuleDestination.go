// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetTopicRuleDestinationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the topic rule destination.
	//
	// Arn is a required field
	Arn *string `location:"uri" locationName:"arn" type:"string" required:"true"`
}

// String returns the string representation
func (s GetTopicRuleDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTopicRuleDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTopicRuleDestinationInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTopicRuleDestinationOutput struct {
	_ struct{} `type:"structure"`

	// The topic rule destination.
	TopicRuleDestination *TopicRuleDestination `locationName:"topicRuleDestination" type:"structure"`
}

// String returns the string representation
func (s GetTopicRuleDestinationOutput) String() string {
	return awsutil.Prettify(s)
}