// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/enums"
)

type UpdateAuthEventFeedbackInput struct {
	_ struct{} `type:"structure"`

	// The event ID.
	//
	// EventId is a required field
	EventId *string `min:"1" type:"string" required:"true"`

	// The feedback token.
	//
	// FeedbackToken is a required field
	FeedbackToken *string `type:"string" required:"true" sensitive:"true"`

	// The authentication event feedback value.
	//
	// FeedbackValue is a required field
	FeedbackValue enums.FeedbackValueType `type:"string" required:"true" enum:"true"`

	// The user pool ID.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`

	// The user pool username.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s UpdateAuthEventFeedbackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAuthEventFeedbackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAuthEventFeedbackInput"}

	if s.EventId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventId"))
	}
	if s.EventId != nil && len(*s.EventId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventId", 1))
	}

	if s.FeedbackToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("FeedbackToken"))
	}
	if len(s.FeedbackValue) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("FeedbackValue"))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAuthEventFeedbackOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAuthEventFeedbackOutput) String() string {
	return awsutil.Prettify(s)
}
