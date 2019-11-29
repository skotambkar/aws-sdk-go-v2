// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/sqs/enums"
)

type GetQueueAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list of attributes for which to retrieve information.
	//
	// In the future, new attributes might be added. If you write code that calls
	// this action, we recommend that you structure your code so that it can handle
	// new attributes gracefully.
	//
	// The following attributes are supported:
	//
	//    * All - Returns all values.
	//
	//    * ApproximateNumberOfMessages - Returns the approximate number of messages
	//    available for retrieval from the queue.
	//
	//    * ApproximateNumberOfMessagesDelayed - Returns the approximate number
	//    of messages in the queue that are delayed and not available for reading
	//    immediately. This can happen when the queue is configured as a delay queue
	//    or when a message has been sent with a delay parameter.
	//
	//    * ApproximateNumberOfMessagesNotVisible - Returns the approximate number
	//    of messages that are in flight. Messages are considered to be in flight
	//    if they have been sent to a client but have not yet been deleted or have
	//    not yet reached the end of their visibility window.
	//
	//    * CreatedTimestamp - Returns the time when the queue was created in seconds
	//    (epoch time (http://en.wikipedia.org/wiki/Unix_time)).
	//
	//    * DelaySeconds - Returns the default delay on the queue in seconds.
	//
	//    * LastModifiedTimestamp - Returns the time when the queue was last changed
	//    in seconds (epoch time (http://en.wikipedia.org/wiki/Unix_time)).
	//
	//    * MaximumMessageSize - Returns the limit of how many bytes a message can
	//    contain before Amazon SQS rejects it.
	//
	//    * MessageRetentionPeriod - Returns the length of time, in seconds, for
	//    which Amazon SQS retains a message.
	//
	//    * Policy - Returns the policy of the queue.
	//
	//    * QueueArn - Returns the Amazon resource name (ARN) of the queue.
	//
	//    * ReceiveMessageWaitTimeSeconds - Returns the length of time, in seconds,
	//    for which the ReceiveMessage action waits for a message to arrive.
	//
	//    * RedrivePolicy - Returns the string that includes the parameters for
	//    dead-letter queue functionality of the source queue. For more information
	//    about the redrive policy and dead-letter queues, see Using Amazon SQS
	//    Dead-Letter Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
	//    in the Amazon Simple Queue Service Developer Guide. deadLetterTargetArn
	//    - The Amazon Resource Name (ARN) of the dead-letter queue to which Amazon
	//    SQS moves messages after the value of maxReceiveCount is exceeded. maxReceiveCount
	//    - The number of times a message is delivered to the source queue before
	//    being moved to the dead-letter queue. When the ReceiveCount for a message
	//    exceeds the maxReceiveCount for a queue, Amazon SQS moves the message
	//    to the dead-letter-queue.
	//
	//    * VisibilityTimeout - Returns the visibility timeout for the queue. For
	//    more information about the visibility timeout, see Visibility Timeout
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	//    in the Amazon Simple Queue Service Developer Guide.
	//
	// The following attributes apply only to server-side-encryption (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html):
	//
	//    * KmsMasterKeyId - Returns the ID of an AWS-managed customer master key
	//    (CMK) for Amazon SQS or a custom CMK. For more information, see Key Terms
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	//
	//    * KmsDataKeyReusePeriodSeconds - Returns the length of time, in seconds,
	//    for which Amazon SQS can reuse a data key to encrypt or decrypt messages
	//    before calling AWS KMS again. For more information, see How Does the Data
	//    Key Reuse Period Work? (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work).
	//
	// The following attributes apply only to FIFO (first-in-first-out) queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html):
	//
	//    * FifoQueue - Returns whether the queue is FIFO. For more information,
	//    see FIFO Queue Logic (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-understanding-logic)
	//    in the Amazon Simple Queue Service Developer Guide. To determine whether
	//    a queue is FIFO (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html),
	//    you can check whether QueueName ends with the .fifo suffix.
	//
	//    * ContentBasedDeduplication - Returns whether content-based deduplication
	//    is enabled for the queue. For more information, see Exactly-Once Processing
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	//    in the Amazon Simple Queue Service Developer Guide.
	AttributeNames []enums.QueueAttributeName `locationNameList:"AttributeName" type:"list" flattened:"true"`

	// The URL of the Amazon SQS queue whose attribute information is retrieved.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetQueueAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueueAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetQueueAttributesInput"}

	if s.QueueUrl == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A list of returned queue attributes.
type GetQueueAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A map of attributes to their respective values.
	Attributes map[string]string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
}

// String returns the string representation
func (s GetQueueAttributesOutput) String() string {
	return awsutil.Prettify(s)
}
