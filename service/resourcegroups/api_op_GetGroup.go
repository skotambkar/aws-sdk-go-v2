// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
)

const opGetGroup = "GetGroup"

// GetGroupRequest returns a request value for making API operation for
// AWS Resource Groups.
//
// Returns information about a specified resource group.
//
//    // Example sending a request using GetGroupRequest.
//    req := client.GetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/GetGroup
func (c *Client) GetGroupRequest(input *types.GetGroupInput) GetGroupRequest {
	op := &aws.Operation{
		Name:       opGetGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/groups/{GroupName}",
	}

	if input == nil {
		input = &types.GetGroupInput{}
	}

	req := c.newRequest(op, input, &types.GetGroupOutput{})
	return GetGroupRequest{Request: req, Input: input, Copy: c.GetGroupRequest}
}

// GetGroupRequest is the request type for the
// GetGroup API operation.
type GetGroupRequest struct {
	*aws.Request
	Input *types.GetGroupInput
	Copy  func(*types.GetGroupInput) GetGroupRequest
}

// Send marshals and sends the GetGroup API request.
func (r GetGroupRequest) Send(ctx context.Context) (*GetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGroupResponse{
		GetGroupOutput: r.Request.Data.(*types.GetGroupOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGroupResponse is the response type for the
// GetGroup API operation.
type GetGroupResponse struct {
	*types.GetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGroup request.
func (r *GetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
