// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetBasePathMapping = "GetBasePathMapping"

// GetBasePathMappingRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Describe a BasePathMapping resource.
//
//    // Example sending a request using GetBasePathMappingRequest.
//    req := client.GetBasePathMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetBasePathMappingRequest(input *types.GetBasePathMappingInput) GetBasePathMappingRequest {
	op := &aws.Operation{
		Name:       opGetBasePathMapping,
		HTTPMethod: "GET",
		HTTPPath:   "/domainnames/{domain_name}/basepathmappings/{base_path}",
	}

	if input == nil {
		input = &types.GetBasePathMappingInput{}
	}

	req := c.newRequest(op, input, &types.GetBasePathMappingOutput{})
	return GetBasePathMappingRequest{Request: req, Input: input, Copy: c.GetBasePathMappingRequest}
}

// GetBasePathMappingRequest is the request type for the
// GetBasePathMapping API operation.
type GetBasePathMappingRequest struct {
	*aws.Request
	Input *types.GetBasePathMappingInput
	Copy  func(*types.GetBasePathMappingInput) GetBasePathMappingRequest
}

// Send marshals and sends the GetBasePathMapping API request.
func (r GetBasePathMappingRequest) Send(ctx context.Context) (*GetBasePathMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBasePathMappingResponse{
		GetBasePathMappingOutput: r.Request.Data.(*types.GetBasePathMappingOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBasePathMappingResponse is the response type for the
// GetBasePathMapping API operation.
type GetBasePathMappingResponse struct {
	*types.GetBasePathMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBasePathMapping request.
func (r *GetBasePathMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
