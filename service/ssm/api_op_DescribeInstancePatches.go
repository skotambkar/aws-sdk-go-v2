// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opDescribeInstancePatches = "DescribeInstancePatches"

// DescribeInstancePatchesRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves information about the patches on the specified instance and their
// state relative to the patch baseline being used for the instance.
//
//    // Example sending a request using DescribeInstancePatchesRequest.
//    req := client.DescribeInstancePatchesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeInstancePatches
func (c *Client) DescribeInstancePatchesRequest(input *types.DescribeInstancePatchesInput) DescribeInstancePatchesRequest {
	op := &aws.Operation{
		Name:       opDescribeInstancePatches,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeInstancePatchesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeInstancePatchesOutput{})
	return DescribeInstancePatchesRequest{Request: req, Input: input, Copy: c.DescribeInstancePatchesRequest}
}

// DescribeInstancePatchesRequest is the request type for the
// DescribeInstancePatches API operation.
type DescribeInstancePatchesRequest struct {
	*aws.Request
	Input *types.DescribeInstancePatchesInput
	Copy  func(*types.DescribeInstancePatchesInput) DescribeInstancePatchesRequest
}

// Send marshals and sends the DescribeInstancePatches API request.
func (r DescribeInstancePatchesRequest) Send(ctx context.Context) (*DescribeInstancePatchesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstancePatchesResponse{
		DescribeInstancePatchesOutput: r.Request.Data.(*types.DescribeInstancePatchesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstancePatchesResponse is the response type for the
// DescribeInstancePatches API operation.
type DescribeInstancePatchesResponse struct {
	*types.DescribeInstancePatchesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstancePatches request.
func (r *DescribeInstancePatchesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
