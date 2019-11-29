// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opDescribeJobExecution = "DescribeJobExecution"

// DescribeJobExecutionRequest returns a request value for making API operation for
// AWS IoT.
//
// Describes a job execution.
//
//    // Example sending a request using DescribeJobExecutionRequest.
//    req := client.DescribeJobExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeJobExecutionRequest(input *types.DescribeJobExecutionInput) DescribeJobExecutionRequest {
	op := &aws.Operation{
		Name:       opDescribeJobExecution,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}/jobs/{jobId}",
	}

	if input == nil {
		input = &types.DescribeJobExecutionInput{}
	}

	req := c.newRequest(op, input, &types.DescribeJobExecutionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeJobExecutionMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeJobExecutionRequest{Request: req, Input: input, Copy: c.DescribeJobExecutionRequest}
}

// DescribeJobExecutionRequest is the request type for the
// DescribeJobExecution API operation.
type DescribeJobExecutionRequest struct {
	*aws.Request
	Input *types.DescribeJobExecutionInput
	Copy  func(*types.DescribeJobExecutionInput) DescribeJobExecutionRequest
}

// Send marshals and sends the DescribeJobExecution API request.
func (r DescribeJobExecutionRequest) Send(ctx context.Context) (*DescribeJobExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJobExecutionResponse{
		DescribeJobExecutionOutput: r.Request.Data.(*types.DescribeJobExecutionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeJobExecutionResponse is the response type for the
// DescribeJobExecution API operation.
type DescribeJobExecutionResponse struct {
	*types.DescribeJobExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJobExecution request.
func (r *DescribeJobExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
