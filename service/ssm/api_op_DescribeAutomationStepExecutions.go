// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opDescribeAutomationStepExecutions = "DescribeAutomationStepExecutions"

// DescribeAutomationStepExecutionsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Information about all active and terminated step executions in an Automation
// workflow.
//
//    // Example sending a request using DescribeAutomationStepExecutionsRequest.
//    req := client.DescribeAutomationStepExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeAutomationStepExecutions
func (c *Client) DescribeAutomationStepExecutionsRequest(input *types.DescribeAutomationStepExecutionsInput) DescribeAutomationStepExecutionsRequest {
	op := &aws.Operation{
		Name:       opDescribeAutomationStepExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeAutomationStepExecutionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAutomationStepExecutionsOutput{})
	return DescribeAutomationStepExecutionsRequest{Request: req, Input: input, Copy: c.DescribeAutomationStepExecutionsRequest}
}

// DescribeAutomationStepExecutionsRequest is the request type for the
// DescribeAutomationStepExecutions API operation.
type DescribeAutomationStepExecutionsRequest struct {
	*aws.Request
	Input *types.DescribeAutomationStepExecutionsInput
	Copy  func(*types.DescribeAutomationStepExecutionsInput) DescribeAutomationStepExecutionsRequest
}

// Send marshals and sends the DescribeAutomationStepExecutions API request.
func (r DescribeAutomationStepExecutionsRequest) Send(ctx context.Context) (*DescribeAutomationStepExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAutomationStepExecutionsResponse{
		DescribeAutomationStepExecutionsOutput: r.Request.Data.(*types.DescribeAutomationStepExecutionsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAutomationStepExecutionsResponse is the response type for the
// DescribeAutomationStepExecutions API operation.
type DescribeAutomationStepExecutionsResponse struct {
	*types.DescribeAutomationStepExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAutomationStepExecutions request.
func (r *DescribeAutomationStepExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
