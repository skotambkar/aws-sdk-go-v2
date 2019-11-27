// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opConfirmPrivateVirtualInterface = "ConfirmPrivateVirtualInterface"

// ConfirmPrivateVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Accepts ownership of a private virtual interface created by another AWS account.
//
// After the virtual interface owner makes this call, the virtual interface
// is created and attached to the specified virtual private gateway or Direct
// Connect gateway, and is made available to handle traffic.
//
//    // Example sending a request using ConfirmPrivateVirtualInterfaceRequest.
//    req := client.ConfirmPrivateVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/ConfirmPrivateVirtualInterface
func (c *Client) ConfirmPrivateVirtualInterfaceRequest(input *types.ConfirmPrivateVirtualInterfaceInput) ConfirmPrivateVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opConfirmPrivateVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ConfirmPrivateVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &types.ConfirmPrivateVirtualInterfaceOutput{})
	return ConfirmPrivateVirtualInterfaceRequest{Request: req, Input: input, Copy: c.ConfirmPrivateVirtualInterfaceRequest}
}

// ConfirmPrivateVirtualInterfaceRequest is the request type for the
// ConfirmPrivateVirtualInterface API operation.
type ConfirmPrivateVirtualInterfaceRequest struct {
	*aws.Request
	Input *types.ConfirmPrivateVirtualInterfaceInput
	Copy  func(*types.ConfirmPrivateVirtualInterfaceInput) ConfirmPrivateVirtualInterfaceRequest
}

// Send marshals and sends the ConfirmPrivateVirtualInterface API request.
func (r ConfirmPrivateVirtualInterfaceRequest) Send(ctx context.Context) (*ConfirmPrivateVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ConfirmPrivateVirtualInterfaceResponse{
		ConfirmPrivateVirtualInterfaceOutput: r.Request.Data.(*types.ConfirmPrivateVirtualInterfaceOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ConfirmPrivateVirtualInterfaceResponse is the response type for the
// ConfirmPrivateVirtualInterface API operation.
type ConfirmPrivateVirtualInterfaceResponse struct {
	*types.ConfirmPrivateVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ConfirmPrivateVirtualInterface request.
func (r *ConfirmPrivateVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
