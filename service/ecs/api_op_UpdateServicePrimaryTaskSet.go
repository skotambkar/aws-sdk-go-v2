// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecs/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opUpdateServicePrimaryTaskSet = "UpdateServicePrimaryTaskSet"

// UpdateServicePrimaryTaskSetRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Modifies which task set in a service is the primary task set. Any parameters
// that are updated on the primary task set in a service will transition to
// the service. This is used when a service uses the EXTERNAL deployment controller
// type. For more information, see Amazon ECS Deployment Types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using UpdateServicePrimaryTaskSetRequest.
//    req := client.UpdateServicePrimaryTaskSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/UpdateServicePrimaryTaskSet
func (c *Client) UpdateServicePrimaryTaskSetRequest(input *types.UpdateServicePrimaryTaskSetInput) UpdateServicePrimaryTaskSetRequest {
	op := &aws.Operation{
		Name:       opUpdateServicePrimaryTaskSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateServicePrimaryTaskSetInput{}
	}

	req := c.newRequest(op, input, &types.UpdateServicePrimaryTaskSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateServicePrimaryTaskSetMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateServicePrimaryTaskSetRequest{Request: req, Input: input, Copy: c.UpdateServicePrimaryTaskSetRequest}
}

// UpdateServicePrimaryTaskSetRequest is the request type for the
// UpdateServicePrimaryTaskSet API operation.
type UpdateServicePrimaryTaskSetRequest struct {
	*aws.Request
	Input *types.UpdateServicePrimaryTaskSetInput
	Copy  func(*types.UpdateServicePrimaryTaskSetInput) UpdateServicePrimaryTaskSetRequest
}

// Send marshals and sends the UpdateServicePrimaryTaskSet API request.
func (r UpdateServicePrimaryTaskSetRequest) Send(ctx context.Context) (*UpdateServicePrimaryTaskSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServicePrimaryTaskSetResponse{
		UpdateServicePrimaryTaskSetOutput: r.Request.Data.(*types.UpdateServicePrimaryTaskSetOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServicePrimaryTaskSetResponse is the response type for the
// UpdateServicePrimaryTaskSet API operation.
type UpdateServicePrimaryTaskSetResponse struct {
	*types.UpdateServicePrimaryTaskSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateServicePrimaryTaskSet request.
func (r *UpdateServicePrimaryTaskSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
