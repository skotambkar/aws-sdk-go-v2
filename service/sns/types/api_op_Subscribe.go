// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for Subscribe action.
type SubscribeInput struct {
	_ struct{} `type:"structure"`

	// A map of attributes with their corresponding values.
	//
	// The following lists the names, descriptions, and values of the special request
	// parameters that the SetTopicAttributes action uses:
	//
	//    * DeliveryPolicy – The policy that defines how Amazon SNS retries failed
	//    deliveries to HTTP/S endpoints.
	//
	//    * FilterPolicy – The simple JSON object that lets your subscriber receive
	//    only a subset of messages, rather than receiving every message published
	//    to the topic.
	//
	//    * RawMessageDelivery – When set to true, enables raw message delivery
	//    to Amazon SQS or HTTP/S endpoints. This eliminates the need for the endpoints
	//    to process JSON formatting, which is otherwise created for Amazon SNS
	//    metadata.
	Attributes map[string]string `type:"map"`

	// The endpoint that you want to receive notifications. Endpoints vary by protocol:
	//
	//    * For the http protocol, the endpoint is an URL beginning with "https://"
	//
	//    * For the https protocol, the endpoint is a URL beginning with "https://"
	//
	//    * For the email protocol, the endpoint is an email address
	//
	//    * For the email-json protocol, the endpoint is an email address
	//
	//    * For the sms protocol, the endpoint is a phone number of an SMS-enabled
	//    device
	//
	//    * For the sqs protocol, the endpoint is the ARN of an Amazon SQS queue
	//
	//    * For the application protocol, the endpoint is the EndpointArn of a mobile
	//    app and device.
	//
	//    * For the lambda protocol, the endpoint is the ARN of an AWS Lambda function.
	Endpoint *string `type:"string"`

	// The protocol you want to use. Supported protocols include:
	//
	//    * http – delivery of JSON-encoded message via HTTP POST
	//
	//    * https – delivery of JSON-encoded message via HTTPS POST
	//
	//    * email – delivery of message via SMTP
	//
	//    * email-json – delivery of JSON-encoded message via SMTP
	//
	//    * sms – delivery of message via SMS
	//
	//    * sqs – delivery of JSON-encoded message to an Amazon SQS queue
	//
	//    * application – delivery of JSON-encoded message to an EndpointArn for
	//    a mobile app and device.
	//
	//    * lambda – delivery of JSON-encoded message to an AWS Lambda function.
	//
	// Protocol is a required field
	Protocol *string `type:"string" required:"true"`

	// Sets whether the response from the Subscribe request includes the subscription
	// ARN, even if the subscription is not yet confirmed.
	//
	// If you set this parameter to false, the response includes the ARN for confirmed
	// subscriptions, but it includes an ARN value of "pending subscription" for
	// subscriptions that are not yet confirmed. A subscription becomes confirmed
	// when the subscriber calls the ConfirmSubscription action with a confirmation
	// token.
	//
	// If you set this parameter to true, the response includes the ARN in all cases,
	// even if the subscription is not yet confirmed.
	//
	// The default value is false.
	ReturnSubscriptionArn *bool `type:"boolean"`

	// The ARN of the topic you want to subscribe to.
	//
	// TopicArn is a required field
	TopicArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SubscribeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubscribeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SubscribeInput"}

	if s.Protocol == nil {
		invalidParams.Add(aws.NewErrParamRequired("Protocol"))
	}

	if s.TopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for Subscribe action.
type SubscribeOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the subscription if it is confirmed, or the string "pending confirmation"
	// if the subscription requires confirmation. However, if the API request parameter
	// ReturnSubscriptionArn is true, then the value is always the subscription
	// ARN, even if the subscription requires confirmation.
	SubscriptionArn *string `type:"string"`
}

// String returns the string representation
func (s SubscribeOutput) String() string {
	return awsutil.Prettify(s)
}