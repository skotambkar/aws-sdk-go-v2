// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opUpdateAuthorizer = "UpdateAuthorizer"

// UpdateAuthorizerRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Updates an Authorizer.
//
//    // Example sending a request using UpdateAuthorizerRequest.
//    req := client.UpdateAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateAuthorizer
func (c *Client) UpdateAuthorizerRequest(input *types.UpdateAuthorizerInput) UpdateAuthorizerRequest {
	op := &aws.Operation{
		Name:       opUpdateAuthorizer,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/apis/{apiId}/authorizers/{authorizerId}",
	}

	if input == nil {
		input = &types.UpdateAuthorizerInput{}
	}

	req := c.newRequest(op, input, &types.UpdateAuthorizerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateAuthorizerMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateAuthorizerRequest{Request: req, Input: input, Copy: c.UpdateAuthorizerRequest}
}

// UpdateAuthorizerRequest is the request type for the
// UpdateAuthorizer API operation.
type UpdateAuthorizerRequest struct {
	*aws.Request
	Input *types.UpdateAuthorizerInput
	Copy  func(*types.UpdateAuthorizerInput) UpdateAuthorizerRequest
}

// Send marshals and sends the UpdateAuthorizer API request.
func (r UpdateAuthorizerRequest) Send(ctx context.Context) (*UpdateAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAuthorizerResponse{
		UpdateAuthorizerOutput: r.Request.Data.(*types.UpdateAuthorizerOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAuthorizerResponse is the response type for the
// UpdateAuthorizer API operation.
type UpdateAuthorizerResponse struct {
	*types.UpdateAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAuthorizer request.
func (r *UpdateAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
