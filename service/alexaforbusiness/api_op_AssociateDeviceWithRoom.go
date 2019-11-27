// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opAssociateDeviceWithRoom = "AssociateDeviceWithRoom"

// AssociateDeviceWithRoomRequest returns a request value for making API operation for
// Alexa For Business.
//
// Associates a device with a given room. This applies all the settings from
// the room profile to the device, and all the skills in any skill groups added
// to that room. This operation requires the device to be online, or else a
// manual sync is required.
//
//    // Example sending a request using AssociateDeviceWithRoomRequest.
//    req := client.AssociateDeviceWithRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/AssociateDeviceWithRoom
func (c *Client) AssociateDeviceWithRoomRequest(input *types.AssociateDeviceWithRoomInput) AssociateDeviceWithRoomRequest {
	op := &aws.Operation{
		Name:       opAssociateDeviceWithRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateDeviceWithRoomInput{}
	}

	req := c.newRequest(op, input, &types.AssociateDeviceWithRoomOutput{})
	return AssociateDeviceWithRoomRequest{Request: req, Input: input, Copy: c.AssociateDeviceWithRoomRequest}
}

// AssociateDeviceWithRoomRequest is the request type for the
// AssociateDeviceWithRoom API operation.
type AssociateDeviceWithRoomRequest struct {
	*aws.Request
	Input *types.AssociateDeviceWithRoomInput
	Copy  func(*types.AssociateDeviceWithRoomInput) AssociateDeviceWithRoomRequest
}

// Send marshals and sends the AssociateDeviceWithRoom API request.
func (r AssociateDeviceWithRoomRequest) Send(ctx context.Context) (*AssociateDeviceWithRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateDeviceWithRoomResponse{
		AssociateDeviceWithRoomOutput: r.Request.Data.(*types.AssociateDeviceWithRoomOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateDeviceWithRoomResponse is the response type for the
// AssociateDeviceWithRoom API operation.
type AssociateDeviceWithRoomResponse struct {
	*types.AssociateDeviceWithRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateDeviceWithRoom request.
func (r *AssociateDeviceWithRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
