// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opCopyOptionGroup = "CopyOptionGroup"

// CopyOptionGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Copies the specified option group.
//
//    // Example sending a request using CopyOptionGroupRequest.
//    req := client.CopyOptionGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CopyOptionGroup
func (c *Client) CopyOptionGroupRequest(input *types.CopyOptionGroupInput) CopyOptionGroupRequest {
	op := &aws.Operation{
		Name:       opCopyOptionGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CopyOptionGroupInput{}
	}

	req := c.newRequest(op, input, &types.CopyOptionGroupOutput{})
	return CopyOptionGroupRequest{Request: req, Input: input, Copy: c.CopyOptionGroupRequest}
}

// CopyOptionGroupRequest is the request type for the
// CopyOptionGroup API operation.
type CopyOptionGroupRequest struct {
	*aws.Request
	Input *types.CopyOptionGroupInput
	Copy  func(*types.CopyOptionGroupInput) CopyOptionGroupRequest
}

// Send marshals and sends the CopyOptionGroup API request.
func (r CopyOptionGroupRequest) Send(ctx context.Context) (*CopyOptionGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyOptionGroupResponse{
		CopyOptionGroupOutput: r.Request.Data.(*types.CopyOptionGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyOptionGroupResponse is the response type for the
// CopyOptionGroup API operation.
type CopyOptionGroupResponse struct {
	*types.CopyOptionGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyOptionGroup request.
func (r *CopyOptionGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
