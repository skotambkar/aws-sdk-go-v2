// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
)

const opDescribeTaskExecution = "DescribeTaskExecution"

// DescribeTaskExecutionRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns detailed metadata about a task that is being executed.
//
//    // Example sending a request using DescribeTaskExecutionRequest.
//    req := client.DescribeTaskExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeTaskExecution
func (c *Client) DescribeTaskExecutionRequest(input *types.DescribeTaskExecutionInput) DescribeTaskExecutionRequest {
	op := &aws.Operation{
		Name:       opDescribeTaskExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeTaskExecutionInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTaskExecutionOutput{})
	return DescribeTaskExecutionRequest{Request: req, Input: input, Copy: c.DescribeTaskExecutionRequest}
}

// DescribeTaskExecutionRequest is the request type for the
// DescribeTaskExecution API operation.
type DescribeTaskExecutionRequest struct {
	*aws.Request
	Input *types.DescribeTaskExecutionInput
	Copy  func(*types.DescribeTaskExecutionInput) DescribeTaskExecutionRequest
}

// Send marshals and sends the DescribeTaskExecution API request.
func (r DescribeTaskExecutionRequest) Send(ctx context.Context) (*DescribeTaskExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTaskExecutionResponse{
		DescribeTaskExecutionOutput: r.Request.Data.(*types.DescribeTaskExecutionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTaskExecutionResponse is the response type for the
// DescribeTaskExecution API operation.
type DescribeTaskExecutionResponse struct {
	*types.DescribeTaskExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTaskExecution request.
func (r *DescribeTaskExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
