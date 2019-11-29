// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sfn/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
)

const opDescribeStateMachine = "DescribeStateMachine"

// DescribeStateMachineRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Describes a state machine.
//
// This operation is eventually consistent. The results are best effort and
// may not reflect very recent updates and changes.
//
//    // Example sending a request using DescribeStateMachineRequest.
//    req := client.DescribeStateMachineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/DescribeStateMachine
func (c *Client) DescribeStateMachineRequest(input *types.DescribeStateMachineInput) DescribeStateMachineRequest {
	op := &aws.Operation{
		Name:       opDescribeStateMachine,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeStateMachineInput{}
	}

	req := c.newRequest(op, input, &types.DescribeStateMachineOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeStateMachineMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeStateMachineRequest{Request: req, Input: input, Copy: c.DescribeStateMachineRequest}
}

// DescribeStateMachineRequest is the request type for the
// DescribeStateMachine API operation.
type DescribeStateMachineRequest struct {
	*aws.Request
	Input *types.DescribeStateMachineInput
	Copy  func(*types.DescribeStateMachineInput) DescribeStateMachineRequest
}

// Send marshals and sends the DescribeStateMachine API request.
func (r DescribeStateMachineRequest) Send(ctx context.Context) (*DescribeStateMachineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStateMachineResponse{
		DescribeStateMachineOutput: r.Request.Data.(*types.DescribeStateMachineOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStateMachineResponse is the response type for the
// DescribeStateMachine API operation.
type DescribeStateMachineResponse struct {
	*types.DescribeStateMachineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStateMachine request.
func (r *DescribeStateMachineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
