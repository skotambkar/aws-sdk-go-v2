// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opGetRouteResponses = "GetRouteResponses"

// GetRouteResponsesRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets the RouteResponses for a Route.
//
//    // Example sending a request using GetRouteResponsesRequest.
//    req := client.GetRouteResponsesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetRouteResponses
func (c *Client) GetRouteResponsesRequest(input *types.GetRouteResponsesInput) GetRouteResponsesRequest {
	op := &aws.Operation{
		Name:       opGetRouteResponses,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}/routeresponses",
	}

	if input == nil {
		input = &types.GetRouteResponsesInput{}
	}

	req := c.newRequest(op, input, &types.GetRouteResponsesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetRouteResponsesMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRouteResponsesRequest{Request: req, Input: input, Copy: c.GetRouteResponsesRequest}
}

// GetRouteResponsesRequest is the request type for the
// GetRouteResponses API operation.
type GetRouteResponsesRequest struct {
	*aws.Request
	Input *types.GetRouteResponsesInput
	Copy  func(*types.GetRouteResponsesInput) GetRouteResponsesRequest
}

// Send marshals and sends the GetRouteResponses API request.
func (r GetRouteResponsesRequest) Send(ctx context.Context) (*GetRouteResponsesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRouteResponsesResponse{
		GetRouteResponsesOutput: r.Request.Data.(*types.GetRouteResponsesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRouteResponsesResponse is the response type for the
// GetRouteResponses API operation.
type GetRouteResponsesResponse struct {
	*types.GetRouteResponsesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRouteResponses request.
func (r *GetRouteResponsesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
