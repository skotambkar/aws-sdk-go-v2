// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteMessageInput struct {
	_ struct{} `type:"structure"`

	// The URL of the Amazon SQS queue from which messages are deleted.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`

	// The receipt handle associated with the message to delete.
	//
	// ReceiptHandle is a required field
	ReceiptHandle *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMessageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMessageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMessageInput"}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}

	if s.ReceiptHandle == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReceiptHandle"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMessageOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMessageOutput) String() string {
	return awsutil.Prettify(s)
}