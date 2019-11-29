// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
)

const opStartInstance = "StartInstance"

// StartInstanceRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Starts a specified instance. For more information, see Starting, Stopping,
// and Rebooting Instances (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-starting.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using StartInstanceRequest.
//    req := client.StartInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/StartInstance
func (c *Client) StartInstanceRequest(input *types.StartInstanceInput) StartInstanceRequest {
	op := &aws.Operation{
		Name:       opStartInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartInstanceInput{}
	}

	req := c.newRequest(op, input, &types.StartInstanceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StartInstanceMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return StartInstanceRequest{Request: req, Input: input, Copy: c.StartInstanceRequest}
}

// StartInstanceRequest is the request type for the
// StartInstance API operation.
type StartInstanceRequest struct {
	*aws.Request
	Input *types.StartInstanceInput
	Copy  func(*types.StartInstanceInput) StartInstanceRequest
}

// Send marshals and sends the StartInstance API request.
func (r StartInstanceRequest) Send(ctx context.Context) (*StartInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartInstanceResponse{
		StartInstanceOutput: r.Request.Data.(*types.StartInstanceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartInstanceResponse is the response type for the
// StartInstance API operation.
type StartInstanceResponse struct {
	*types.StartInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartInstance request.
func (r *StartInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
