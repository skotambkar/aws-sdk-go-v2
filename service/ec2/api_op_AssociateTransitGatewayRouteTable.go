// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opAssociateTransitGatewayRouteTable = "AssociateTransitGatewayRouteTable"

// AssociateTransitGatewayRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates the specified attachment with the specified transit gateway route
// table. You can associate only one route table with an attachment.
//
//    // Example sending a request using AssociateTransitGatewayRouteTableRequest.
//    req := client.AssociateTransitGatewayRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateTransitGatewayRouteTable
func (c *Client) AssociateTransitGatewayRouteTableRequest(input *types.AssociateTransitGatewayRouteTableInput) AssociateTransitGatewayRouteTableRequest {
	op := &aws.Operation{
		Name:       opAssociateTransitGatewayRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateTransitGatewayRouteTableInput{}
	}

	req := c.newRequest(op, input, &types.AssociateTransitGatewayRouteTableOutput{})
	return AssociateTransitGatewayRouteTableRequest{Request: req, Input: input, Copy: c.AssociateTransitGatewayRouteTableRequest}
}

// AssociateTransitGatewayRouteTableRequest is the request type for the
// AssociateTransitGatewayRouteTable API operation.
type AssociateTransitGatewayRouteTableRequest struct {
	*aws.Request
	Input *types.AssociateTransitGatewayRouteTableInput
	Copy  func(*types.AssociateTransitGatewayRouteTableInput) AssociateTransitGatewayRouteTableRequest
}

// Send marshals and sends the AssociateTransitGatewayRouteTable API request.
func (r AssociateTransitGatewayRouteTableRequest) Send(ctx context.Context) (*AssociateTransitGatewayRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateTransitGatewayRouteTableResponse{
		AssociateTransitGatewayRouteTableOutput: r.Request.Data.(*types.AssociateTransitGatewayRouteTableOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateTransitGatewayRouteTableResponse is the response type for the
// AssociateTransitGatewayRouteTable API operation.
type AssociateTransitGatewayRouteTableResponse struct {
	*types.AssociateTransitGatewayRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateTransitGatewayRouteTable request.
func (r *AssociateTransitGatewayRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
