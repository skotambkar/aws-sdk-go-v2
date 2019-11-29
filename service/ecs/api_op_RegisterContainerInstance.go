// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opRegisterContainerInstance = "RegisterContainerInstance"

// RegisterContainerInstanceRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
//
// This action is only used by the Amazon ECS agent, and it is not intended
// for use outside of the agent.
//
// Registers an EC2 instance into the specified cluster. This instance becomes
// available to place containers on.
//
//    // Example sending a request using RegisterContainerInstanceRequest.
//    req := client.RegisterContainerInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/RegisterContainerInstance
func (c *Client) RegisterContainerInstanceRequest(input *types.RegisterContainerInstanceInput) RegisterContainerInstanceRequest {
	op := &aws.Operation{
		Name:       opRegisterContainerInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterContainerInstanceInput{}
	}

	req := c.newRequest(op, input, &types.RegisterContainerInstanceOutput{})
	return RegisterContainerInstanceRequest{Request: req, Input: input, Copy: c.RegisterContainerInstanceRequest}
}

// RegisterContainerInstanceRequest is the request type for the
// RegisterContainerInstance API operation.
type RegisterContainerInstanceRequest struct {
	*aws.Request
	Input *types.RegisterContainerInstanceInput
	Copy  func(*types.RegisterContainerInstanceInput) RegisterContainerInstanceRequest
}

// Send marshals and sends the RegisterContainerInstance API request.
func (r RegisterContainerInstanceRequest) Send(ctx context.Context) (*RegisterContainerInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterContainerInstanceResponse{
		RegisterContainerInstanceOutput: r.Request.Data.(*types.RegisterContainerInstanceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterContainerInstanceResponse is the response type for the
// RegisterContainerInstance API operation.
type RegisterContainerInstanceResponse struct {
	*types.RegisterContainerInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterContainerInstance request.
func (r *RegisterContainerInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
