// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opUpdateThingGroup = "UpdateThingGroup"

// UpdateThingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Update a thing group.
//
//    // Example sending a request using UpdateThingGroupRequest.
//    req := client.UpdateThingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateThingGroupRequest(input *types.UpdateThingGroupInput) UpdateThingGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateThingGroup,
		HTTPMethod: "PATCH",
		HTTPPath:   "/thing-groups/{thingGroupName}",
	}

	if input == nil {
		input = &types.UpdateThingGroupInput{}
	}

	req := c.newRequest(op, input, &types.UpdateThingGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateThingGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateThingGroupRequest{Request: req, Input: input, Copy: c.UpdateThingGroupRequest}
}

// UpdateThingGroupRequest is the request type for the
// UpdateThingGroup API operation.
type UpdateThingGroupRequest struct {
	*aws.Request
	Input *types.UpdateThingGroupInput
	Copy  func(*types.UpdateThingGroupInput) UpdateThingGroupRequest
}

// Send marshals and sends the UpdateThingGroup API request.
func (r UpdateThingGroupRequest) Send(ctx context.Context) (*UpdateThingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateThingGroupResponse{
		UpdateThingGroupOutput: r.Request.Data.(*types.UpdateThingGroupOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateThingGroupResponse is the response type for the
// UpdateThingGroup API operation.
type UpdateThingGroupResponse struct {
	*types.UpdateThingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateThingGroup request.
func (r *UpdateThingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
