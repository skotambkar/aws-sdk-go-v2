// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDisassociateTransitGatewayRouteTable = "DisassociateTransitGatewayRouteTable"

// DisassociateTransitGatewayRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disassociates a resource attachment from a transit gateway route table.
//
//    // Example sending a request using DisassociateTransitGatewayRouteTableRequest.
//    req := client.DisassociateTransitGatewayRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateTransitGatewayRouteTable
func (c *Client) DisassociateTransitGatewayRouteTableRequest(input *types.DisassociateTransitGatewayRouteTableInput) DisassociateTransitGatewayRouteTableRequest {
	op := &aws.Operation{
		Name:       opDisassociateTransitGatewayRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociateTransitGatewayRouteTableInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateTransitGatewayRouteTableOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DisassociateTransitGatewayRouteTableMarshaler{Input: input}.GetNamedBuildHandler())

	return DisassociateTransitGatewayRouteTableRequest{Request: req, Input: input, Copy: c.DisassociateTransitGatewayRouteTableRequest}
}

// DisassociateTransitGatewayRouteTableRequest is the request type for the
// DisassociateTransitGatewayRouteTable API operation.
type DisassociateTransitGatewayRouteTableRequest struct {
	*aws.Request
	Input *types.DisassociateTransitGatewayRouteTableInput
	Copy  func(*types.DisassociateTransitGatewayRouteTableInput) DisassociateTransitGatewayRouteTableRequest
}

// Send marshals and sends the DisassociateTransitGatewayRouteTable API request.
func (r DisassociateTransitGatewayRouteTableRequest) Send(ctx context.Context) (*DisassociateTransitGatewayRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateTransitGatewayRouteTableResponse{
		DisassociateTransitGatewayRouteTableOutput: r.Request.Data.(*types.DisassociateTransitGatewayRouteTableOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateTransitGatewayRouteTableResponse is the response type for the
// DisassociateTransitGatewayRouteTable API operation.
type DisassociateTransitGatewayRouteTableResponse struct {
	*types.DisassociateTransitGatewayRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateTransitGatewayRouteTable request.
func (r *DisassociateTransitGatewayRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
