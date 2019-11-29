// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
)

const opUpdateEndpointGroup = "UpdateEndpointGroup"

// UpdateEndpointGroupRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Update an endpoint group. To see an AWS CLI example of updating an endpoint
// group, scroll down to Example.
//
//    // Example sending a request using UpdateEndpointGroupRequest.
//    req := client.UpdateEndpointGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/UpdateEndpointGroup
func (c *Client) UpdateEndpointGroupRequest(input *types.UpdateEndpointGroupInput) UpdateEndpointGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateEndpointGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateEndpointGroupInput{}
	}

	req := c.newRequest(op, input, &types.UpdateEndpointGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateEndpointGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateEndpointGroupRequest{Request: req, Input: input, Copy: c.UpdateEndpointGroupRequest}
}

// UpdateEndpointGroupRequest is the request type for the
// UpdateEndpointGroup API operation.
type UpdateEndpointGroupRequest struct {
	*aws.Request
	Input *types.UpdateEndpointGroupInput
	Copy  func(*types.UpdateEndpointGroupInput) UpdateEndpointGroupRequest
}

// Send marshals and sends the UpdateEndpointGroup API request.
func (r UpdateEndpointGroupRequest) Send(ctx context.Context) (*UpdateEndpointGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEndpointGroupResponse{
		UpdateEndpointGroupOutput: r.Request.Data.(*types.UpdateEndpointGroupOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEndpointGroupResponse is the response type for the
// UpdateEndpointGroup API operation.
type UpdateEndpointGroupResponse struct {
	*types.UpdateEndpointGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEndpointGroup request.
func (r *UpdateEndpointGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
