// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opCreateIntegrationResponse = "CreateIntegrationResponse"

// CreateIntegrationResponseRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Creates an IntegrationResponses.
//
//    // Example sending a request using CreateIntegrationResponseRequest.
//    req := client.CreateIntegrationResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateIntegrationResponse
func (c *Client) CreateIntegrationResponseRequest(input *types.CreateIntegrationResponseInput) CreateIntegrationResponseRequest {
	op := &aws.Operation{
		Name:       opCreateIntegrationResponse,
		HTTPMethod: "POST",
		HTTPPath:   "/v2/apis/{apiId}/integrations/{integrationId}/integrationresponses",
	}

	if input == nil {
		input = &types.CreateIntegrationResponseInput{}
	}

	req := c.newRequest(op, input, &types.CreateIntegrationResponseOutput{})
	return CreateIntegrationResponseRequest{Request: req, Input: input, Copy: c.CreateIntegrationResponseRequest}
}

// CreateIntegrationResponseRequest is the request type for the
// CreateIntegrationResponse API operation.
type CreateIntegrationResponseRequest struct {
	*aws.Request
	Input *types.CreateIntegrationResponseInput
	Copy  func(*types.CreateIntegrationResponseInput) CreateIntegrationResponseRequest
}

// Send marshals and sends the CreateIntegrationResponse API request.
func (r CreateIntegrationResponseRequest) Send(ctx context.Context) (*CreateIntegrationResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateIntegrationResponseResponse{
		CreateIntegrationResponseOutput: r.Request.Data.(*types.CreateIntegrationResponseOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateIntegrationResponseResponse is the response type for the
// CreateIntegrationResponse API operation.
type CreateIntegrationResponseResponse struct {
	*types.CreateIntegrationResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateIntegrationResponse request.
func (r *CreateIntegrationResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
