// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opUpdateIntegration = "UpdateIntegration"

// UpdateIntegrationRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Represents an update integration.
//
//    // Example sending a request using UpdateIntegrationRequest.
//    req := client.UpdateIntegrationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateIntegrationRequest(input *types.UpdateIntegrationInput) UpdateIntegrationRequest {
	op := &aws.Operation{
		Name:       opUpdateIntegration,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}/integration",
	}

	if input == nil {
		input = &types.UpdateIntegrationInput{}
	}

	req := c.newRequest(op, input, &types.UpdateIntegrationOutput{})
	return UpdateIntegrationRequest{Request: req, Input: input, Copy: c.UpdateIntegrationRequest}
}

// UpdateIntegrationRequest is the request type for the
// UpdateIntegration API operation.
type UpdateIntegrationRequest struct {
	*aws.Request
	Input *types.UpdateIntegrationInput
	Copy  func(*types.UpdateIntegrationInput) UpdateIntegrationRequest
}

// Send marshals and sends the UpdateIntegration API request.
func (r UpdateIntegrationRequest) Send(ctx context.Context) (*UpdateIntegrationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateIntegrationResponse{
		UpdateIntegrationOutput: r.Request.Data.(*types.UpdateIntegrationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateIntegrationResponse is the response type for the
// UpdateIntegration API operation.
type UpdateIntegrationResponse struct {
	*types.UpdateIntegrationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateIntegration request.
func (r *UpdateIntegrationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
