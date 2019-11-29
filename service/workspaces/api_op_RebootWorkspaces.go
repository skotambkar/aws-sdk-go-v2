// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
)

const opRebootWorkspaces = "RebootWorkspaces"

// RebootWorkspacesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Reboots the specified WorkSpaces.
//
// You cannot reboot a WorkSpace unless its state is AVAILABLE or UNHEALTHY.
//
// This operation is asynchronous and returns before the WorkSpaces have rebooted.
//
//    // Example sending a request using RebootWorkspacesRequest.
//    req := client.RebootWorkspacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/RebootWorkspaces
func (c *Client) RebootWorkspacesRequest(input *types.RebootWorkspacesInput) RebootWorkspacesRequest {
	op := &aws.Operation{
		Name:       opRebootWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RebootWorkspacesInput{}
	}

	req := c.newRequest(op, input, &types.RebootWorkspacesOutput{})
	return RebootWorkspacesRequest{Request: req, Input: input, Copy: c.RebootWorkspacesRequest}
}

// RebootWorkspacesRequest is the request type for the
// RebootWorkspaces API operation.
type RebootWorkspacesRequest struct {
	*aws.Request
	Input *types.RebootWorkspacesInput
	Copy  func(*types.RebootWorkspacesInput) RebootWorkspacesRequest
}

// Send marshals and sends the RebootWorkspaces API request.
func (r RebootWorkspacesRequest) Send(ctx context.Context) (*RebootWorkspacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebootWorkspacesResponse{
		RebootWorkspacesOutput: r.Request.Data.(*types.RebootWorkspacesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebootWorkspacesResponse is the response type for the
// RebootWorkspaces API operation.
type RebootWorkspacesResponse struct {
	*types.RebootWorkspacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebootWorkspaces request.
func (r *RebootWorkspacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
