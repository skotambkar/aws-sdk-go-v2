// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateTransitGatewayRouteTable = "CreateTransitGatewayRouteTable"

// CreateTransitGatewayRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a route table for the specified transit gateway.
//
//    // Example sending a request using CreateTransitGatewayRouteTableRequest.
//    req := client.CreateTransitGatewayRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayRouteTable
func (c *Client) CreateTransitGatewayRouteTableRequest(input *types.CreateTransitGatewayRouteTableInput) CreateTransitGatewayRouteTableRequest {
	op := &aws.Operation{
		Name:       opCreateTransitGatewayRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateTransitGatewayRouteTableInput{}
	}

	req := c.newRequest(op, input, &types.CreateTransitGatewayRouteTableOutput{})
	return CreateTransitGatewayRouteTableRequest{Request: req, Input: input, Copy: c.CreateTransitGatewayRouteTableRequest}
}

// CreateTransitGatewayRouteTableRequest is the request type for the
// CreateTransitGatewayRouteTable API operation.
type CreateTransitGatewayRouteTableRequest struct {
	*aws.Request
	Input *types.CreateTransitGatewayRouteTableInput
	Copy  func(*types.CreateTransitGatewayRouteTableInput) CreateTransitGatewayRouteTableRequest
}

// Send marshals and sends the CreateTransitGatewayRouteTable API request.
func (r CreateTransitGatewayRouteTableRequest) Send(ctx context.Context) (*CreateTransitGatewayRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTransitGatewayRouteTableResponse{
		CreateTransitGatewayRouteTableOutput: r.Request.Data.(*types.CreateTransitGatewayRouteTableOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTransitGatewayRouteTableResponse is the response type for the
// CreateTransitGatewayRouteTable API operation.
type CreateTransitGatewayRouteTableResponse struct {
	*types.CreateTransitGatewayRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTransitGatewayRouteTable request.
func (r *CreateTransitGatewayRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
