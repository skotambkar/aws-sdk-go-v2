// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const opCreateQueue = "CreateQueue"

// CreateQueueRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// Creates a new standard or FIFO queue. You can pass one or more attributes
// in the request. Keep the following caveats in mind:
//
//    * If you don't specify the FifoQueue attribute, Amazon SQS creates a standard
//    queue. You can't change the queue type after you create it and you can't
//    convert an existing standard queue into a FIFO queue. You must either
//    create a new FIFO queue for your application or delete your existing standard
//    queue and recreate it as a FIFO queue. For more information, see Moving
//    From a Standard Queue to a FIFO Queue (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-moving)
//    in the Amazon Simple Queue Service Developer Guide.
//
//    * If you don't provide a value for an attribute, the queue is created
//    with the default value for the attribute.
//
//    * If you delete a queue, you must wait at least 60 seconds before creating
//    a queue with the same name.
//
// To successfully create a new queue, you must provide a queue name that adheres
// to the limits related to queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html)
// and is unique within the scope of your queues.
//
// To get the queue URL, use the GetQueueUrl action. GetQueueUrl requires only
// the QueueName parameter. be aware of existing queue names:
//
//    * If you provide the name of an existing queue along with the exact names
//    and values of all the queue's attributes, CreateQueue returns the queue
//    URL for the existing queue.
//
//    * If the queue name, attribute names, or attribute values don't match
//    an existing queue, CreateQueue returns an error.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
//    // Example sending a request using CreateQueueRequest.
//    req := client.CreateQueueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/CreateQueue
func (c *Client) CreateQueueRequest(input *types.CreateQueueInput) CreateQueueRequest {
	op := &aws.Operation{
		Name:       opCreateQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateQueueInput{}
	}

	req := c.newRequest(op, input, &types.CreateQueueOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateQueueMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateQueueRequest{Request: req, Input: input, Copy: c.CreateQueueRequest}
}

// CreateQueueRequest is the request type for the
// CreateQueue API operation.
type CreateQueueRequest struct {
	*aws.Request
	Input *types.CreateQueueInput
	Copy  func(*types.CreateQueueInput) CreateQueueRequest
}

// Send marshals and sends the CreateQueue API request.
func (r CreateQueueRequest) Send(ctx context.Context) (*CreateQueueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateQueueResponse{
		CreateQueueOutput: r.Request.Data.(*types.CreateQueueOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateQueueResponse is the response type for the
// CreateQueue API operation.
type CreateQueueResponse struct {
	*types.CreateQueueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateQueue request.
func (r *CreateQueueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
