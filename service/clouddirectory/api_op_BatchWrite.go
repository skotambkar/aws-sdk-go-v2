// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opBatchWrite = "BatchWrite"

// BatchWriteRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Performs all the write operations in a batch. Either all the operations succeed
// or none.
//
//    // Example sending a request using BatchWriteRequest.
//    req := client.BatchWriteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/BatchWrite
func (c *Client) BatchWriteRequest(input *types.BatchWriteInput) BatchWriteRequest {
	op := &aws.Operation{
		Name:       opBatchWrite,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/batchwrite",
	}

	if input == nil {
		input = &types.BatchWriteInput{}
	}

	req := c.newRequest(op, input, &types.BatchWriteOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.BatchWriteMarshaler{Input: input}.GetNamedBuildHandler())

	return BatchWriteRequest{Request: req, Input: input, Copy: c.BatchWriteRequest}
}

// BatchWriteRequest is the request type for the
// BatchWrite API operation.
type BatchWriteRequest struct {
	*aws.Request
	Input *types.BatchWriteInput
	Copy  func(*types.BatchWriteInput) BatchWriteRequest
}

// Send marshals and sends the BatchWrite API request.
func (r BatchWriteRequest) Send(ctx context.Context) (*BatchWriteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchWriteResponse{
		BatchWriteOutput: r.Request.Data.(*types.BatchWriteOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchWriteResponse is the response type for the
// BatchWrite API operation.
type BatchWriteResponse struct {
	*types.BatchWriteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchWrite request.
func (r *BatchWriteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
