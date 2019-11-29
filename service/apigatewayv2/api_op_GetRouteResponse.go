// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opGetRouteResponse = "GetRouteResponse"

// GetRouteResponseRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets a RouteResponse.
//
//    // Example sending a request using GetRouteResponseRequest.
//    req := client.GetRouteResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetRouteResponse
func (c *Client) GetRouteResponseRequest(input *types.GetRouteResponseInput) GetRouteResponseRequest {
	op := &aws.Operation{
		Name:       opGetRouteResponse,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}/routeresponses/{routeResponseId}",
	}

	if input == nil {
		input = &types.GetRouteResponseInput{}
	}

	req := c.newRequest(op, input, &types.GetRouteResponseOutput{})
	return GetRouteResponseRequest{Request: req, Input: input, Copy: c.GetRouteResponseRequest}
}

// GetRouteResponseRequest is the request type for the
// GetRouteResponse API operation.
type GetRouteResponseRequest struct {
	*aws.Request
	Input *types.GetRouteResponseInput
	Copy  func(*types.GetRouteResponseInput) GetRouteResponseRequest
}

// Send marshals and sends the GetRouteResponse API request.
func (r GetRouteResponseRequest) Send(ctx context.Context) (*GetRouteResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRouteResponseResponse{
		GetRouteResponseOutput: r.Request.Data.(*types.GetRouteResponseOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRouteResponseResponse is the response type for the
// GetRouteResponse API operation.
type GetRouteResponseResponse struct {
	*types.GetRouteResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRouteResponse request.
func (r *GetRouteResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
