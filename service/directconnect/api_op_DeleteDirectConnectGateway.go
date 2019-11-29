// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opDeleteDirectConnectGateway = "DeleteDirectConnectGateway"

// DeleteDirectConnectGatewayRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Deletes the specified Direct Connect gateway. You must first delete all virtual
// interfaces that are attached to the Direct Connect gateway and disassociate
// all virtual private gateways associated with the Direct Connect gateway.
//
//    // Example sending a request using DeleteDirectConnectGatewayRequest.
//    req := client.DeleteDirectConnectGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/DeleteDirectConnectGateway
func (c *Client) DeleteDirectConnectGatewayRequest(input *types.DeleteDirectConnectGatewayInput) DeleteDirectConnectGatewayRequest {
	op := &aws.Operation{
		Name:       opDeleteDirectConnectGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteDirectConnectGatewayInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDirectConnectGatewayOutput{})
	return DeleteDirectConnectGatewayRequest{Request: req, Input: input, Copy: c.DeleteDirectConnectGatewayRequest}
}

// DeleteDirectConnectGatewayRequest is the request type for the
// DeleteDirectConnectGateway API operation.
type DeleteDirectConnectGatewayRequest struct {
	*aws.Request
	Input *types.DeleteDirectConnectGatewayInput
	Copy  func(*types.DeleteDirectConnectGatewayInput) DeleteDirectConnectGatewayRequest
}

// Send marshals and sends the DeleteDirectConnectGateway API request.
func (r DeleteDirectConnectGatewayRequest) Send(ctx context.Context) (*DeleteDirectConnectGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDirectConnectGatewayResponse{
		DeleteDirectConnectGatewayOutput: r.Request.Data.(*types.DeleteDirectConnectGatewayOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDirectConnectGatewayResponse is the response type for the
// DeleteDirectConnectGateway API operation.
type DeleteDirectConnectGatewayResponse struct {
	*types.DeleteDirectConnectGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDirectConnectGateway request.
func (r *DeleteDirectConnectGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
