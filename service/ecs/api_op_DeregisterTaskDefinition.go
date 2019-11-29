// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecs/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opDeregisterTaskDefinition = "DeregisterTaskDefinition"

// DeregisterTaskDefinitionRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Deregisters the specified task definition by family and revision. Upon deregistration,
// the task definition is marked as INACTIVE. Existing tasks and services that
// reference an INACTIVE task definition continue to run without disruption.
// Existing services that reference an INACTIVE task definition can still scale
// up or down by modifying the service's desired count.
//
// You cannot use an INACTIVE task definition to run new tasks or create new
// services, and you cannot update an existing service to reference an INACTIVE
// task definition. However, there may be up to a 10-minute window following
// deregistration where these restrictions have not yet taken effect.
//
// At this time, INACTIVE task definitions remain discoverable in your account
// indefinitely. However, this behavior is subject to change in the future,
// so you should not rely on INACTIVE task definitions persisting beyond the
// lifecycle of any associated tasks and services.
//
//    // Example sending a request using DeregisterTaskDefinitionRequest.
//    req := client.DeregisterTaskDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeregisterTaskDefinition
func (c *Client) DeregisterTaskDefinitionRequest(input *types.DeregisterTaskDefinitionInput) DeregisterTaskDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeregisterTaskDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeregisterTaskDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.DeregisterTaskDefinitionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeregisterTaskDefinitionMarshaler{Input: input}.GetNamedBuildHandler())

	return DeregisterTaskDefinitionRequest{Request: req, Input: input, Copy: c.DeregisterTaskDefinitionRequest}
}

// DeregisterTaskDefinitionRequest is the request type for the
// DeregisterTaskDefinition API operation.
type DeregisterTaskDefinitionRequest struct {
	*aws.Request
	Input *types.DeregisterTaskDefinitionInput
	Copy  func(*types.DeregisterTaskDefinitionInput) DeregisterTaskDefinitionRequest
}

// Send marshals and sends the DeregisterTaskDefinition API request.
func (r DeregisterTaskDefinitionRequest) Send(ctx context.Context) (*DeregisterTaskDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterTaskDefinitionResponse{
		DeregisterTaskDefinitionOutput: r.Request.Data.(*types.DeregisterTaskDefinitionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterTaskDefinitionResponse is the response type for the
// DeregisterTaskDefinition API operation.
type DeregisterTaskDefinitionResponse struct {
	*types.DeregisterTaskDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterTaskDefinition request.
func (r *DeregisterTaskDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
