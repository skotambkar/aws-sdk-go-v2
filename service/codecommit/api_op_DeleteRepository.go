// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opDeleteRepository = "DeleteRepository"

// DeleteRepositoryRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Deletes a repository. If a specified repository was already deleted, a null
// repository ID is returned.
//
// Deleting a repository also deletes all associated objects and metadata. After
// a repository is deleted, all future push calls to the deleted repository
// fail.
//
//    // Example sending a request using DeleteRepositoryRequest.
//    req := client.DeleteRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DeleteRepository
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

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteRepositoryMarshaler{Input: input}.GetNamedBuildHandler())

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
