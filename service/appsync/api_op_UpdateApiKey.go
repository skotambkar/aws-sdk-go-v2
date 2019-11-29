// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/appsync/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
)

const opUpdateApiKey = "UpdateApiKey"

// UpdateApiKeyRequest returns a request value for making API operation for
// AWS AppSync.
//
// Updates an API key.
//
//    // Example sending a request using UpdateApiKeyRequest.
//    req := client.UpdateApiKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateApiKey
func (c *Client) UpdateApiKeyRequest(input *types.UpdateApiKeyInput) UpdateApiKeyRequest {
	op := &aws.Operation{
		Name:       opUpdateApiKey,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/apikeys/{id}",
	}

	if input == nil {
		input = &types.UpdateApiKeyInput{}
	}

	req := c.newRequest(op, input, &types.UpdateApiKeyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateApiKeyMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateApiKeyRequest{Request: req, Input: input, Copy: c.UpdateApiKeyRequest}
}

// UpdateApiKeyRequest is the request type for the
// UpdateApiKey API operation.
type UpdateApiKeyRequest struct {
	*aws.Request
	Input *types.UpdateApiKeyInput
	Copy  func(*types.UpdateApiKeyInput) UpdateApiKeyRequest
}

// Send marshals and sends the UpdateApiKey API request.
func (r UpdateApiKeyRequest) Send(ctx context.Context) (*UpdateApiKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApiKeyResponse{
		UpdateApiKeyOutput: r.Request.Data.(*types.UpdateApiKeyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApiKeyResponse is the response type for the
// UpdateApiKey API operation.
type UpdateApiKeyResponse struct {
	*types.UpdateApiKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApiKey request.
func (r *UpdateApiKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
