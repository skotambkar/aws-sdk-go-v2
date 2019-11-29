// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
)

const opDescribeStackSummary = "DescribeStackSummary"

// DescribeStackSummaryRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Describes the number of layers and apps in a specified stack, and the number
// of instances in each state, such as running_setup or online.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeStackSummaryRequest.
//    req := client.DescribeStackSummaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeStackSummary
func (c *Client) DescribeStackSummaryRequest(input *types.DescribeStackSummaryInput) DescribeStackSummaryRequest {
	op := &aws.Operation{
		Name:       opDescribeStackSummary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeStackSummaryInput{}
	}

	req := c.newRequest(op, input, &types.DescribeStackSummaryOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeStackSummaryMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeStackSummaryRequest{Request: req, Input: input, Copy: c.DescribeStackSummaryRequest}
}

// DescribeStackSummaryRequest is the request type for the
// DescribeStackSummary API operation.
type DescribeStackSummaryRequest struct {
	*aws.Request
	Input *types.DescribeStackSummaryInput
	Copy  func(*types.DescribeStackSummaryInput) DescribeStackSummaryRequest
}

// Send marshals and sends the DescribeStackSummary API request.
func (r DescribeStackSummaryRequest) Send(ctx context.Context) (*DescribeStackSummaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStackSummaryResponse{
		DescribeStackSummaryOutput: r.Request.Data.(*types.DescribeStackSummaryOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStackSummaryResponse is the response type for the
// DescribeStackSummary API operation.
type DescribeStackSummaryResponse struct {
	*types.DescribeStackSummaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStackSummary request.
func (r *DescribeStackSummaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
