// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
)

const opDescribeStacks = "DescribeStacks"

// DescribeStacksRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Retrieves a list that describes one or more specified stacks, if the stack
// names are provided. Otherwise, all stacks in the account are described.
//
//    // Example sending a request using DescribeStacksRequest.
//    req := client.DescribeStacksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeStacks
func (c *Client) DescribeStacksRequest(input *types.DescribeStacksInput) DescribeStacksRequest {
	op := &aws.Operation{
		Name:       opDescribeStacks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeStacksInput{}
	}

	req := c.newRequest(op, input, &types.DescribeStacksOutput{})
	return DescribeStacksRequest{Request: req, Input: input, Copy: c.DescribeStacksRequest}
}

// DescribeStacksRequest is the request type for the
// DescribeStacks API operation.
type DescribeStacksRequest struct {
	*aws.Request
	Input *types.DescribeStacksInput
	Copy  func(*types.DescribeStacksInput) DescribeStacksRequest
}

// Send marshals and sends the DescribeStacks API request.
func (r DescribeStacksRequest) Send(ctx context.Context) (*DescribeStacksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStacksResponse{
		DescribeStacksOutput: r.Request.Data.(*types.DescribeStacksOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStacksResponse is the response type for the
// DescribeStacks API operation.
type DescribeStacksResponse struct {
	*types.DescribeStacksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStacks request.
func (r *DescribeStacksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
