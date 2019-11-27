// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opDeleteCodeRepository = "DeleteCodeRepository"

// DeleteCodeRepositoryRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Deletes the specified Git repository from your account.
//
//    // Example sending a request using DeleteCodeRepositoryRequest.
//    req := client.DeleteCodeRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteCodeRepository
func (c *Client) DeleteCodeRepositoryRequest(input *types.DeleteCodeRepositoryInput) DeleteCodeRepositoryRequest {
	op := &aws.Operation{
		Name:       opDeleteCodeRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteCodeRepositoryInput{}
	}

	req := c.newRequest(op, input, &types.DeleteCodeRepositoryOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCodeRepositoryRequest{Request: req, Input: input, Copy: c.DeleteCodeRepositoryRequest}
}

// DeleteCodeRepositoryRequest is the request type for the
// DeleteCodeRepository API operation.
type DeleteCodeRepositoryRequest struct {
	*aws.Request
	Input *types.DeleteCodeRepositoryInput
	Copy  func(*types.DeleteCodeRepositoryInput) DeleteCodeRepositoryRequest
}

// Send marshals and sends the DeleteCodeRepository API request.
func (r DeleteCodeRepositoryRequest) Send(ctx context.Context) (*DeleteCodeRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCodeRepositoryResponse{
		DeleteCodeRepositoryOutput: r.Request.Data.(*types.DeleteCodeRepositoryOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCodeRepositoryResponse is the response type for the
// DeleteCodeRepository API operation.
type DeleteCodeRepositoryResponse struct {
	*types.DeleteCodeRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCodeRepository request.
func (r *DeleteCodeRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
