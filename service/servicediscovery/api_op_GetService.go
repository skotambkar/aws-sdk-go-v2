// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
)

const opGetService = "GetService"

// GetServiceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Gets the settings for a specified service.
//
//    // Example sending a request using GetServiceRequest.
//    req := client.GetServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/GetService
func (c *Client) GetServiceRequest(input *types.GetServiceInput) GetServiceRequest {
	op := &aws.Operation{
		Name:       opGetService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetServiceInput{}
	}

	req := c.newRequest(op, input, &types.GetServiceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetServiceMarshaler{Input: input}.GetNamedBuildHandler())

	return GetServiceRequest{Request: req, Input: input, Copy: c.GetServiceRequest}
}

// GetServiceRequest is the request type for the
// GetService API operation.
type GetServiceRequest struct {
	*aws.Request
	Input *types.GetServiceInput
	Copy  func(*types.GetServiceInput) GetServiceRequest
}

// Send marshals and sends the GetService API request.
func (r GetServiceRequest) Send(ctx context.Context) (*GetServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetServiceResponse{
		GetServiceOutput: r.Request.Data.(*types.GetServiceOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetServiceResponse is the response type for the
// GetService API operation.
type GetServiceResponse struct {
	*types.GetServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetService request.
func (r *GetServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
