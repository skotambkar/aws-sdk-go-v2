// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opDeleteVirtualInterface = "DeleteVirtualInterface"

// DeleteVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deletes a virtual interface.
//
//    // Example sending a request using DeleteVirtualInterfaceRequest.
//    req := client.DeleteVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DeleteVirtualInterface
func (c *Client) DeleteVirtualInterfaceRequest(input *types.DeleteVirtualInterfaceInput) DeleteVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opDeleteVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &types.DeleteVirtualInterfaceOutput{})
	return DeleteVirtualInterfaceRequest{Request: req, Input: input, Copy: c.DeleteVirtualInterfaceRequest}
}

// DeleteVirtualInterfaceRequest is the request type for the
// DeleteVirtualInterface API operation.
type DeleteVirtualInterfaceRequest struct {
	*aws.Request
	Input *types.DeleteVirtualInterfaceInput
	Copy  func(*types.DeleteVirtualInterfaceInput) DeleteVirtualInterfaceRequest
}

// Send marshals and sends the DeleteVirtualInterface API request.
func (r DeleteVirtualInterfaceRequest) Send(ctx context.Context) (*DeleteVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVirtualInterfaceResponse{
		DeleteVirtualInterfaceOutput: r.Request.Data.(*types.DeleteVirtualInterfaceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVirtualInterfaceResponse is the response type for the
// DeleteVirtualInterface API operation.
type DeleteVirtualInterfaceResponse struct {
	*types.DeleteVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVirtualInterface request.
func (r *DeleteVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
