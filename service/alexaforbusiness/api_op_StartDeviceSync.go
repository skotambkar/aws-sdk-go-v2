// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opStartDeviceSync = "StartDeviceSync"

// StartDeviceSyncRequest returns a request value for making API operation for
// Alexa For Business.
//
// Resets a device and its account to the known default settings. This clears
// all information and settings set by previous users in the following ways:
//
//    * Bluetooth - This unpairs all bluetooth devices paired with your echo
//    device.
//
//    * Volume - This resets the echo device's volume to the default value.
//
//    * Notifications - This clears all notifications from your echo device.
//
//    * Lists - This clears all to-do items from your echo device.
//
//    * Settings - This internally syncs the room's profile (if the device is
//    assigned to a room), contacts, address books, delegation access for account
//    linking, and communications (if enabled on the room profile).
//
//    // Example sending a request using StartDeviceSyncRequest.
//    req := client.StartDeviceSyncRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/StartDeviceSync
func (c *Client) StartDeviceSyncRequest(input *types.StartDeviceSyncInput) StartDeviceSyncRequest {
	op := &aws.Operation{
		Name:       opStartDeviceSync,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.StartDeviceSyncInput{}
	}

	req := c.newRequest(op, input, &types.StartDeviceSyncOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.StartDeviceSyncMarshaler{Input: input}.GetNamedBuildHandler())

	return StartDeviceSyncRequest{Request: req, Input: input, Copy: c.StartDeviceSyncRequest}
}

// StartDeviceSyncRequest is the request type for the
// StartDeviceSync API operation.
type StartDeviceSyncRequest struct {
	*aws.Request
	Input *types.StartDeviceSyncInput
	Copy  func(*types.StartDeviceSyncInput) StartDeviceSyncRequest
}

// Send marshals and sends the StartDeviceSync API request.
func (r StartDeviceSyncRequest) Send(ctx context.Context) (*StartDeviceSyncResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDeviceSyncResponse{
		StartDeviceSyncOutput: r.Request.Data.(*types.StartDeviceSyncOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDeviceSyncResponse is the response type for the
// StartDeviceSync API operation.
type StartDeviceSyncResponse struct {
	*types.StartDeviceSyncOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDeviceSync request.
func (r *StartDeviceSyncResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
