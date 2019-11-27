// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotdataplane

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane/types"
)

const opUpdateThingShadow = "UpdateThingShadow"

// UpdateThingShadowRequest returns a request value for making API operation for
// AWS IoT Data Plane.
//
// Updates the thing shadow for the specified thing.
//
// For more information, see UpdateThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_UpdateThingShadow.html)
// in the AWS IoT Developer Guide.
//
//    // Example sending a request using UpdateThingShadowRequest.
//    req := client.UpdateThingShadowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateThingShadowRequest(input *types.UpdateThingShadowInput) UpdateThingShadowRequest {
	op := &aws.Operation{
		Name:       opUpdateThingShadow,
		HTTPMethod: "POST",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &types.UpdateThingShadowInput{}
	}

	req := c.newRequest(op, input, &types.UpdateThingShadowOutput{})
	return UpdateThingShadowRequest{Request: req, Input: input, Copy: c.UpdateThingShadowRequest}
}

// UpdateThingShadowRequest is the request type for the
// UpdateThingShadow API operation.
type UpdateThingShadowRequest struct {
	*aws.Request
	Input *types.UpdateThingShadowInput
	Copy  func(*types.UpdateThingShadowInput) UpdateThingShadowRequest
}

// Send marshals and sends the UpdateThingShadow API request.
func (r UpdateThingShadowRequest) Send(ctx context.Context) (*UpdateThingShadowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateThingShadowResponse{
		UpdateThingShadowOutput: r.Request.Data.(*types.UpdateThingShadowOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateThingShadowResponse is the response type for the
// UpdateThingShadow API operation.
type UpdateThingShadowResponse struct {
	*types.UpdateThingShadowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateThingShadow request.
func (r *UpdateThingShadowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
