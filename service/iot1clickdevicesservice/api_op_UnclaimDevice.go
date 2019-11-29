// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickdevicesservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice/types"
)

const opUnclaimDevice = "UnclaimDevice"

// UnclaimDeviceRequest returns a request value for making API operation for
// AWS IoT 1-Click Devices Service.
//
// Disassociates a device from your AWS account using its device ID.
//
//    // Example sending a request using UnclaimDeviceRequest.
//    req := client.UnclaimDeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devices-2018-05-14/UnclaimDevice
func (c *Client) UnclaimDeviceRequest(input *types.UnclaimDeviceInput) UnclaimDeviceRequest {
	op := &aws.Operation{
		Name:       opUnclaimDevice,
		HTTPMethod: "PUT",
		HTTPPath:   "/devices/{deviceId}/unclaim",
	}

	if input == nil {
		input = &types.UnclaimDeviceInput{}
	}

	req := c.newRequest(op, input, &types.UnclaimDeviceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UnclaimDeviceMarshaler{Input: input}.GetNamedBuildHandler())

	return UnclaimDeviceRequest{Request: req, Input: input, Copy: c.UnclaimDeviceRequest}
}

// UnclaimDeviceRequest is the request type for the
// UnclaimDevice API operation.
type UnclaimDeviceRequest struct {
	*aws.Request
	Input *types.UnclaimDeviceInput
	Copy  func(*types.UnclaimDeviceInput) UnclaimDeviceRequest
}

// Send marshals and sends the UnclaimDevice API request.
func (r UnclaimDeviceRequest) Send(ctx context.Context) (*UnclaimDeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UnclaimDeviceResponse{
		UnclaimDeviceOutput: r.Request.Data.(*types.UnclaimDeviceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UnclaimDeviceResponse is the response type for the
// UnclaimDevice API operation.
type UnclaimDeviceResponse struct {
	*types.UnclaimDeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UnclaimDevice request.
func (r *UnclaimDeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
