// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

const opDeleteRepository = "DeleteRepository"

// DeleteRepositoryRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Deletes an existing image repository. If a repository contains images, you
// must use the force option to delete it.
//
//    // Example sending a request using DeleteRepositoryRequest.
//    req := client.DeleteRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/DeleteRepository
func (c *Client) DeleteRepositoryRequest(input *types.DeleteRepositoryInput) DeleteRepositoryRequest {
	op := &aws.Operation{
		Name:       opDeleteRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteRepositoryInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRepositoryOutput{})
	return DeleteRepositoryRequest{Request: req, Input: input, Copy: c.DeleteRepositoryRequest}
}

// DeleteRepositoryRequest is the request type for the
// DeleteRepository API operation.
type DeleteRepositoryRequest struct {
	*aws.Request
	Input *types.DeleteRepositoryInput
	Copy  func(*types.DeleteRepositoryInput) DeleteRepositoryRequest
}

// Send marshals and sends the DeleteRepository API request.
func (r DeleteRepositoryRequest) Send(ctx context.Context) (*DeleteRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRepositoryResponse{
		DeleteRepositoryOutput: r.Request.Data.(*types.DeleteRepositoryOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRepositoryResponse is the response type for the
// DeleteRepository API operation.
type DeleteRepositoryResponse struct {
	*types.DeleteRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRepository request.
func (r *DeleteRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
