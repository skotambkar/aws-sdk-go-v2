// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opEnableTransitGatewayRouteTablePropagation = "EnableTransitGatewayRouteTablePropagation"

// EnableTransitGatewayRouteTablePropagationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables the specified attachment to propagate routes to the specified propagation
// route table.
//
//    // Example sending a request using EnableTransitGatewayRouteTablePropagationRequest.
//    req := client.EnableTransitGatewayRouteTablePropagationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableTransitGatewayRouteTablePropagation
func (c *Client) EnableTransitGatewayRouteTablePropagationRequest(input *types.EnableTransitGatewayRouteTablePropagationInput) EnableTransitGatewayRouteTablePropagationRequest {
	op := &aws.Operation{
		Name:       opEnableTransitGatewayRouteTablePropagation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.EnableTransitGatewayRouteTablePropagationInput{}
	}

	req := c.newRequest(op, input, &types.EnableTransitGatewayRouteTablePropagationOutput{})
	return EnableTransitGatewayRouteTablePropagationRequest{Request: req, Input: input, Copy: c.EnableTransitGatewayRouteTablePropagationRequest}
}

// EnableTransitGatewayRouteTablePropagationRequest is the request type for the
// EnableTransitGatewayRouteTablePropagation API operation.
type EnableTransitGatewayRouteTablePropagationRequest struct {
	*aws.Request
	Input *types.EnableTransitGatewayRouteTablePropagationInput
	Copy  func(*types.EnableTransitGatewayRouteTablePropagationInput) EnableTransitGatewayRouteTablePropagationRequest
}

// Send marshals and sends the EnableTransitGatewayRouteTablePropagation API request.
func (r EnableTransitGatewayRouteTablePropagationRequest) Send(ctx context.Context) (*EnableTransitGatewayRouteTablePropagationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableTransitGatewayRouteTablePropagationResponse{
		EnableTransitGatewayRouteTablePropagationOutput: r.Request.Data.(*types.EnableTransitGatewayRouteTablePropagationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableTransitGatewayRouteTablePropagationResponse is the response type for the
// EnableTransitGatewayRouteTablePropagation API operation.
type EnableTransitGatewayRouteTablePropagationResponse struct {
	*types.EnableTransitGatewayRouteTablePropagationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableTransitGatewayRouteTablePropagation request.
func (r *EnableTransitGatewayRouteTablePropagationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
