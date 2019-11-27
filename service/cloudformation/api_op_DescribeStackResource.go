// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opDescribeStackResource = "DescribeStackResource"

// DescribeStackResourceRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns a description of the specified resource in the specified stack.
//
// For deleted stacks, DescribeStackResource returns resource information for
// up to 90 days after the stack has been deleted.
//
//    // Example sending a request using DescribeStackResourceRequest.
//    req := client.DescribeStackResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeStackResource
func (c *Client) DescribeStackResourceRequest(input *types.DescribeStackResourceInput) DescribeStackResourceRequest {
	op := &aws.Operation{
		Name:       opDescribeStackResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeStackResourceInput{}
	}

	req := c.newRequest(op, input, &types.DescribeStackResourceOutput{})
	return DescribeStackResourceRequest{Request: req, Input: input, Copy: c.DescribeStackResourceRequest}
}

// DescribeStackResourceRequest is the request type for the
// DescribeStackResource API operation.
type DescribeStackResourceRequest struct {
	*aws.Request
	Input *types.DescribeStackResourceInput
	Copy  func(*types.DescribeStackResourceInput) DescribeStackResourceRequest
}

// Send marshals and sends the DescribeStackResource API request.
func (r DescribeStackResourceRequest) Send(ctx context.Context) (*DescribeStackResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackResourceResponse{
		DescribeStackResourceOutput: r.Request.Data.(*types.DescribeStackResourceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackResourceResponse is the response type for the
// DescribeStackResource API operation.
type DescribeStackResourceResponse struct {
	*types.DescribeStackResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackResource request.
func (r *DescribeStackResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
