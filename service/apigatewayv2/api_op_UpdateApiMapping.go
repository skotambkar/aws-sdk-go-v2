// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opUpdateApiMapping = "UpdateApiMapping"

// UpdateApiMappingRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// The API mapping.
//
//    // Example sending a request using UpdateApiMappingRequest.
//    req := client.UpdateApiMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateApiMapping
func (c *Client) UpdateApiMappingRequest(input *types.UpdateApiMappingInput) UpdateApiMappingRequest {
	op := &aws.Operation{
		Name:       opUpdateApiMapping,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/domainnames/{domainName}/apimappings/{apiMappingId}",
	}

	if input == nil {
		input = &types.UpdateApiMappingInput{}
	}

	req := c.newRequest(op, input, &types.UpdateApiMappingOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateApiMappingMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateApiMappingRequest{Request: req, Input: input, Copy: c.UpdateApiMappingRequest}
}

// UpdateApiMappingRequest is the request type for the
// UpdateApiMapping API operation.
type UpdateApiMappingRequest struct {
	*aws.Request
	Input *types.UpdateApiMappingInput
	Copy  func(*types.UpdateApiMappingInput) UpdateApiMappingRequest
}

// Send marshals and sends the UpdateApiMapping API request.
func (r UpdateApiMappingRequest) Send(ctx context.Context) (*UpdateApiMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApiMappingResponse{
		UpdateApiMappingOutput: r.Request.Data.(*types.UpdateApiMappingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApiMappingResponse is the response type for the
// UpdateApiMapping API operation.
type UpdateApiMappingResponse struct {
	*types.UpdateApiMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApiMapping request.
func (r *UpdateApiMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
