// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opCreateApiKey = "CreateApiKey"

// CreateApiKeyRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Create an ApiKey resource.
//
// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/create-api-key.html)
//
//    // Example sending a request using CreateApiKeyRequest.
//    req := client.CreateApiKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateApiKeyRequest(input *types.CreateApiKeyInput) CreateApiKeyRequest {
	op := &aws.Operation{
		Name:       opCreateApiKey,
		HTTPMethod: "POST",
		HTTPPath:   "/apikeys",
	}

	if input == nil {
		input = &types.CreateApiKeyInput{}
	}

	req := c.newRequest(op, input, &types.CreateApiKeyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateApiKeyMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateApiKeyRequest{Request: req, Input: input, Copy: c.CreateApiKeyRequest}
}

// CreateApiKeyRequest is the request type for the
// CreateApiKey API operation.
type CreateApiKeyRequest struct {
	*aws.Request
	Input *types.CreateApiKeyInput
	Copy  func(*types.CreateApiKeyInput) CreateApiKeyRequest
}

// Send marshals and sends the CreateApiKey API request.
func (r CreateApiKeyRequest) Send(ctx context.Context) (*CreateApiKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApiKeyResponse{
		CreateApiKeyOutput: r.Request.Data.(*types.CreateApiKeyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApiKeyResponse is the response type for the
// CreateApiKey API operation.
type CreateApiKeyResponse struct {
	*types.CreateApiKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApiKey request.
func (r *CreateApiKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
