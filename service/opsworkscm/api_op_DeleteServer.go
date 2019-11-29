// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
)

const opDeleteServer = "DeleteServer"

// DeleteServerRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Deletes the server and the underlying AWS CloudFormation stacks (including
// the server's EC2 instance). When you run this command, the server state is
// updated to DELETING. After the server is deleted, it is no longer returned
// by DescribeServer requests. If the AWS CloudFormation stack cannot be deleted,
// the server cannot be deleted.
//
// This operation is asynchronous.
//
// An InvalidStateException is thrown when a server deletion is already in progress.
// A ResourceNotFoundException is thrown when the server does not exist. A ValidationException
// is raised when parameters of the request are not valid.
//
//    // Example sending a request using DeleteServerRequest.
//    req := client.DeleteServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/DeleteServer
func (c *Client) DeleteServerRequest(input *types.DeleteServerInput) DeleteServerRequest {
	op := &aws.Operation{
		Name:       opDeleteServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteServerInput{}
	}

	req := c.newRequest(op, input, &types.DeleteServerOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteServerMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteServerRequest{Request: req, Input: input, Copy: c.DeleteServerRequest}
}

// DeleteServerRequest is the request type for the
// DeleteServer API operation.
type DeleteServerRequest struct {
	*aws.Request
	Input *types.DeleteServerInput
	Copy  func(*types.DeleteServerInput) DeleteServerRequest
}

// Send marshals and sends the DeleteServer API request.
func (r DeleteServerRequest) Send(ctx context.Context) (*DeleteServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteServerResponse{
		DeleteServerOutput: r.Request.Data.(*types.DeleteServerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteServerResponse is the response type for the
// DeleteServer API operation.
type DeleteServerResponse struct {
	*types.DeleteServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteServer request.
func (r *DeleteServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
