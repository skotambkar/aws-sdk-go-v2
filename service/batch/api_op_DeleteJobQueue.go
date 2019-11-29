// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
)

const opDeleteJobQueue = "DeleteJobQueue"

// DeleteJobQueueRequest returns a request value for making API operation for
// AWS Batch.
//
// Deletes the specified job queue. You must first disable submissions for a
// queue with the UpdateJobQueue operation. All jobs in the queue are terminated
// when you delete a job queue.
//
// It is not necessary to disassociate compute environments from a queue before
// submitting a DeleteJobQueue request.
//
//    // Example sending a request using DeleteJobQueueRequest.
//    req := client.DeleteJobQueueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DeleteJobQueue
func (c *Client) DeleteJobQueueRequest(input *types.DeleteJobQueueInput) DeleteJobQueueRequest {
	op := &aws.Operation{
		Name:       opDeleteJobQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/deletejobqueue",
	}

	if input == nil {
		input = &types.DeleteJobQueueInput{}
	}

	req := c.newRequest(op, input, &types.DeleteJobQueueOutput{})
	return DeleteJobQueueRequest{Request: req, Input: input, Copy: c.DeleteJobQueueRequest}
}

// DeleteJobQueueRequest is the request type for the
// DeleteJobQueue API operation.
type DeleteJobQueueRequest struct {
	*aws.Request
	Input *types.DeleteJobQueueInput
	Copy  func(*types.DeleteJobQueueInput) DeleteJobQueueRequest
}

// Send marshals and sends the DeleteJobQueue API request.
func (r DeleteJobQueueRequest) Send(ctx context.Context) (*DeleteJobQueueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteJobQueueResponse{
		DeleteJobQueueOutput: r.Request.Data.(*types.DeleteJobQueueOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteJobQueueResponse is the response type for the
// DeleteJobQueue API operation.
type DeleteJobQueueResponse struct {
	*types.DeleteJobQueueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteJobQueue request.
func (r *DeleteJobQueueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
