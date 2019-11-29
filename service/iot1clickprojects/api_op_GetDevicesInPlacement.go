// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot1clickprojects

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects/types"
)

const opGetDevicesInPlacement = "GetDevicesInPlacement"

// GetDevicesInPlacementRequest returns a request value for making API operation for
// AWS IoT 1-Click Projects Service.
//
// Returns an object enumerating the devices in a placement.
//
//    // Example sending a request using GetDevicesInPlacementRequest.
//    req := client.GetDevicesInPlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot1click-projects-2018-05-14/GetDevicesInPlacement
func (c *Client) GetDevicesInPlacementRequest(input *types.GetDevicesInPlacementInput) GetDevicesInPlacementRequest {
	op := &aws.Operation{
		Name:       opGetDevicesInPlacement,
		HTTPMethod: "GET",
		HTTPPath:   "/projects/{projectName}/placements/{placementName}/devices",
	}

	if input == nil {
		input = &types.GetDevicesInPlacementInput{}
	}

	req := c.newRequest(op, input, &types.GetDevicesInPlacementOutput{})
	return GetDevicesInPlacementRequest{Request: req, Input: input, Copy: c.GetDevicesInPlacementRequest}
}

// GetDevicesInPlacementRequest is the request type for the
// GetDevicesInPlacement API operation.
type GetDevicesInPlacementRequest struct {
	*aws.Request
	Input *types.GetDevicesInPlacementInput
	Copy  func(*types.GetDevicesInPlacementInput) GetDevicesInPlacementRequest
}

// Send marshals and sends the GetDevicesInPlacement API request.
func (r GetDevicesInPlacementRequest) Send(ctx context.Context) (*GetDevicesInPlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDevicesInPlacementResponse{
		GetDevicesInPlacementOutput: r.Request.Data.(*types.GetDevicesInPlacementOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDevicesInPlacementResponse is the response type for the
// GetDevicesInPlacement API operation.
type GetDevicesInPlacementResponse struct {
	*types.GetDevicesInPlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDevicesInPlacement request.
func (r *GetDevicesInPlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
