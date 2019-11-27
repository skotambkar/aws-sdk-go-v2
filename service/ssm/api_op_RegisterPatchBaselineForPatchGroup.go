// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opRegisterPatchBaselineForPatchGroup = "RegisterPatchBaselineForPatchGroup"

// RegisterPatchBaselineForPatchGroupRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Registers a patch baseline for a patch group.
//
//    // Example sending a request using RegisterPatchBaselineForPatchGroupRequest.
//    req := client.RegisterPatchBaselineForPatchGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/RegisterPatchBaselineForPatchGroup
func (c *Client) RegisterPatchBaselineForPatchGroupRequest(input *types.RegisterPatchBaselineForPatchGroupInput) RegisterPatchBaselineForPatchGroupRequest {
	op := &aws.Operation{
		Name:       opRegisterPatchBaselineForPatchGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterPatchBaselineForPatchGroupInput{}
	}

	req := c.newRequest(op, input, &types.RegisterPatchBaselineForPatchGroupOutput{})
	return RegisterPatchBaselineForPatchGroupRequest{Request: req, Input: input, Copy: c.RegisterPatchBaselineForPatchGroupRequest}
}

// RegisterPatchBaselineForPatchGroupRequest is the request type for the
// RegisterPatchBaselineForPatchGroup API operation.
type RegisterPatchBaselineForPatchGroupRequest struct {
	*aws.Request
	Input *types.RegisterPatchBaselineForPatchGroupInput
	Copy  func(*types.RegisterPatchBaselineForPatchGroupInput) RegisterPatchBaselineForPatchGroupRequest
}

// Send marshals and sends the RegisterPatchBaselineForPatchGroup API request.
func (r RegisterPatchBaselineForPatchGroupRequest) Send(ctx context.Context) (*RegisterPatchBaselineForPatchGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterPatchBaselineForPatchGroupResponse{
		RegisterPatchBaselineForPatchGroupOutput: r.Request.Data.(*types.RegisterPatchBaselineForPatchGroupOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterPatchBaselineForPatchGroupResponse is the response type for the
// RegisterPatchBaselineForPatchGroup API operation.
type RegisterPatchBaselineForPatchGroupResponse struct {
	*types.RegisterPatchBaselineForPatchGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterPatchBaselineForPatchGroup request.
func (r *RegisterPatchBaselineForPatchGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
