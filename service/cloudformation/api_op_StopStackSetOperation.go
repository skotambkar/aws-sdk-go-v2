// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opStopStackSetOperation = "StopStackSetOperation"

// StopStackSetOperationRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Stops an in-progress operation on a stack set and its associated stack instances.
//
//    // Example sending a request using StopStackSetOperationRequest.
//    req := client.StopStackSetOperationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/StopStackSetOperation
func (c *Client) StopStackSetOperationRequest(input *types.StopStackSetOperationInput) StopStackSetOperationRequest {
	op := &aws.Operation{
		Name:       opStopStackSetOperation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StopStackSetOperationInput{}
	}

	req := c.newRequest(op, input, &types.StopStackSetOperationOutput{})
	return StopStackSetOperationRequest{Request: req, Input: input, Copy: c.StopStackSetOperationRequest}
}

// StopStackSetOperationRequest is the request type for the
// StopStackSetOperation API operation.
type StopStackSetOperationRequest struct {
	*aws.Request
	Input *types.StopStackSetOperationInput
	Copy  func(*types.StopStackSetOperationInput) StopStackSetOperationRequest
}

// Send marshals and sends the StopStackSetOperation API request.
func (r StopStackSetOperationRequest) Send(ctx context.Context) (*StopStackSetOperationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopStackSetOperationResponse{
		StopStackSetOperationOutput: r.Request.Data.(*types.StopStackSetOperationOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopStackSetOperationResponse is the response type for the
// StopStackSetOperation API operation.
type StopStackSetOperationResponse struct {
	*types.StopStackSetOperationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopStackSetOperation request.
func (r *StopStackSetOperationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
