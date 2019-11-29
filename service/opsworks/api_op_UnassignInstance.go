// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
)

const opUnassignInstance = "UnassignInstance"

// UnassignInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Unassigns a registered instance from all layers that are using the instance.
// The instance remains in the stack as an unassigned instance, and can be assigned
// to another layer as needed. You cannot use this action with instances that
// were created with AWS OpsWorks Stacks.
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack or an attached policy that explicitly grants
// permissions. For more information about user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using UnassignInstanceRequest.
//    req := client.UnassignInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UnassignInstance
func (c *Client) UnassignInstanceRequest(input *types.UnassignInstanceInput) UnassignInstanceRequest {
	op := &aws.Operation{
		Name:       opUnassignInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UnassignInstanceInput{}
	}

	req := c.newRequest(op, input, &types.UnassignInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UnassignInstanceRequest{Request: req, Input: input, Copy: c.UnassignInstanceRequest}
}

// UnassignInstanceRequest is the request type for the
// UnassignInstance API operation.
type UnassignInstanceRequest struct {
	*aws.Request
	Input *types.UnassignInstanceInput
	Copy  func(*types.UnassignInstanceInput) UnassignInstanceRequest
}

// Send marshals and sends the UnassignInstance API request.
func (r UnassignInstanceRequest) Send(ctx context.Context) (*UnassignInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UnassignInstanceResponse{
		UnassignInstanceOutput: r.Request.Data.(*types.UnassignInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UnassignInstanceResponse is the response type for the
// UnassignInstance API operation.
type UnassignInstanceResponse struct {
	*types.UnassignInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UnassignInstance request.
func (r *UnassignInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
