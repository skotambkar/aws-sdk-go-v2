// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
)

const opDescribeAgent = "DescribeAgent"

// DescribeAgentRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns metadata such as the name, the network interfaces, and the status
// (that is, whether the agent is running or not) for an agent. To specify which
// agent to describe, use the Amazon Resource Name (ARN) of the agent in your
// request.
//
//    // Example sending a request using DescribeAgentRequest.
//    req := client.DescribeAgentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeAgent
func (c *Client) DescribeAgentRequest(input *types.DescribeAgentInput) DescribeAgentRequest {
	op := &aws.Operation{
		Name:       opDescribeAgent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeAgentInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAgentOutput{})
	return DescribeAgentRequest{Request: req, Input: input, Copy: c.DescribeAgentRequest}
}

// DescribeAgentRequest is the request type for the
// DescribeAgent API operation.
type DescribeAgentRequest struct {
	*aws.Request
	Input *types.DescribeAgentInput
	Copy  func(*types.DescribeAgentInput) DescribeAgentRequest
}

// Send marshals and sends the DescribeAgent API request.
func (r DescribeAgentRequest) Send(ctx context.Context) (*DescribeAgentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAgentResponse{
		DescribeAgentOutput: r.Request.Data.(*types.DescribeAgentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAgentResponse is the response type for the
// DescribeAgent API operation.
type DescribeAgentResponse struct {
	*types.DescribeAgentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAgent request.
func (r *DescribeAgentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
