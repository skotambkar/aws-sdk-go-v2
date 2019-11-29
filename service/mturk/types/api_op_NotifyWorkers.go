// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type NotifyWorkersInput struct {
	_ struct{} `type:"structure"`

	// The text of the email message to send. Can include up to 4,096 characters
	//
	// MessageText is a required field
	MessageText *string `type:"string" required:"true"`

	// The subject line of the email message to send. Can include up to 200 characters.
	//
	// Subject is a required field
	Subject *string `type:"string" required:"true"`

	// A list of Worker IDs you wish to notify. You can notify upto 100 Workers
	// at a time.
	//
	// WorkerIds is a required field
	WorkerIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s NotifyWorkersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NotifyWorkersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NotifyWorkersInput"}

	if s.MessageText == nil {
		invalidParams.Add(aws.NewErrParamRequired("MessageText"))
	}

	if s.Subject == nil {
		invalidParams.Add(aws.NewErrParamRequired("Subject"))
	}

	if s.WorkerIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkerIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NotifyWorkersOutput struct {
	_ struct{} `type:"structure"`

	// When MTurk sends notifications to the list of Workers, it returns back any
	// failures it encounters in this list of NotifyWorkersFailureStatus objects.
	NotifyWorkersFailureStatuses []NotifyWorkersFailureStatus `type:"list"`
}

// String returns the string representation
func (s NotifyWorkersOutput) String() string {
	return awsutil.Prettify(s)
}