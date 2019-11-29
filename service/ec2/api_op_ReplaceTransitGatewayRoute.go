// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opReplaceTransitGatewayRoute = "ReplaceTransitGatewayRoute"

// ReplaceTransitGatewayRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Replaces the specified route in the specified transit gateway route table.
//
//    // Example sending a request using ReplaceTransitGatewayRouteRequest.
//    req := client.ReplaceTransitGatewayRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReplaceTransitGatewayRoute
func (c *Client) ReplaceTransitGatewayRouteRequest(input *types.ReplaceTransitGatewayRouteInput) ReplaceTransitGatewayRouteRequest {
	op := &aws.Operation{
		Name:       opReplaceTransitGatewayRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ReplaceTransitGatewayRouteInput{}
	}

	req := c.newRequest(op, input, &types.ReplaceTransitGatewayRouteOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.ReplaceTransitGatewayRouteMarshaler{Input: input}.GetNamedBuildHandler())

	return ReplaceTransitGatewayRouteRequest{Request: req, Input: input, Copy: c.ReplaceTransitGatewayRouteRequest}
}

// ReplaceTransitGatewayRouteRequest is the request type for the
// ReplaceTransitGatewayRoute API operation.
type ReplaceTransitGatewayRouteRequest struct {
	*aws.Request
	Input *types.ReplaceTransitGatewayRouteInput
	Copy  func(*types.ReplaceTransitGatewayRouteInput) ReplaceTransitGatewayRouteRequest
}

// Send marshals and sends the ReplaceTransitGatewayRoute API request.
func (r ReplaceTransitGatewayRouteRequest) Send(ctx context.Context) (*ReplaceTransitGatewayRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReplaceTransitGatewayRouteResponse{
		ReplaceTransitGatewayRouteOutput: r.Request.Data.(*types.ReplaceTransitGatewayRouteOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReplaceTransitGatewayRouteResponse is the response type for the
// ReplaceTransitGatewayRoute API operation.
type ReplaceTransitGatewayRouteResponse struct {
	*types.ReplaceTransitGatewayRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReplaceTransitGatewayRoute request.
func (r *ReplaceTransitGatewayRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
