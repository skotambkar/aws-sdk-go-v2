// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for ListSubscriptions action.
type ListSubscriptionsInput struct {
	_ struct{} `type:"structure"`

	// Token returned by the previous ListSubscriptions request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListSubscriptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Response for ListSubscriptions action
type ListSubscriptionsOutput struct {
	_ struct{} `type:"structure"`

	// Token to pass along to the next ListSubscriptions request. This element is
	// returned if there are more subscriptions to retrieve.
	NextToken *string `type:"string"`

	// A list of subscriptions.
	Subscriptions []Subscription `type:"list"`
}

// String returns the string representation
func (s ListSubscriptionsOutput) String() string {
	return awsutil.Prettify(s)
}
