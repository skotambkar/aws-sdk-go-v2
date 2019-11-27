// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloud9

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
)

const opDescribeEnvironmentStatus = "DescribeEnvironmentStatus"

// DescribeEnvironmentStatusRequest returns a request value for making API operation for
// AWS Cloud9.
//
// Gets status information for an AWS Cloud9 development environment.
//
//    // Example sending a request using DescribeEnvironmentStatusRequest.
//    req := client.DescribeEnvironmentStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloud9-2017-09-23/DescribeEnvironmentStatus
func (c *Client) DescribeEnvironmentStatusRequest(input *types.DescribeEnvironmentStatusInput) DescribeEnvironmentStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeEnvironmentStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEnvironmentStatusInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEnvironmentStatusOutput{})
	return DescribeEnvironmentStatusRequest{Request: req, Input: input, Copy: c.DescribeEnvironmentStatusRequest}
}

// DescribeEnvironmentStatusRequest is the request type for the
// DescribeEnvironmentStatus API operation.
type DescribeEnvironmentStatusRequest struct {
	*aws.Request
	Input *types.DescribeEnvironmentStatusInput
	Copy  func(*types.DescribeEnvironmentStatusInput) DescribeEnvironmentStatusRequest
}

// Send marshals and sends the DescribeEnvironmentStatus API request.
func (r DescribeEnvironmentStatusRequest) Send(ctx context.Context) (*DescribeEnvironmentStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEnvironmentStatusResponse{
		DescribeEnvironmentStatusOutput: r.Request.Data.(*types.DescribeEnvironmentStatusOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEnvironmentStatusResponse is the response type for the
// DescribeEnvironmentStatus API operation.
type DescribeEnvironmentStatusResponse struct {
	*types.DescribeEnvironmentStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEnvironmentStatus request.
func (r *DescribeEnvironmentStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
