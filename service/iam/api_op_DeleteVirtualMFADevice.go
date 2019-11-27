// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opDeleteVirtualMFADevice = "DeleteVirtualMFADevice"

// DeleteVirtualMFADeviceRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes a virtual MFA device.
//
// You must deactivate a user's virtual MFA device before you can delete it.
// For information about deactivating MFA devices, see DeactivateMFADevice.
//
//    // Example sending a request using DeleteVirtualMFADeviceRequest.
//    req := client.DeleteVirtualMFADeviceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteVirtualMFADevice
func (c *Client) DeleteVirtualMFADeviceRequest(input *types.DeleteVirtualMFADeviceInput) DeleteVirtualMFADeviceRequest {
	op := &aws.Operation{
		Name:       opDeleteVirtualMFADevice,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteVirtualMFADeviceInput{}
	}

	req := c.newRequest(op, input, &types.DeleteVirtualMFADeviceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVirtualMFADeviceRequest{Request: req, Input: input, Copy: c.DeleteVirtualMFADeviceRequest}
}

// DeleteVirtualMFADeviceRequest is the request type for the
// DeleteVirtualMFADevice API operation.
type DeleteVirtualMFADeviceRequest struct {
	*aws.Request
	Input *types.DeleteVirtualMFADeviceInput
	Copy  func(*types.DeleteVirtualMFADeviceInput) DeleteVirtualMFADeviceRequest
}

// Send marshals and sends the DeleteVirtualMFADevice API request.
func (r DeleteVirtualMFADeviceRequest) Send(ctx context.Context) (*DeleteVirtualMFADeviceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVirtualMFADeviceResponse{
		DeleteVirtualMFADeviceOutput: r.Request.Data.(*types.DeleteVirtualMFADeviceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVirtualMFADeviceResponse is the response type for the
// DeleteVirtualMFADevice API operation.
type DeleteVirtualMFADeviceResponse struct {
	*types.DeleteVirtualMFADeviceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVirtualMFADevice request.
func (r *DeleteVirtualMFADeviceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
