// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetInstancePortStates = "GetInstancePortStates"

// GetInstancePortStatesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns the port states for a specific virtual private server, or instance.
//
//    // Example sending a request using GetInstancePortStatesRequest.
//    req := client.GetInstancePortStatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstancePortStates
func (c *Client) GetInstancePortStatesRequest(input *types.GetInstancePortStatesInput) GetInstancePortStatesRequest {
	op := &aws.Operation{
		Name:       opGetInstancePortStates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetInstancePortStatesInput{}
	}

	req := c.newRequest(op, input, &types.GetInstancePortStatesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetInstancePortStatesMarshaler{Input: input}.GetNamedBuildHandler())

	return GetInstancePortStatesRequest{Request: req, Input: input, Copy: c.GetInstancePortStatesRequest}
}

// GetInstancePortStatesRequest is the request type for the
// GetInstancePortStates API operation.
type GetInstancePortStatesRequest struct {
	*aws.Request
	Input *types.GetInstancePortStatesInput
	Copy  func(*types.GetInstancePortStatesInput) GetInstancePortStatesRequest
}

// Send marshals and sends the GetInstancePortStates API request.
func (r GetInstancePortStatesRequest) Send(ctx context.Context) (*GetInstancePortStatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstancePortStatesResponse{
		GetInstancePortStatesOutput: r.Request.Data.(*types.GetInstancePortStatesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstancePortStatesResponse is the response type for the
// GetInstancePortStates API operation.
type GetInstancePortStatesResponse struct {
	*types.GetInstancePortStatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstancePortStates request.
func (r *GetInstancePortStatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
