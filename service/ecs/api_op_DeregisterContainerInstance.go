// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opDeregisterContainerInstance = "DeregisterContainerInstance"

// DeregisterContainerInstanceRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Deregisters an Amazon ECS container instance from the specified cluster.
// This instance is no longer available to run tasks.
//
// If you intend to use the container instance for some other purpose after
// deregistration, you should stop all of the tasks running on the container
// instance before deregistration. That prevents any orphaned tasks from consuming
// resources.
//
// Deregistering a container instance removes the instance from a cluster, but
// it does not terminate the EC2 instance. If you are finished using the instance,
// be sure to terminate it in the Amazon EC2 console to stop billing.
//
// If you terminate a running container instance, Amazon ECS automatically deregisters
// the instance from your cluster (stopped container instances or instances
// with disconnected agents are not automatically deregistered when terminated).
//
//    // Example sending a request using DeregisterContainerInstanceRequest.
//    req := client.DeregisterContainerInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeregisterContainerInstance
func (c *Client) DeregisterContainerInstanceRequest(input *types.DeregisterContainerInstanceInput) DeregisterContainerInstanceRequest {
	op := &aws.Operation{
		Name:       opDeregisterContainerInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeregisterContainerInstanceInput{}
	}

	req := c.newRequest(op, input, &types.DeregisterContainerInstanceOutput{})
	return DeregisterContainerInstanceRequest{Request: req, Input: input, Copy: c.DeregisterContainerInstanceRequest}
}

// DeregisterContainerInstanceRequest is the request type for the
// DeregisterContainerInstance API operation.
type DeregisterContainerInstanceRequest struct {
	*aws.Request
	Input *types.DeregisterContainerInstanceInput
	Copy  func(*types.DeregisterContainerInstanceInput) DeregisterContainerInstanceRequest
}

// Send marshals and sends the DeregisterContainerInstance API request.
func (r DeregisterContainerInstanceRequest) Send(ctx context.Context) (*DeregisterContainerInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterContainerInstanceResponse{
		DeregisterContainerInstanceOutput: r.Request.Data.(*types.DeregisterContainerInstanceOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterContainerInstanceResponse is the response type for the
// DeregisterContainerInstance API operation.
type DeregisterContainerInstanceResponse struct {
	*types.DeregisterContainerInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterContainerInstance request.
func (r *DeregisterContainerInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
