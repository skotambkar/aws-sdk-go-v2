// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotevents/types"
)

const opUpdateInput = "UpdateInput"

// UpdateInputRequest returns a request value for making API operation for
// AWS IoT Events.
//
// Updates an input.
//
//    // Example sending a request using UpdateInputRequest.
//    req := client.UpdateInputRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotevents-2018-07-27/UpdateInput
func (c *Client) UpdateInputRequest(input *types.UpdateInputInput) UpdateInputRequest {
	op := &aws.Operation{
		Name:       opUpdateInput,
		HTTPMethod: "PUT",
		HTTPPath:   "/inputs/{inputName}",
	}

	if input == nil {
		input = &types.UpdateInputInput{}
	}

	req := c.newRequest(op, input, &types.UpdateInputOutput{})
	return UpdateInputRequest{Request: req, Input: input, Copy: c.UpdateInputRequest}
}

// UpdateInputRequest is the request type for the
// UpdateInput API operation.
type UpdateInputRequest struct {
	*aws.Request
	Input *types.UpdateInputInput
	Copy  func(*types.UpdateInputInput) UpdateInputRequest
}

// Send marshals and sends the UpdateInput API request.
func (r UpdateInputRequest) Send(ctx context.Context) (*UpdateInputResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateInputResponse{
		UpdateInputOutput: r.Request.Data.(*types.UpdateInputOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateInputResponse is the response type for the
// UpdateInput API operation.
type UpdateInputResponse struct {
	*types.UpdateInputOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateInput request.
func (r *UpdateInputResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
