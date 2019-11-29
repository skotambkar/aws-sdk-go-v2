// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opDescribeType = "DescribeType"

// DescribeTypeRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns detailed information about a type that has been registered.
//
// If you specify a VersionId, DescribeType returns information about that specific
// type version. Otherwise, it returns information about the default type version.
//
//    // Example sending a request using DescribeTypeRequest.
//    req := client.DescribeTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DescribeType
func (c *Client) DescribeTypeRequest(input *types.DescribeTypeInput) DescribeTypeRequest {
	op := &aws.Operation{
		Name:       opDescribeType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeTypeInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTypeOutput{})
	return DescribeTypeRequest{Request: req, Input: input, Copy: c.DescribeTypeRequest}
}

// DescribeTypeRequest is the request type for the
// DescribeType API operation.
type DescribeTypeRequest struct {
	*aws.Request
	Input *types.DescribeTypeInput
	Copy  func(*types.DescribeTypeInput) DescribeTypeRequest
}

// Send marshals and sends the DescribeType API request.
func (r DescribeTypeRequest) Send(ctx context.Context) (*DescribeTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTypeResponse{
		DescribeTypeOutput: r.Request.Data.(*types.DescribeTypeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTypeResponse is the response type for the
// DescribeType API operation.
type DescribeTypeResponse struct {
	*types.DescribeTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeType request.
func (r *DescribeTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
