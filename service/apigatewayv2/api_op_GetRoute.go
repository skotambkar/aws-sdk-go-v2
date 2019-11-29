// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opGetRoute = "GetRoute"

// GetRouteRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets a Route.
//
//    // Example sending a request using GetRouteRequest.
//    req := client.GetRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetRoute
func (c *Client) GetRouteRequest(input *types.GetRouteInput) GetRouteRequest {
	op := &aws.Operation{
		Name:       opGetRoute,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/routes/{routeId}",
	}

	if input == nil {
		input = &types.GetRouteInput{}
	}

	req := c.newRequest(op, input, &types.GetRouteOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetRouteMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRouteRequest{Request: req, Input: input, Copy: c.GetRouteRequest}
}

// GetRouteRequest is the request type for the
// GetRoute API operation.
type GetRouteRequest struct {
	*aws.Request
	Input *types.GetRouteInput
	Copy  func(*types.GetRouteInput) GetRouteRequest
}

// Send marshals and sends the GetRoute API request.
func (r GetRouteRequest) Send(ctx context.Context) (*GetRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRouteResponse{
		GetRouteOutput: r.Request.Data.(*types.GetRouteOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRouteResponse is the response type for the
// GetRoute API operation.
type GetRouteResponse struct {
	*types.GetRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRoute request.
func (r *GetRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
