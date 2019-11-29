// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetUserEndpoints = "GetUserEndpoints"

// GetUserEndpointsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about all the endpoints that are associated with a
// specific user ID.
//
//    // Example sending a request using GetUserEndpointsRequest.
//    req := client.GetUserEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetUserEndpoints
func (c *Client) GetUserEndpointsRequest(input *types.GetUserEndpointsInput) GetUserEndpointsRequest {
	op := &aws.Operation{
		Name:       opGetUserEndpoints,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/users/{user-id}",
	}

	if input == nil {
		input = &types.GetUserEndpointsInput{}
	}

	req := c.newRequest(op, input, &types.GetUserEndpointsOutput{})
	return GetUserEndpointsRequest{Request: req, Input: input, Copy: c.GetUserEndpointsRequest}
}

// GetUserEndpointsRequest is the request type for the
// GetUserEndpoints API operation.
type GetUserEndpointsRequest struct {
	*aws.Request
	Input *types.GetUserEndpointsInput
	Copy  func(*types.GetUserEndpointsInput) GetUserEndpointsRequest
}

// Send marshals and sends the GetUserEndpoints API request.
func (r GetUserEndpointsRequest) Send(ctx context.Context) (*GetUserEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUserEndpointsResponse{
		GetUserEndpointsOutput: r.Request.Data.(*types.GetUserEndpointsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUserEndpointsResponse is the response type for the
// GetUserEndpoints API operation.
type GetUserEndpointsResponse struct {
	*types.GetUserEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUserEndpoints request.
func (r *GetUserEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
