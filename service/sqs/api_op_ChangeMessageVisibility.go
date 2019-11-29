// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const opChangeMessageVisibility = "ChangeMessageVisibility"

// ChangeMessageVisibilityRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Changes the visibility timeout of a specified message in a queue to a new
// value. The default visibility timeout for a message is 30 seconds. The minimum
// is 0 seconds. The maximum is 12 hours. For more information, see Visibility
// Timeout (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// For example, you have a message with a visibility timeout of 5 minutes. After
// 3 minutes, you call ChangeMessageVisibility with a timeout of 10 minutes.
// You can continue to call ChangeMessageVisibility to extend the visibility
// timeout to the maximum allowed time. If you try to extend the visibility
// timeout beyond the maximum, your request is rejected.
//
// An Amazon SQS message has three basic states:
//
// Sent to a queue by a producer.
//
// Received from the queue by a consumer.
//
// Deleted from the queue.
//
// A message is considered to be stored after it is sent to a queue by a producer,
// but not yet received from the queue by a consumer (that is, between states
// 1 and 2). There is no limit to the number of stored messages. A message is
// considered to be in flight after it is received from a queue by a consumer,
// but not yet deleted from the queue (that is, between states 2 and 3). There
// is a limit to the number of inflight messages.
//
// Limits that apply to inflight messages are unrelated to the unlimited number
// of stored messages.
//
// For most standard queues (depending on queue traffic and message backlog),
// there can be a maximum of approximately 120,000 inflight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns the OverLimit error message. To avoid reaching
// the limit, you should delete messages from the queue after they're processed.
// You can also increase the number of queues you use to process your messages.
// To request a limit increase, file a support request (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-sqs).
//
// For FIFO queues, there can be a maximum of 20,000 inflight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns no error messages.
//
// If you attempt to set the VisibilityTimeout to a value greater than the maximum
// time left, Amazon SQS returns an error. Amazon SQS doesn't automatically
// recalculate and increase the timeout to the maximum remaining time.
//
// Unlike with a queue, when you change the visibility timeout for a specific
// message the timeout value is applied immediately but isn't saved in memory
// for that message. If you don't delete a message after it is received, the
// visibility timeout for the message reverts to the original timeout value
// (not to the value you set using the ChangeMessageVisibility action) the next
// time the message is received.
//
//    // Example sending a request using ChangeMessageVisibilityRequest.
//    req := client.ChangeMessageVisibilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ChangeMessageVisibility
func (c *Client) ChangeMessageVisibilityRequest(input *types.ChangeMessageVisibilityInput) ChangeMessageVisibilityRequest {
	op := &aws.Operation{
		Name:       opChangeMessageVisibility,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ChangeMessageVisibilityInput{}
	}

	req := c.newRequest(op, input, &types.ChangeMessageVisibilityOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.ChangeMessageVisibilityMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ChangeMessageVisibilityRequest{Request: req, Input: input, Copy: c.ChangeMessageVisibilityRequest}
}

// ChangeMessageVisibilityRequest is the request type for the
// ChangeMessageVisibility API operation.
type ChangeMessageVisibilityRequest struct {
	*aws.Request
	Input *types.ChangeMessageVisibilityInput
	Copy  func(*types.ChangeMessageVisibilityInput) ChangeMessageVisibilityRequest
}

// Send marshals and sends the ChangeMessageVisibility API request.
func (r ChangeMessageVisibilityRequest) Send(ctx context.Context) (*ChangeMessageVisibilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ChangeMessageVisibilityResponse{
		ChangeMessageVisibilityOutput: r.Request.Data.(*types.ChangeMessageVisibilityOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ChangeMessageVisibilityResponse is the response type for the
// ChangeMessageVisibility API operation.
type ChangeMessageVisibilityResponse struct {
	*types.ChangeMessageVisibilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ChangeMessageVisibility request.
func (r *ChangeMessageVisibilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
