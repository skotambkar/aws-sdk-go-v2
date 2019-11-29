// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opCreatePrivateVirtualInterface = "CreatePrivateVirtualInterface"

// CreatePrivateVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates a private virtual interface. A virtual interface is the VLAN that
// transports AWS Direct Connect traffic. A private virtual interface can be
// connected to either a Direct Connect gateway or a Virtual Private Gateway
// (VGW). Connecting the private virtual interface to a Direct Connect gateway
// enables the possibility for connecting to multiple VPCs, including VPCs in
// different AWS Regions. Connecting the private virtual interface to a VGW
// only provides access to a single VPC within the same Region.
//
//    // Example sending a request using CreatePrivateVirtualInterfaceRequest.
//    req := client.CreatePrivateVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreatePrivateVirtualInterface
func (c *Client) CreatePrivateVirtualInterfaceRequest(input *types.CreatePrivateVirtualInterfaceInput) CreatePrivateVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opCreatePrivateVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreatePrivateVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &types.CreatePrivateVirtualInterfaceOutput{})
	return CreatePrivateVirtualInterfaceRequest{Request: req, Input: input, Copy: c.CreatePrivateVirtualInterfaceRequest}
}

// CreatePrivateVirtualInterfaceRequest is the request type for the
// CreatePrivateVirtualInterface API operation.
type CreatePrivateVirtualInterfaceRequest struct {
	*aws.Request
	Input *types.CreatePrivateVirtualInterfaceInput
	Copy  func(*types.CreatePrivateVirtualInterfaceInput) CreatePrivateVirtualInterfaceRequest
}

// Send marshals and sends the CreatePrivateVirtualInterface API request.
func (r CreatePrivateVirtualInterfaceRequest) Send(ctx context.Context) (*CreatePrivateVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePrivateVirtualInterfaceResponse{
		CreatePrivateVirtualInterfaceOutput: r.Request.Data.(*types.CreatePrivateVirtualInterfaceOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePrivateVirtualInterfaceResponse is the response type for the
// CreatePrivateVirtualInterface API operation.
type CreatePrivateVirtualInterfaceResponse struct {
	*types.CreatePrivateVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePrivateVirtualInterface request.
func (r *CreatePrivateVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
