// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opAllocateTransitVirtualInterface = "AllocateTransitVirtualInterface"

// AllocateTransitVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Provisions a transit virtual interface to be owned by the specified AWS account.
// Use this type of interface to connect a transit gateway to your Direct Connect
// gateway.
//
// The owner of a connection provisions a transit virtual interface to be owned
// by the specified AWS account.
//
// After you create a transit virtual interface, it must be confirmed by the
// owner using ConfirmTransitVirtualInterface. Until this step has been completed,
// the transit virtual interface is in the requested state and is not available
// to handle traffic.
//
//    // Example sending a request using AllocateTransitVirtualInterfaceRequest.
//    req := client.AllocateTransitVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocateTransitVirtualInterface
func (c *Client) AllocateTransitVirtualInterfaceRequest(input *types.AllocateTransitVirtualInterfaceInput) AllocateTransitVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opAllocateTransitVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AllocateTransitVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &types.AllocateTransitVirtualInterfaceOutput{})
	return AllocateTransitVirtualInterfaceRequest{Request: req, Input: input, Copy: c.AllocateTransitVirtualInterfaceRequest}
}

// AllocateTransitVirtualInterfaceRequest is the request type for the
// AllocateTransitVirtualInterface API operation.
type AllocateTransitVirtualInterfaceRequest struct {
	*aws.Request
	Input *types.AllocateTransitVirtualInterfaceInput
	Copy  func(*types.AllocateTransitVirtualInterfaceInput) AllocateTransitVirtualInterfaceRequest
}

// Send marshals and sends the AllocateTransitVirtualInterface API request.
func (r AllocateTransitVirtualInterfaceRequest) Send(ctx context.Context) (*AllocateTransitVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AllocateTransitVirtualInterfaceResponse{
		AllocateTransitVirtualInterfaceOutput: r.Request.Data.(*types.AllocateTransitVirtualInterfaceOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AllocateTransitVirtualInterfaceResponse is the response type for the
// AllocateTransitVirtualInterface API operation.
type AllocateTransitVirtualInterfaceResponse struct {
	*types.AllocateTransitVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AllocateTransitVirtualInterface request.
func (r *AllocateTransitVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
