// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opCreateDirectConnectGateway = "CreateDirectConnectGateway"

// CreateDirectConnectGatewayRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Creates a Direct Connect gateway, which is an intermediate object that enables
// you to connect a set of virtual interfaces and virtual private gateways.
// A Direct Connect gateway is global and visible in any AWS Region after it
// is created. The virtual interfaces and virtual private gateways that are
// connected through a Direct Connect gateway can be in different AWS Regions.
// This enables you to connect to a VPC in any Region, regardless of the Region
// in which the virtual interfaces are located, and pass traffic between them.
//
//    // Example sending a request using CreateDirectConnectGatewayRequest.
//    req := client.CreateDirectConnectGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/CreateDirectConnectGateway
func (c *Client) CreateDirectConnectGatewayRequest(input *types.CreateDirectConnectGatewayInput) CreateDirectConnectGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateDirectConnectGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDirectConnectGatewayInput{}
	}

	req := c.newRequest(op, input, &types.CreateDirectConnectGatewayOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateDirectConnectGatewayMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDirectConnectGatewayRequest{Request: req, Input: input, Copy: c.CreateDirectConnectGatewayRequest}
}

// CreateDirectConnectGatewayRequest is the request type for the
// CreateDirectConnectGateway API operation.
type CreateDirectConnectGatewayRequest struct {
	*aws.Request
	Input *types.CreateDirectConnectGatewayInput
	Copy  func(*types.CreateDirectConnectGatewayInput) CreateDirectConnectGatewayRequest
}

// Send marshals and sends the CreateDirectConnectGateway API request.
func (r CreateDirectConnectGatewayRequest) Send(ctx context.Context) (*CreateDirectConnectGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDirectConnectGatewayResponse{
		CreateDirectConnectGatewayOutput: r.Request.Data.(*types.CreateDirectConnectGatewayOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDirectConnectGatewayResponse is the response type for the
// CreateDirectConnectGateway API operation.
type CreateDirectConnectGatewayResponse struct {
	*types.CreateDirectConnectGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDirectConnectGateway request.
func (r *CreateDirectConnectGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
