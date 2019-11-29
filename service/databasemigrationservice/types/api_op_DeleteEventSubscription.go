// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteEventSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// The name of the DMS event notification subscription to be deleted.
	//
	// SubscriptionName is a required field
	SubscriptionName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEventSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEventSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEventSubscriptionInput"}

	if s.SubscriptionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteEventSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// The event subscription that was deleted.
	EventSubscription *EventSubscription `type:"structure"`
}

// String returns the string representation
func (s DeleteEventSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}