// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opDeleteRemoteAccessSession = "DeleteRemoteAccessSession"

// DeleteRemoteAccessSessionRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Deletes a completed remote access session and its results.
//
//    // Example sending a request using DeleteRemoteAccessSessionRequest.
//    req := client.DeleteRemoteAccessSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/DeleteRemoteAccessSession
func (c *Client) DeleteRemoteAccessSessionRequest(input *types.DeleteRemoteAccessSessionInput) DeleteRemoteAccessSessionRequest {
	op := &aws.Operation{
		Name:       opDeleteRemoteAccessSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteRemoteAccessSessionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRemoteAccessSessionOutput{})
	return DeleteRemoteAccessSessionRequest{Request: req, Input: input, Copy: c.DeleteRemoteAccessSessionRequest}
}

// DeleteRemoteAccessSessionRequest is the request type for the
// DeleteRemoteAccessSession API operation.
type DeleteRemoteAccessSessionRequest struct {
	*aws.Request
	Input *types.DeleteRemoteAccessSessionInput
	Copy  func(*types.DeleteRemoteAccessSessionInput) DeleteRemoteAccessSessionRequest
}

// Send marshals and sends the DeleteRemoteAccessSession API request.
func (r DeleteRemoteAccessSessionRequest) Send(ctx context.Context) (*DeleteRemoteAccessSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRemoteAccessSessionResponse{
		DeleteRemoteAccessSessionOutput: r.Request.Data.(*types.DeleteRemoteAccessSessionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRemoteAccessSessionResponse is the response type for the
// DeleteRemoteAccessSession API operation.
type DeleteRemoteAccessSessionResponse struct {
	*types.DeleteRemoteAccessSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRemoteAccessSession request.
func (r *DeleteRemoteAccessSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
