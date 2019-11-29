// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opAssociateVirtualInterface = "AssociateVirtualInterface"

// AssociateVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Associates a virtual interface with a specified link aggregation group (LAG)
// or connection. Connectivity to AWS is temporarily interrupted as the virtual
// interface is being migrated. If the target connection or LAG has an associated
// virtual interface with a conflicting VLAN number or a conflicting IP address,
// the operation fails.
//
// Virtual interfaces associated with a hosted connection cannot be associated
// with a LAG; hosted connections must be migrated along with their virtual
// interfaces using AssociateHostedConnection.
//
// To reassociate a virtual interface to a new connection or LAG, the requester
// must own either the virtual interface itself or the connection to which the
// virtual interface is currently associated. Additionally, the requester must
// own the connection or LAG for the association.
//
//    // Example sending a request using AssociateVirtualInterfaceRequest.
//    req := client.AssociateVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AssociateVirtualInterface
func (c *Client) AssociateVirtualInterfaceRequest(input *types.AssociateVirtualInterfaceInput) AssociateVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opAssociateVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &types.AssociateVirtualInterfaceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AssociateVirtualInterfaceMarshaler{Input: input}.GetNamedBuildHandler())

	return AssociateVirtualInterfaceRequest{Request: req, Input: input, Copy: c.AssociateVirtualInterfaceRequest}
}

// AssociateVirtualInterfaceRequest is the request type for the
// AssociateVirtualInterface API operation.
type AssociateVirtualInterfaceRequest struct {
	*aws.Request
	Input *types.AssociateVirtualInterfaceInput
	Copy  func(*types.AssociateVirtualInterfaceInput) AssociateVirtualInterfaceRequest
}

// Send marshals and sends the AssociateVirtualInterface API request.
func (r AssociateVirtualInterfaceRequest) Send(ctx context.Context) (*AssociateVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateVirtualInterfaceResponse{
		AssociateVirtualInterfaceOutput: r.Request.Data.(*types.AssociateVirtualInterfaceOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateVirtualInterfaceResponse is the response type for the
// AssociateVirtualInterface API operation.
type AssociateVirtualInterfaceResponse struct {
	*types.AssociateVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateVirtualInterface request.
func (r *AssociateVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
