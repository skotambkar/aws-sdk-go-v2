// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetApiKey = "GetApiKey"

// GetApiKeyRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets information about the current ApiKey resource.
//
//    // Example sending a request using GetApiKeyRequest.
//    req := client.GetApiKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetApiKeyRequest(input *types.GetApiKeyInput) GetApiKeyRequest {
	op := &aws.Operation{
		Name:       opGetApiKey,
		HTTPMethod: "GET",
		HTTPPath:   "/apikeys/{api_Key}",
	}

	if input == nil {
		input = &types.GetApiKeyInput{}
	}

	req := c.newRequest(op, input, &types.GetApiKeyOutput{})
	return GetApiKeyRequest{Request: req, Input: input, Copy: c.GetApiKeyRequest}
}

// GetApiKeyRequest is the request type for the
// GetApiKey API operation.
type GetApiKeyRequest struct {
	*aws.Request
	Input *types.GetApiKeyInput
	Copy  func(*types.GetApiKeyInput) GetApiKeyRequest
}

// Send marshals and sends the GetApiKey API request.
func (r GetApiKeyRequest) Send(ctx context.Context) (*GetApiKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetApiKeyResponse{
		GetApiKeyOutput: r.Request.Data.(*types.GetApiKeyOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetApiKeyResponse is the response type for the
// GetApiKey API operation.
type GetApiKeyResponse struct {
	*types.GetApiKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetApiKey request.
func (r *GetApiKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
