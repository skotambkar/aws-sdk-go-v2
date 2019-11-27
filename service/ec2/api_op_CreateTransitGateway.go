// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateTransitGateway = "CreateTransitGateway"

// CreateTransitGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a transit gateway.
//
// You can use a transit gateway to interconnect your virtual private clouds
// (VPC) and on-premises networks. After the transit gateway enters the available
// state, you can attach your VPCs and VPN connections to the transit gateway.
//
// To attach your VPCs, use CreateTransitGatewayVpcAttachment.
//
// To attach a VPN connection, use CreateCustomerGateway to create a customer
// gateway and specify the ID of the customer gateway and the ID of the transit
// gateway in a call to CreateVpnConnection.
//
// When you create a transit gateway, we create a default transit gateway route
// table and use it as the default association route table and the default propagation
// route table. You can use CreateTransitGatewayRouteTable to create additional
// transit gateway route tables. If you disable automatic route propagation,
// we do not create a default transit gateway route table. You can use EnableTransitGatewayRouteTablePropagation
// to propagate routes from a resource attachment to a transit gateway route
// table. If you disable automatic associations, you can use AssociateTransitGatewayRouteTable
// to associate a resource attachment with a transit gateway route table.
//
//    // Example sending a request using CreateTransitGatewayRequest.
//    req := client.CreateTransitGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGateway
func (c *Client) CreateTransitGatewayRequest(input *types.CreateTransitGatewayInput) CreateTransitGatewayRequest {
	op := &aws.Operation{
		Name:       opCreateTransitGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateTransitGatewayInput{}
	}

	req := c.newRequest(op, input, &types.CreateTransitGatewayOutput{})
	return CreateTransitGatewayRequest{Request: req, Input: input, Copy: c.CreateTransitGatewayRequest}
}

// CreateTransitGatewayRequest is the request type for the
// CreateTransitGateway API operation.
type CreateTransitGatewayRequest struct {
	*aws.Request
	Input *types.CreateTransitGatewayInput
	Copy  func(*types.CreateTransitGatewayInput) CreateTransitGatewayRequest
}

// Send marshals and sends the CreateTransitGateway API request.
func (r CreateTransitGatewayRequest) Send(ctx context.Context) (*CreateTransitGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTransitGatewayResponse{
		CreateTransitGatewayOutput: r.Request.Data.(*types.CreateTransitGatewayOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTransitGatewayResponse is the response type for the
// CreateTransitGateway API operation.
type CreateTransitGatewayResponse struct {
	*types.CreateTransitGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTransitGateway request.
func (r *CreateTransitGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
