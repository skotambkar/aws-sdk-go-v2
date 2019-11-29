// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opAllocateConnectionOnInterconnect = "AllocateConnectionOnInterconnect"

// AllocateConnectionOnInterconnectRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deprecated. Use AllocateHostedConnection instead.
//
// Creates a hosted connection on an interconnect.
//
// Allocates a VLAN number and a specified amount of bandwidth for use by a
// hosted connection on the specified interconnect.
//
// Intended for use by AWS Direct Connect Partners only.
//
//    // Example sending a request using AllocateConnectionOnInterconnectRequest.
//    req := client.AllocateConnectionOnInterconnectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocateConnectionOnInterconnect
func (c *Client) AllocateConnectionOnInterconnectRequest(input *types.AllocateConnectionOnInterconnectInput) AllocateConnectionOnInterconnectRequest {
	if c.Client.Config.Logger != nil {
		c.Client.Config.Logger.Log("This operation, AllocateConnectionOnInterconnect, has been deprecated")
	}
	op := &aws.Operation{
		Name:       opAllocateConnectionOnInterconnect,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AllocateConnectionOnInterconnectInput{}
	}

	req := c.newRequest(op, input, &types.AllocateConnectionOnInterconnectOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AllocateConnectionOnInterconnectMarshaler{Input: input}.GetNamedBuildHandler())

	return AllocateConnectionOnInterconnectRequest{Request: req, Input: input, Copy: c.AllocateConnectionOnInterconnectRequest}
}

// AllocateConnectionOnInterconnectRequest is the request type for the
// AllocateConnectionOnInterconnect API operation.
type AllocateConnectionOnInterconnectRequest struct {
	*aws.Request
	Input *types.AllocateConnectionOnInterconnectInput
	Copy  func(*types.AllocateConnectionOnInterconnectInput) AllocateConnectionOnInterconnectRequest
}

// Send marshals and sends the AllocateConnectionOnInterconnect API request.
func (r AllocateConnectionOnInterconnectRequest) Send(ctx context.Context) (*AllocateConnectionOnInterconnectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AllocateConnectionOnInterconnectResponse{
		AllocateConnectionOnInterconnectOutput: r.Request.Data.(*types.AllocateConnectionOnInterconnectOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AllocateConnectionOnInterconnectResponse is the response type for the
// AllocateConnectionOnInterconnect API operation.
type AllocateConnectionOnInterconnectResponse struct {
	*types.AllocateConnectionOnInterconnectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AllocateConnectionOnInterconnect request.
func (r *AllocateConnectionOnInterconnectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
