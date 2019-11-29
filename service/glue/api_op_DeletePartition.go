// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opDeletePartition = "DeletePartition"

// DeletePartitionRequest returns a request value for making API operation for
// AWS Glue.
//
// Deletes a specified partition.
//
//    // Example sending a request using DeletePartitionRequest.
//    req := client.DeletePartitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/DeletePartition
func (c *Client) DeletePartitionRequest(input *types.DeletePartitionInput) DeletePartitionRequest {
	op := &aws.Operation{
		Name:       opDeletePartition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeletePartitionInput{}
	}

	req := c.newRequest(op, input, &types.DeletePartitionOutput{})
	return DeletePartitionRequest{Request: req, Input: input, Copy: c.DeletePartitionRequest}
}

// DeletePartitionRequest is the request type for the
// DeletePartition API operation.
type DeletePartitionRequest struct {
	*aws.Request
	Input *types.DeletePartitionInput
	Copy  func(*types.DeletePartitionInput) DeletePartitionRequest
}

// Send marshals and sends the DeletePartition API request.
func (r DeletePartitionRequest) Send(ctx context.Context) (*DeletePartitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePartitionResponse{
		DeletePartitionOutput: r.Request.Data.(*types.DeletePartitionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePartitionResponse is the response type for the
// DeletePartition API operation.
type DeletePartitionResponse struct {
	*types.DeletePartitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePartition request.
func (r *DeletePartitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
