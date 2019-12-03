// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetGatewayResponse = "GetGatewayResponse"

// GetGatewayResponseRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets a GatewayResponse of a specified response type on the given RestApi.
//
//    // Example sending a request using GetGatewayResponseRequest.
//    req := client.GetGatewayResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetGatewayResponseRequest(input *types.GetGatewayResponseInput) GetGatewayResponseRequest {
	op := &aws.Operation{
		Name:       opGetGatewayResponse,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/gatewayresponses/{response_type}",
	}

	if input == nil {
		input = &types.GetGatewayResponseInput{}
	}

	req := c.newRequest(op, input, &types.GetGatewayResponseOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetGatewayResponseMarshaler{Input: input}.GetNamedBuildHandler())

	return GetGatewayResponseRequest{Request: req, Input: input, Copy: c.GetGatewayResponseRequest}
}

// GetGatewayResponseRequest is the request type for the
// GetGatewayResponse API operation.
type GetGatewayResponseRequest struct {
	*aws.Request
	Input *types.GetGatewayResponseInput
	Copy  func(*types.GetGatewayResponseInput) GetGatewayResponseRequest
}

// Send marshals and sends the GetGatewayResponse API request.
func (r GetGatewayResponseRequest) Send(ctx context.Context) (*GetGatewayResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGatewayResponseResponse{
		GetGatewayResponseOutput: r.Request.Data.(*types.GetGatewayResponseOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGatewayResponseResponse is the response type for the
// GetGatewayResponse API operation.
type GetGatewayResponseResponse struct {
	*types.GetGatewayResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGatewayResponse request.
func (r *GetGatewayResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
