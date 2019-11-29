// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
)

const opRebuildWorkspaces = "RebuildWorkspaces"

// RebuildWorkspacesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Rebuilds the specified WorkSpace.
//
// You cannot rebuild a WorkSpace unless its state is AVAILABLE, ERROR, or UNHEALTHY.
//
// Rebuilding a WorkSpace is a potentially destructive action that can result
// in the loss of data. For more information, see Rebuild a WorkSpace (https://docs.aws.amazon.com/workspaces/latest/adminguide/reset-workspace.html).
//
// This operation is asynchronous and returns before the WorkSpaces have been
// completely rebuilt.
//
//    // Example sending a request using RebuildWorkspacesRequest.
//    req := client.RebuildWorkspacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/RebuildWorkspaces
func (c *Client) RebuildWorkspacesRequest(input *types.RebuildWorkspacesInput) RebuildWorkspacesRequest {
	op := &aws.Operation{
		Name:       opRebuildWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RebuildWorkspacesInput{}
	}

	req := c.newRequest(op, input, &types.RebuildWorkspacesOutput{})
	return RebuildWorkspacesRequest{Request: req, Input: input, Copy: c.RebuildWorkspacesRequest}
}

// RebuildWorkspacesRequest is the request type for the
// RebuildWorkspaces API operation.
type RebuildWorkspacesRequest struct {
	*aws.Request
	Input *types.RebuildWorkspacesInput
	Copy  func(*types.RebuildWorkspacesInput) RebuildWorkspacesRequest
}

// Send marshals and sends the RebuildWorkspaces API request.
func (r RebuildWorkspacesRequest) Send(ctx context.Context) (*RebuildWorkspacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebuildWorkspacesResponse{
		RebuildWorkspacesOutput: r.Request.Data.(*types.RebuildWorkspacesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebuildWorkspacesResponse is the response type for the
// RebuildWorkspaces API operation.
type RebuildWorkspacesResponse struct {
	*types.RebuildWorkspacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebuildWorkspaces request.
func (r *RebuildWorkspacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
