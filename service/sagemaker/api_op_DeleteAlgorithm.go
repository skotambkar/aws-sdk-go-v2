// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opDeleteAlgorithm = "DeleteAlgorithm"

// DeleteAlgorithmRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Removes the specified algorithm from your account.
//
//    // Example sending a request using DeleteAlgorithmRequest.
//    req := client.DeleteAlgorithmRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteAlgorithm
func (c *Client) DeleteAlgorithmRequest(input *types.DeleteAlgorithmInput) DeleteAlgorithmRequest {
	op := &aws.Operation{
		Name:       opDeleteAlgorithm,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteAlgorithmInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAlgorithmOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAlgorithmRequest{Request: req, Input: input, Copy: c.DeleteAlgorithmRequest}
}

// DeleteAlgorithmRequest is the request type for the
// DeleteAlgorithm API operation.
type DeleteAlgorithmRequest struct {
	*aws.Request
	Input *types.DeleteAlgorithmInput
	Copy  func(*types.DeleteAlgorithmInput) DeleteAlgorithmRequest
}

// Send marshals and sends the DeleteAlgorithm API request.
func (r DeleteAlgorithmRequest) Send(ctx context.Context) (*DeleteAlgorithmResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAlgorithmResponse{
		DeleteAlgorithmOutput: r.Request.Data.(*types.DeleteAlgorithmOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAlgorithmResponse is the response type for the
// DeleteAlgorithm API operation.
type DeleteAlgorithmResponse struct {
	*types.DeleteAlgorithmOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAlgorithm request.
func (r *DeleteAlgorithmResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
