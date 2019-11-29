// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetEndpoint = "GetEndpoint"

// GetEndpointRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the settings and attributes of a specific endpoint
// for an application.
//
//    // Example sending a request using GetEndpointRequest.
//    req := client.GetEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEndpoint
func (c *Client) GetEndpointRequest(input *types.GetEndpointInput) GetEndpointRequest {
	op := &aws.Operation{
		Name:       opGetEndpoint,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/endpoints/{endpoint-id}",
	}

	if input == nil {
		input = &types.GetEndpointInput{}
	}

	req := c.newRequest(op, input, &types.GetEndpointOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetEndpointMarshaler{Input: input}.GetNamedBuildHandler())

	return GetEndpointRequest{Request: req, Input: input, Copy: c.GetEndpointRequest}
}

// GetEndpointRequest is the request type for the
// GetEndpoint API operation.
type GetEndpointRequest struct {
	*aws.Request
	Input *types.GetEndpointInput
	Copy  func(*types.GetEndpointInput) GetEndpointRequest
}

// Send marshals and sends the GetEndpoint API request.
func (r GetEndpointRequest) Send(ctx context.Context) (*GetEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEndpointResponse{
		GetEndpointOutput: r.Request.Data.(*types.GetEndpointOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEndpointResponse is the response type for the
// GetEndpoint API operation.
type GetEndpointResponse struct {
	*types.GetEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEndpoint request.
func (r *GetEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
