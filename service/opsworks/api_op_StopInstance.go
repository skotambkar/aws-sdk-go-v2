// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
)

const opStopInstance = "StopInstance"

// StopInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Stops a specified instance. When you stop a standard instance, the data disappears
// and must be reinstalled when you restart the instance. You can stop an Amazon
// EBS-backed instance without losing data. For more information, see Starting,
// Stopping, and Rebooting Instances (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-starting.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using StopInstanceRequest.
//    req := client.StopInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/StopInstance
func (c *Client) StopInstanceRequest(input *types.StopInstanceInput) StopInstanceRequest {
	op := &aws.Operation{
		Name:       opStopInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopInstanceInput{}
	}

	req := c.newRequest(op, input, &types.StopInstanceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StopInstanceRequest{Request: req, Input: input, Copy: c.StopInstanceRequest}
}

// StopInstanceRequest is the request type for the
// StopInstance API operation.
type StopInstanceRequest struct {
	*aws.Request
	Input *types.StopInstanceInput
	Copy  func(*types.StopInstanceInput) StopInstanceRequest
}

// Send marshals and sends the StopInstance API request.
func (r StopInstanceRequest) Send(ctx context.Context) (*StopInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopInstanceResponse{
		StopInstanceOutput: r.Request.Data.(*types.StopInstanceOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopInstanceResponse is the response type for the
// StopInstance API operation.
type StopInstanceResponse struct {
	*types.StopInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopInstance request.
func (r *StopInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
