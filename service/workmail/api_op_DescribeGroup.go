// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
)

const opDescribeGroup = "DescribeGroup"

// DescribeGroupRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Returns the data available for the group.
//
//    // Example sending a request using DescribeGroupRequest.
//    req := client.DescribeGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DescribeGroup
func (c *Client) DescribeGroupRequest(input *types.DescribeGroupInput) DescribeGroupRequest {
	op := &aws.Operation{
		Name:       opDescribeGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeGroupInput{}
	}

	req := c.newRequest(op, input, &types.DescribeGroupOutput{})
	return DescribeGroupRequest{Request: req, Input: input, Copy: c.DescribeGroupRequest}
}

// DescribeGroupRequest is the request type for the
// DescribeGroup API operation.
type DescribeGroupRequest struct {
	*aws.Request
	Input *types.DescribeGroupInput
	Copy  func(*types.DescribeGroupInput) DescribeGroupRequest
}

// Send marshals and sends the DescribeGroup API request.
func (r DescribeGroupRequest) Send(ctx context.Context) (*DescribeGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGroupResponse{
		DescribeGroupOutput: r.Request.Data.(*types.DescribeGroupOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGroupResponse is the response type for the
// DescribeGroup API operation.
type DescribeGroupResponse struct {
	*types.DescribeGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGroup request.
func (r *DescribeGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
