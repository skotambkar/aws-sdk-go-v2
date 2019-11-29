// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opBatchDeleteTable = "BatchDeleteTable"

// BatchDeleteTableRequest returns a request value for making API operation for
// AWS Glue.
//
// Deletes multiple tables at once.
//
// After completing this operation, you no longer have access to the table versions
// and partitions that belong to the deleted table. AWS Glue deletes these "orphaned"
// resources asynchronously in a timely manner, at the discretion of the service.
//
// To ensure the immediate deletion of all related resources, before calling
// BatchDeleteTable, use DeleteTableVersion or BatchDeleteTableVersion, and
// DeletePartition or BatchDeletePartition, to delete any resources that belong
// to the table.
//
//    // Example sending a request using BatchDeleteTableRequest.
//    req := client.BatchDeleteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchDeleteTable
func (c *Client) BatchDeleteTableRequest(input *types.BatchDeleteTableInput) BatchDeleteTableRequest {
	op := &aws.Operation{
		Name:       opBatchDeleteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchDeleteTableInput{}
	}

	req := c.newRequest(op, input, &types.BatchDeleteTableOutput{})
	return BatchDeleteTableRequest{Request: req, Input: input, Copy: c.BatchDeleteTableRequest}
}

// BatchDeleteTableRequest is the request type for the
// BatchDeleteTable API operation.
type BatchDeleteTableRequest struct {
	*aws.Request
	Input *types.BatchDeleteTableInput
	Copy  func(*types.BatchDeleteTableInput) BatchDeleteTableRequest
}

// Send marshals and sends the BatchDeleteTable API request.
func (r BatchDeleteTableRequest) Send(ctx context.Context) (*BatchDeleteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeleteTableResponse{
		BatchDeleteTableOutput: r.Request.Data.(*types.BatchDeleteTableOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeleteTableResponse is the response type for the
// BatchDeleteTable API operation.
type BatchDeleteTableResponse struct {
	*types.BatchDeleteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeleteTable request.
func (r *BatchDeleteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
