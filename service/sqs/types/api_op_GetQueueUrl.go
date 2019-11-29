// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetQueueUrlInput struct {
	_ struct{} `type:"structure"`

	// The name of the queue whose URL must be fetched. Maximum 80 characters. Valid
	// values: alphanumeric characters, hyphens (-), and underscores (_).
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueName is a required field
	QueueName *string `type:"string" required:"true"`

	// The AWS account ID of the account that created the queue.
	QueueOwnerAWSAccountId *string `type:"string"`
}

// String returns the string representation
func (s GetQueueUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueueUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetQueueUrlInput"}

	if s.QueueName == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// For more information, see Interpreting Responses (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-api-responses.html)
// in the Amazon Simple Queue Service Developer Guide.
type GetQueueUrlOutput struct {
	_ struct{} `type:"structure"`

	// The URL of the queue.
	QueueUrl *string `type:"string"`
}

// String returns the string representation
func (s GetQueueUrlOutput) String() string {
	return awsutil.Prettify(s)
}
