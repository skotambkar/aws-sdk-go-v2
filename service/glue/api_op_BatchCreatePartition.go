// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opBatchCreatePartition = "BatchCreatePartition"

// BatchCreatePartitionRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates one or more partitions in a batch operation.
//
//    // Example sending a request using BatchCreatePartitionRequest.
//    req := client.BatchCreatePartitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchCreatePartition
func (c *Client) BatchCreatePartitionRequest(input *types.BatchCreatePartitionInput) BatchCreatePartitionRequest {
	op := &aws.Operation{
		Name:       opBatchCreatePartition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchCreatePartitionInput{}
	}

	req := c.newRequest(op, input, &types.BatchCreatePartitionOutput{})
	return BatchCreatePartitionRequest{Request: req, Input: input, Copy: c.BatchCreatePartitionRequest}
}

// BatchCreatePartitionRequest is the request type for the
// BatchCreatePartition API operation.
type BatchCreatePartitionRequest struct {
	*aws.Request
	Input *types.BatchCreatePartitionInput
	Copy  func(*types.BatchCreatePartitionInput) BatchCreatePartitionRequest
}

// Send marshals and sends the BatchCreatePartition API request.
func (r BatchCreatePartitionRequest) Send(ctx context.Context) (*BatchCreatePartitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchCreatePartitionResponse{
		BatchCreatePartitionOutput: r.Request.Data.(*types.BatchCreatePartitionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchCreatePartitionResponse is the response type for the
// BatchCreatePartition API operation.
type BatchCreatePartitionResponse struct {
	*types.BatchCreatePartitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchCreatePartition request.
func (r *BatchCreatePartitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
