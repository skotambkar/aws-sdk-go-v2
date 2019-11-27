// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitosync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
)

const opRegisterDevice = "RegisterDevice"

// RegisterDeviceRequest returns a request value for making API operation for
// Amazon Cognito Sync.
//
// Registers a device to receive push sync notifications.
//
// This API can only be called with temporary credentials provided by Cognito
// Identity. You cannot call this API with developer credentials.
//
//    // Example sending a request using RegisterDeviceRequest.
//    req := client.RegisterDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-sync-2014-06-30/RegisterDevice
func (c *Client) RegisterDeviceRequest(input *types.RegisterDeviceInput) RegisterDeviceRequest {
	op := &aws.Operation{
		Name:       opRegisterDevice,
		HTTPMethod: "POST",
		HTTPPath:   "/identitypools/{IdentityPoolId}/identity/{IdentityId}/device",
	}

	if input == nil {
		input = &types.RegisterDeviceInput{}
	}

	req := c.newRequest(op, input, &types.RegisterDeviceOutput{})
	return RegisterDeviceRequest{Request: req, Input: input, Copy: c.RegisterDeviceRequest}
}

// RegisterDeviceRequest is the request type for the
// RegisterDevice API operation.
type RegisterDeviceRequest struct {
	*aws.Request
	Input *types.RegisterDeviceInput
	Copy  func(*types.RegisterDeviceInput) RegisterDeviceRequest
}

// Send marshals and sends the RegisterDevice API request.
func (r RegisterDeviceRequest) Send(ctx context.Context) (*RegisterDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterDeviceResponse{
		RegisterDeviceOutput: r.Request.Data.(*types.RegisterDeviceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterDeviceResponse is the response type for the
// RegisterDevice API operation.
type RegisterDeviceResponse struct {
	*types.RegisterDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterDevice request.
func (r *RegisterDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
