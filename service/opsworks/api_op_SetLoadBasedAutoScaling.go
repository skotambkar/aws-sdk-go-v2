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

const opSetLoadBasedAutoScaling = "SetLoadBasedAutoScaling"

// SetLoadBasedAutoScalingRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Specify the load-based auto scaling configuration for a specified layer.
// For more information, see Managing Load with Time-based and Load-based Instances
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-autoscaling.html).
//
// To use load-based auto scaling, you must create a set of load-based auto
// scaling instances. Load-based auto scaling operates only on the instances
// from that set, so you must ensure that you have created enough instances
// to handle the maximum anticipated load.
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using SetLoadBasedAutoScalingRequest.
//    req := client.SetLoadBasedAutoScalingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/SetLoadBasedAutoScaling
func (c *Client) SetLoadBasedAutoScalingRequest(input *types.SetLoadBasedAutoScalingInput) SetLoadBasedAutoScalingRequest {
	op := &aws.Operation{
		Name:       opSetLoadBasedAutoScaling,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetLoadBasedAutoScalingInput{}
	}

	req := c.newRequest(op, input, &types.SetLoadBasedAutoScalingOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.SetLoadBasedAutoScalingMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetLoadBasedAutoScalingRequest{Request: req, Input: input, Copy: c.SetLoadBasedAutoScalingRequest}
}

// SetLoadBasedAutoScalingRequest is the request type for the
// SetLoadBasedAutoScaling API operation.
type SetLoadBasedAutoScalingRequest struct {
	*aws.Request
	Input *types.SetLoadBasedAutoScalingInput
	Copy  func(*types.SetLoadBasedAutoScalingInput) SetLoadBasedAutoScalingRequest
}

// Send marshals and sends the SetLoadBasedAutoScaling API request.
func (r SetLoadBasedAutoScalingRequest) Send(ctx context.Context) (*SetLoadBasedAutoScalingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetLoadBasedAutoScalingResponse{
		SetLoadBasedAutoScalingOutput: r.Request.Data.(*types.SetLoadBasedAutoScalingOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetLoadBasedAutoScalingResponse is the response type for the
// SetLoadBasedAutoScaling API operation.
type SetLoadBasedAutoScalingResponse struct {
	*types.SetLoadBasedAutoScalingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetLoadBasedAutoScaling request.
func (r *SetLoadBasedAutoScalingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
