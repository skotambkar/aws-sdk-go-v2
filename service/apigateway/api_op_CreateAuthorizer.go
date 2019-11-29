// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opCreateAuthorizer = "CreateAuthorizer"

// CreateAuthorizerRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Adds a new Authorizer resource to an existing RestApi resource.
//
// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/create-authorizer.html)
//
//    // Example sending a request using CreateAuthorizerRequest.
//    req := client.CreateAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateAuthorizerRequest(input *types.CreateAuthorizerInput) CreateAuthorizerRequest {
	op := &aws.Operation{
		Name:       opCreateAuthorizer,
		HTTPMethod: "POST",
		HTTPPath:   "/restapis/{restapi_id}/authorizers",
	}

	if input == nil {
		input = &types.CreateAuthorizerInput{}
	}

	req := c.newRequest(op, input, &types.CreateAuthorizerOutput{})
	return CreateAuthorizerRequest{Request: req, Input: input, Copy: c.CreateAuthorizerRequest}
}

// CreateAuthorizerRequest is the request type for the
// CreateAuthorizer API operation.
type CreateAuthorizerRequest struct {
	*aws.Request
	Input *types.CreateAuthorizerInput
	Copy  func(*types.CreateAuthorizerInput) CreateAuthorizerRequest
}

// Send marshals and sends the CreateAuthorizer API request.
func (r CreateAuthorizerRequest) Send(ctx context.Context) (*CreateAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAuthorizerResponse{
		CreateAuthorizerOutput: r.Request.Data.(*types.CreateAuthorizerOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAuthorizerResponse is the response type for the
// CreateAuthorizer API operation.
type CreateAuthorizerResponse struct {
	*types.CreateAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAuthorizer request.
func (r *CreateAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
