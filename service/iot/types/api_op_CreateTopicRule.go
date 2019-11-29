// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the CreateTopicRule operation.
type CreateTopicRuleInput struct {
	_ struct{} `type:"structure" payload:"TopicRulePayload"`

	// The name of the rule.
	//
	// RuleName is a required field
	RuleName *string `location:"uri" locationName:"ruleName" min:"1" type:"string" required:"true"`

	// Metadata which can be used to manage the topic rule.
	//
	// For URI Request parameters use format: ...key1=value1&key2=value2...
	//
	// For the CLI command-line parameter use format: --tags "key1=value1&key2=value2..."
	//
	// For the cli-input-json file use format: "tags": "key1=value1&key2=value2..."
	Tags *string `location:"header" locationName:"x-amz-tagging" type:"string"`

	// The rule payload.
	//
	// TopicRulePayload is a required field
	TopicRulePayload *TopicRulePayload `locationName:"topicRulePayload" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateTopicRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTopicRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTopicRuleInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}

	if s.TopicRulePayload == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicRulePayload"))
	}
	if s.TopicRulePayload != nil {
		if err := s.TopicRulePayload.Validate(); err != nil {
			invalidParams.AddNested("TopicRulePayload", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateTopicRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTopicRuleOutput) String() string {
	return awsutil.Prettify(s)
}