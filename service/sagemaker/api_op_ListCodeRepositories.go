// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opListCodeRepositories = "ListCodeRepositories"

// ListCodeRepositoriesRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Gets a list of the Git repositories in your account.
//
//    // Example sending a request using ListCodeRepositoriesRequest.
//    req := client.ListCodeRepositoriesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListCodeRepositories
func (c *Client) ListCodeRepositoriesRequest(input *types.ListCodeRepositoriesInput) ListCodeRepositoriesRequest {
	op := &aws.Operation{
		Name:       opListCodeRepositories,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListCodeRepositoriesInput{}
	}

	req := c.newRequest(op, input, &types.ListCodeRepositoriesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListCodeRepositoriesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListCodeRepositoriesRequest{Request: req, Input: input, Copy: c.ListCodeRepositoriesRequest}
}

// ListCodeRepositoriesRequest is the request type for the
// ListCodeRepositories API operation.
type ListCodeRepositoriesRequest struct {
	*aws.Request
	Input *types.ListCodeRepositoriesInput
	Copy  func(*types.ListCodeRepositoriesInput) ListCodeRepositoriesRequest
}

// Send marshals and sends the ListCodeRepositories API request.
func (r ListCodeRepositoriesRequest) Send(ctx context.Context) (*ListCodeRepositoriesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCodeRepositoriesResponse{
		ListCodeRepositoriesOutput: r.Request.Data.(*types.ListCodeRepositoriesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListCodeRepositoriesResponse is the response type for the
// ListCodeRepositories API operation.
type ListCodeRepositoriesResponse struct {
	*types.ListCodeRepositoriesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCodeRepositories request.
func (r *ListCodeRepositoriesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
