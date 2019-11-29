// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
)

const opGetDataflowEndpointGroup = "GetDataflowEndpointGroup"

// GetDataflowEndpointGroupRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns the dataflow endpoint group.
//
//    // Example sending a request using GetDataflowEndpointGroupRequest.
//    req := client.GetDataflowEndpointGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/GetDataflowEndpointGroup
func (c *Client) GetDataflowEndpointGroupRequest(input *types.GetDataflowEndpointGroupInput) GetDataflowEndpointGroupRequest {
	op := &aws.Operation{
		Name:       opGetDataflowEndpointGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/dataflowEndpointGroup/{dataflowEndpointGroupId}",
	}

	if input == nil {
		input = &types.GetDataflowEndpointGroupInput{}
	}

	req := c.newRequest(op, input, &types.GetDataflowEndpointGroupOutput{})
	return GetDataflowEndpointGroupRequest{Request: req, Input: input, Copy: c.GetDataflowEndpointGroupRequest}
}

// GetDataflowEndpointGroupRequest is the request type for the
// GetDataflowEndpointGroup API operation.
type GetDataflowEndpointGroupRequest struct {
	*aws.Request
	Input *types.GetDataflowEndpointGroupInput
	Copy  func(*types.GetDataflowEndpointGroupInput) GetDataflowEndpointGroupRequest
}

// Send marshals and sends the GetDataflowEndpointGroup API request.
func (r GetDataflowEndpointGroupRequest) Send(ctx context.Context) (*GetDataflowEndpointGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDataflowEndpointGroupResponse{
		GetDataflowEndpointGroupOutput: r.Request.Data.(*types.GetDataflowEndpointGroupOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDataflowEndpointGroupResponse is the response type for the
// GetDataflowEndpointGroup API operation.
type GetDataflowEndpointGroupResponse struct {
	*types.GetDataflowEndpointGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDataflowEndpointGroup request.
func (r *GetDataflowEndpointGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
