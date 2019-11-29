// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for ConfirmSubscription action.
type ConfirmSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// Disallows unauthenticated unsubscribes of the subscription. If the value
	// of this parameter is true and the request has an AWS signature, then only
	// the topic owner and the subscription owner can unsubscribe the endpoint.
	// The unsubscribe action requires AWS authentication.
	AuthenticateOnUnsubscribe *string `type:"string"`

	// Short-lived token sent to an endpoint during the Subscribe action.
	//
	// Token is a required field
	Token *string `type:"string" required:"true"`

	// The ARN of the topic for which you wish to confirm a subscription.
	//
	// TopicArn is a required field
	TopicArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ConfirmSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmSubscriptionInput"}

	if s.Token == nil {
		invalidParams.Add(aws.NewErrParamRequired("Token"))
	}

	if s.TopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for ConfirmSubscriptions action.
type ConfirmSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the created subscription.
	SubscriptionArn *string `type:"string"`
}

// String returns the string representation
func (s ConfirmSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}
