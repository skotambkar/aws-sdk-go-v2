// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opDiscoverPollEndpoint = "DiscoverPollEndpoint"

// DiscoverPollEndpointRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
//
// This action is only used by the Amazon ECS agent, and it is not intended
// for use outside of the agent.
//
// Returns an endpoint for the Amazon ECS agent to poll for updates.
//
//    // Example sending a request using DiscoverPollEndpointRequest.
//    req := client.DiscoverPollEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DiscoverPollEndpoint
func (c *Client) DiscoverPollEndpointRequest(input *types.DiscoverPollEndpointInput) DiscoverPollEndpointRequest {
	op := &aws.Operation{
		Name:       opDiscoverPollEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DiscoverPollEndpointInput{}
	}

	req := c.newRequest(op, input, &types.DiscoverPollEndpointOutput{})
	return DiscoverPollEndpointRequest{Request: req, Input: input, Copy: c.DiscoverPollEndpointRequest}
}

// DiscoverPollEndpointRequest is the request type for the
// DiscoverPollEndpoint API operation.
type DiscoverPollEndpointRequest struct {
	*aws.Request
	Input *types.DiscoverPollEndpointInput
	Copy  func(*types.DiscoverPollEndpointInput) DiscoverPollEndpointRequest
}

// Send marshals and sends the DiscoverPollEndpoint API request.
func (r DiscoverPollEndpointRequest) Send(ctx context.Context) (*DiscoverPollEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DiscoverPollEndpointResponse{
		DiscoverPollEndpointOutput: r.Request.Data.(*types.DiscoverPollEndpointOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DiscoverPollEndpointResponse is the response type for the
// DiscoverPollEndpoint API operation.
type DiscoverPollEndpointResponse struct {
	*types.DiscoverPollEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DiscoverPollEndpoint request.
func (r *DiscoverPollEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
