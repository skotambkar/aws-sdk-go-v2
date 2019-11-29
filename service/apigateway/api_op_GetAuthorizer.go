// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetAuthorizer = "GetAuthorizer"

// GetAuthorizerRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Describe an existing Authorizer resource.
//
// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/get-authorizer.html)
//
//    // Example sending a request using GetAuthorizerRequest.
//    req := client.GetAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetAuthorizerRequest(input *types.GetAuthorizerInput) GetAuthorizerRequest {
	op := &aws.Operation{
		Name:       opGetAuthorizer,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/authorizers/{authorizer_id}",
	}

	if input == nil {
		input = &types.GetAuthorizerInput{}
	}

	req := c.newRequest(op, input, &types.GetAuthorizerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetAuthorizerMarshaler{Input: input}.GetNamedBuildHandler())

	return GetAuthorizerRequest{Request: req, Input: input, Copy: c.GetAuthorizerRequest}
}

// GetAuthorizerRequest is the request type for the
// GetAuthorizer API operation.
type GetAuthorizerRequest struct {
	*aws.Request
	Input *types.GetAuthorizerInput
	Copy  func(*types.GetAuthorizerInput) GetAuthorizerRequest
}

// Send marshals and sends the GetAuthorizer API request.
func (r GetAuthorizerRequest) Send(ctx context.Context) (*GetAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAuthorizerResponse{
		GetAuthorizerOutput: r.Request.Data.(*types.GetAuthorizerOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAuthorizerResponse is the response type for the
// GetAuthorizer API operation.
type GetAuthorizerResponse struct {
	*types.GetAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAuthorizer request.
func (r *GetAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
