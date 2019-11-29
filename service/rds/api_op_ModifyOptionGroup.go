// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opModifyOptionGroup = "ModifyOptionGroup"

// ModifyOptionGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Modifies an existing option group.
//
//    // Example sending a request using ModifyOptionGroupRequest.
//    req := client.ModifyOptionGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyOptionGroup
func (c *Client) ModifyOptionGroupRequest(input *types.ModifyOptionGroupInput) ModifyOptionGroupRequest {
	op := &aws.Operation{
		Name:       opModifyOptionGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyOptionGroupInput{}
	}

	req := c.newRequest(op, input, &types.ModifyOptionGroupOutput{})
	return ModifyOptionGroupRequest{Request: req, Input: input, Copy: c.ModifyOptionGroupRequest}
}

// ModifyOptionGroupRequest is the request type for the
// ModifyOptionGroup API operation.
type ModifyOptionGroupRequest struct {
	*aws.Request
	Input *types.ModifyOptionGroupInput
	Copy  func(*types.ModifyOptionGroupInput) ModifyOptionGroupRequest
}

// Send marshals and sends the ModifyOptionGroup API request.
func (r ModifyOptionGroupRequest) Send(ctx context.Context) (*ModifyOptionGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyOptionGroupResponse{
		ModifyOptionGroupOutput: r.Request.Data.(*types.ModifyOptionGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyOptionGroupResponse is the response type for the
// ModifyOptionGroup API operation.
type ModifyOptionGroupResponse struct {
	*types.ModifyOptionGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyOptionGroup request.
func (r *ModifyOptionGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
