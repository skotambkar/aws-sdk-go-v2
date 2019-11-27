// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opRegisterTaskDefinition = "RegisterTaskDefinition"

// RegisterTaskDefinitionRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Registers a new task definition from the supplied family and containerDefinitions.
// Optionally, you can add data volumes to your containers with the volumes
// parameter. For more information about task definition parameters and defaults,
// see Amazon ECS Task Definitions (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_defintions.html)
// in the Amazon Elastic Container Service Developer Guide.
//
// You can specify an IAM role for your task with the taskRoleArn parameter.
// When you specify an IAM role for a task, its containers can then use the
// latest versions of the AWS CLI or SDKs to make API requests to the AWS services
// that are specified in the IAM policy associated with the role. For more information,
// see IAM Roles for Tasks (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-iam-roles.html)
// in the Amazon Elastic Container Service Developer Guide.
//
// You can specify a Docker networking mode for the containers in your task
// definition with the networkMode parameter. The available network modes correspond
// to those described in Network settings (https://docs.docker.com/engine/reference/run/#/network-settings)
// in the Docker run reference. If you specify the awsvpc network mode, the
// task is allocated an elastic network interface, and you must specify a NetworkConfiguration
// when you create a service or run a task with the task definition. For more
// information, see Task Networking (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using RegisterTaskDefinitionRequest.
//    req := client.RegisterTaskDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/RegisterTaskDefinition
func (c *Client) RegisterTaskDefinitionRequest(input *types.RegisterTaskDefinitionInput) RegisterTaskDefinitionRequest {
	op := &aws.Operation{
		Name:       opRegisterTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterTaskDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.RegisterTaskDefinitionOutput{})
	return RegisterTaskDefinitionRequest{Request: req, Input: input, Copy: c.RegisterTaskDefinitionRequest}
}

// RegisterTaskDefinitionRequest is the request type for the
// RegisterTaskDefinition API operation.
type RegisterTaskDefinitionRequest struct {
	*aws.Request
	Input *types.RegisterTaskDefinitionInput
	Copy  func(*types.RegisterTaskDefinitionInput) RegisterTaskDefinitionRequest
}

// Send marshals and sends the RegisterTaskDefinition API request.
func (r RegisterTaskDefinitionRequest) Send(ctx context.Context) (*RegisterTaskDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterTaskDefinitionResponse{
		RegisterTaskDefinitionOutput: r.Request.Data.(*types.RegisterTaskDefinitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterTaskDefinitionResponse is the response type for the
// RegisterTaskDefinition API operation.
type RegisterTaskDefinitionResponse struct {
	*types.RegisterTaskDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterTaskDefinition request.
func (r *RegisterTaskDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
