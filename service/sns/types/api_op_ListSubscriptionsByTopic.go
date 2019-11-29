// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for ListSubscriptionsByTopic action.
type ListSubscriptionsByTopicInput struct {
	_ struct{} `type:"structure"`

	// Token returned by the previous ListSubscriptionsByTopic request.
	NextToken *string `type:"string"`

	// The ARN of the topic for which you wish to find subscriptions.
	//
	// TopicArn is a required field
	TopicArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListSubscriptionsByTopicInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSubscriptionsByTopicInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSubscriptionsByTopicInput"}

	if s.TopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for ListSubscriptionsByTopic action.
type ListSubscriptionsByTopicOutput struct {
	_ struct{} `type:"structure"`

	// Token to pass along to the next ListSubscriptionsByTopic request. This element
	// is returned if there are more subscriptions to retrieve.
	NextToken *string `type:"string"`

	// A list of subscriptions.
	Subscriptions []Subscription `type:"list"`
}

// String returns the string representation
func (s ListSubscriptionsByTopicOutput) String() string {
	return awsutil.Prettify(s)
}