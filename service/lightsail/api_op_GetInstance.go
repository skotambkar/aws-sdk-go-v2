// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetInstance = "GetInstance"

// GetInstanceRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about a specific Amazon Lightsail instance, which is
// a virtual private server.
//
//    // Example sending a request using GetInstanceRequest.
//    req := client.GetInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstance
func (c *Client) GetInstanceRequest(input *types.GetInstanceInput) GetInstanceRequest {
	op := &aws.Operation{
		Name:       opGetInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetInstanceInput{}
	}

	req := c.newRequest(op, input, &types.GetInstanceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetInstanceMarshaler{Input: input}.GetNamedBuildHandler())

	return GetInstanceRequest{Request: req, Input: input, Copy: c.GetInstanceRequest}
}

// GetInstanceRequest is the request type for the
// GetInstance API operation.
type GetInstanceRequest struct {
	*aws.Request
	Input *types.GetInstanceInput
	Copy  func(*types.GetInstanceInput) GetInstanceRequest
}

// Send marshals and sends the GetInstance API request.
func (r GetInstanceRequest) Send(ctx context.Context) (*GetInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstanceResponse{
		GetInstanceOutput: r.Request.Data.(*types.GetInstanceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstanceResponse is the response type for the
// GetInstance API operation.
type GetInstanceResponse struct {
	*types.GetInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstance request.
func (r *GetInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
