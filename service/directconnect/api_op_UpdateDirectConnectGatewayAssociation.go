// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opUpdateDirectConnectGatewayAssociation = "UpdateDirectConnectGatewayAssociation"

// UpdateDirectConnectGatewayAssociationRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Updates the specified attributes of the Direct Connect gateway association.
//
// Add or remove prefixes from the association.
//
//    // Example sending a request using UpdateDirectConnectGatewayAssociationRequest.
//    req := client.UpdateDirectConnectGatewayAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/UpdateDirectConnectGatewayAssociation
func (c *Client) UpdateDirectConnectGatewayAssociationRequest(input *types.UpdateDirectConnectGatewayAssociationInput) UpdateDirectConnectGatewayAssociationRequest {
	op := &aws.Operation{
		Name:       opUpdateDirectConnectGatewayAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateDirectConnectGatewayAssociationInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDirectConnectGatewayAssociationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateDirectConnectGatewayAssociationMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateDirectConnectGatewayAssociationRequest{Request: req, Input: input, Copy: c.UpdateDirectConnectGatewayAssociationRequest}
}

// UpdateDirectConnectGatewayAssociationRequest is the request type for the
// UpdateDirectConnectGatewayAssociation API operation.
type UpdateDirectConnectGatewayAssociationRequest struct {
	*aws.Request
	Input *types.UpdateDirectConnectGatewayAssociationInput
	Copy  func(*types.UpdateDirectConnectGatewayAssociationInput) UpdateDirectConnectGatewayAssociationRequest
}

// Send marshals and sends the UpdateDirectConnectGatewayAssociation API request.
func (r UpdateDirectConnectGatewayAssociationRequest) Send(ctx context.Context) (*UpdateDirectConnectGatewayAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDirectConnectGatewayAssociationResponse{
		UpdateDirectConnectGatewayAssociationOutput: r.Request.Data.(*types.UpdateDirectConnectGatewayAssociationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDirectConnectGatewayAssociationResponse is the response type for the
// UpdateDirectConnectGatewayAssociation API operation.
type UpdateDirectConnectGatewayAssociationResponse struct {
	*types.UpdateDirectConnectGatewayAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDirectConnectGatewayAssociation request.
func (r *UpdateDirectConnectGatewayAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
