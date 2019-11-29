// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opRegisterThing = "RegisterThing"

// RegisterThingRequest returns a request value for making API operation for
// AWS IoT.
//
// Provisions a thing.
//
//    // Example sending a request using RegisterThingRequest.
//    req := client.RegisterThingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RegisterThingRequest(input *types.RegisterThingInput) RegisterThingRequest {
	op := &aws.Operation{
		Name:       opRegisterThing,
		HTTPMethod: "POST",
		HTTPPath:   "/things",
	}

	if input == nil {
		input = &types.RegisterThingInput{}
	}

	req := c.newRequest(op, input, &types.RegisterThingOutput{})
	return RegisterThingRequest{Request: req, Input: input, Copy: c.RegisterThingRequest}
}

// RegisterThingRequest is the request type for the
// RegisterThing API operation.
type RegisterThingRequest struct {
	*aws.Request
	Input *types.RegisterThingInput
	Copy  func(*types.RegisterThingInput) RegisterThingRequest
}

// Send marshals and sends the RegisterThing API request.
func (r RegisterThingRequest) Send(ctx context.Context) (*RegisterThingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterThingResponse{
		RegisterThingOutput: r.Request.Data.(*types.RegisterThingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterThingResponse is the response type for the
// RegisterThing API operation.
type RegisterThingResponse struct {
	*types.RegisterThingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterThing request.
func (r *RegisterThingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
