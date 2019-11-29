// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetMethod = "GetMethod"

// GetMethodRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Describe an existing Method resource.
//
//    // Example sending a request using GetMethodRequest.
//    req := client.GetMethodRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetMethodRequest(input *types.GetMethodInput) GetMethodRequest {
	op := &aws.Operation{
		Name:       opGetMethod,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}",
	}

	if input == nil {
		input = &types.GetMethodInput{}
	}

	req := c.newRequest(op, input, &types.GetMethodOutput{})
	return GetMethodRequest{Request: req, Input: input, Copy: c.GetMethodRequest}
}

// GetMethodRequest is the request type for the
// GetMethod API operation.
type GetMethodRequest struct {
	*aws.Request
	Input *types.GetMethodInput
	Copy  func(*types.GetMethodInput) GetMethodRequest
}

// Send marshals and sends the GetMethod API request.
func (r GetMethodRequest) Send(ctx context.Context) (*GetMethodResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMethodResponse{
		GetMethodOutput: r.Request.Data.(*types.GetMethodOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMethodResponse is the response type for the
// GetMethod API operation.
type GetMethodResponse struct {
	*types.GetMethodOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMethod request.
func (r *GetMethodResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
