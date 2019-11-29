// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opUpdateMethodResponse = "UpdateMethodResponse"

// UpdateMethodResponseRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Updates an existing MethodResponse resource.
//
//    // Example sending a request using UpdateMethodResponseRequest.
//    req := client.UpdateMethodResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateMethodResponseRequest(input *types.UpdateMethodResponseInput) UpdateMethodResponseRequest {
	op := &aws.Operation{
		Name:       opUpdateMethodResponse,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}/responses/{status_code}",
	}

	if input == nil {
		input = &types.UpdateMethodResponseInput{}
	}

	req := c.newRequest(op, input, &types.UpdateMethodResponseOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateMethodResponseMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateMethodResponseRequest{Request: req, Input: input, Copy: c.UpdateMethodResponseRequest}
}

// UpdateMethodResponseRequest is the request type for the
// UpdateMethodResponse API operation.
type UpdateMethodResponseRequest struct {
	*aws.Request
	Input *types.UpdateMethodResponseInput
	Copy  func(*types.UpdateMethodResponseInput) UpdateMethodResponseRequest
}

// Send marshals and sends the UpdateMethodResponse API request.
func (r UpdateMethodResponseRequest) Send(ctx context.Context) (*UpdateMethodResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMethodResponseResponse{
		UpdateMethodResponseOutput: r.Request.Data.(*types.UpdateMethodResponseOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMethodResponseResponse is the response type for the
// UpdateMethodResponse API operation.
type UpdateMethodResponseResponse struct {
	*types.UpdateMethodResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMethodResponse request.
func (r *UpdateMethodResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
