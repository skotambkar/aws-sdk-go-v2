// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecs/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opDeleteTaskSet = "DeleteTaskSet"

// DeleteTaskSetRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Deletes a specified task set within a service. This is used when a service
// uses the EXTERNAL deployment controller type. For more information, see Amazon
// ECS Deployment Types (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using DeleteTaskSetRequest.
//    req := client.DeleteTaskSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DeleteTaskSet
func (c *Client) DeleteTaskSetRequest(input *types.DeleteTaskSetInput) DeleteTaskSetRequest {
	op := &aws.Operation{
		Name:       opDeleteTaskSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTaskSetInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTaskSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteTaskSetMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteTaskSetRequest{Request: req, Input: input, Copy: c.DeleteTaskSetRequest}
}

// DeleteTaskSetRequest is the request type for the
// DeleteTaskSet API operation.
type DeleteTaskSetRequest struct {
	*aws.Request
	Input *types.DeleteTaskSetInput
	Copy  func(*types.DeleteTaskSetInput) DeleteTaskSetRequest
}

// Send marshals and sends the DeleteTaskSet API request.
func (r DeleteTaskSetRequest) Send(ctx context.Context) (*DeleteTaskSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTaskSetResponse{
		DeleteTaskSetOutput: r.Request.Data.(*types.DeleteTaskSetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTaskSetResponse is the response type for the
// DeleteTaskSet API operation.
type DeleteTaskSetResponse struct {
	*types.DeleteTaskSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTaskSet request.
func (r *DeleteTaskSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
