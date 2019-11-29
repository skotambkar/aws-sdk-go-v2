// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const opReceiveMessage = "ReceiveMessage"

// ReceiveMessageRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Retrieves one or more messages (up to 10), from the specified queue. Using
// the WaitTimeSeconds parameter enables long-poll support. For more information,
// see Amazon SQS Long Polling (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// Short poll is the default behavior where a weighted random set of machines
// is sampled on a ReceiveMessage call. Thus, only the messages on the sampled
// machines are returned. If the number of messages in the queue is small (fewer
// than 1,000), you most likely get fewer messages than you requested per ReceiveMessage
// call. If the number of messages in the queue is extremely small, you might
// not receive any messages in a particular ReceiveMessage response. If this
// happens, repeat the request.
//
// For each message returned, the response includes the following:
//
//    * The message body.
//
//    * An MD5 digest of the message body. For information about MD5, see RFC1321
//    (https://www.ietf.org/rfc/rfc1321.txt).
//
//    * The MessageId you received when you sent the message to the queue.
//
//    * The receipt handle.
//
//    * The message attributes.
//
//    * An MD5 digest of the message attributes.
//
// The receipt handle is the identifier you must provide when deleting the message.
// For more information, see Queue and Message Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// You can provide the VisibilityTimeout parameter in your request. The parameter
// is applied to the messages that Amazon SQS returns in the response. If you
// don't include the parameter, the overall visibility timeout for the queue
// is used for the returned messages. For more information, see Visibility Timeout
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// A message that isn't deleted or a message whose visibility isn't extended
// before the visibility timeout expires counts as a failed receive. Depending
// on the configuration of the queue, the message might be sent to the dead-letter
// queue.
//
// In the future, new attributes might be added. If you write code that calls
// this action, we recommend that you structure your code so that it can handle
// new attributes gracefully.
//
//    // Example sending a request using ReceiveMessageRequest.
//    req := client.ReceiveMessageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ReceiveMessage
func (c *Client) ReceiveMessageRequest(input *types.ReceiveMessageInput) ReceiveMessageRequest {
	op := &aws.Operation{
		Name:       opReceiveMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ReceiveMessageInput{}
	}

	req := c.newRequest(op, input, &types.ReceiveMessageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.ReceiveMessageMarshaler{Input: input}.GetNamedBuildHandler())

	return ReceiveMessageRequest{Request: req, Input: input, Copy: c.ReceiveMessageRequest}
}

// ReceiveMessageRequest is the request type for the
// ReceiveMessage API operation.
type ReceiveMessageRequest struct {
	*aws.Request
	Input *types.ReceiveMessageInput
	Copy  func(*types.ReceiveMessageInput) ReceiveMessageRequest
}

// Send marshals and sends the ReceiveMessage API request.
func (r ReceiveMessageRequest) Send(ctx context.Context) (*ReceiveMessageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReceiveMessageResponse{
		ReceiveMessageOutput: r.Request.Data.(*types.ReceiveMessageOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReceiveMessageResponse is the response type for the
// ReceiveMessage API operation.
type ReceiveMessageResponse struct {
	*types.ReceiveMessageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReceiveMessage request.
func (r *ReceiveMessageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
