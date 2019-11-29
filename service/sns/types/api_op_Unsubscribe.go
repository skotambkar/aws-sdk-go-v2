// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for Unsubscribe action.
type UnsubscribeInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the subscription to be deleted.
	//
	// SubscriptionArn is a required field
	SubscriptionArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UnsubscribeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnsubscribeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UnsubscribeInput"}

	if s.SubscriptionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UnsubscribeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UnsubscribeOutput) String() string {
	return awsutil.Prettify(s)
}
