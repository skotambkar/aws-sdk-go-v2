// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateTransitGatewayRoute = "CreateTransitGatewayRoute"

// CreateTransitGatewayRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a static route for the specified transit gateway route table.
//
//    // Example sending a request using CreateTransitGatewayRouteRequest.
//    req := client.CreateTransitGatewayRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateTransitGatewayRoute
func (c *Client) CreateTransitGatewayRouteRequest(input *types.CreateTransitGatewayRouteInput) CreateTransitGatewayRouteRequest {
	op := &aws.Operation{
		Name:       opCreateTransitGatewayRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateTransitGatewayRouteInput{}
	}

	req := c.newRequest(op, input, &types.CreateTransitGatewayRouteOutput{})
	return CreateTransitGatewayRouteRequest{Request: req, Input: input, Copy: c.CreateTransitGatewayRouteRequest}
}

// CreateTransitGatewayRouteRequest is the request type for the
// CreateTransitGatewayRoute API operation.
type CreateTransitGatewayRouteRequest struct {
	*aws.Request
	Input *types.CreateTransitGatewayRouteInput
	Copy  func(*types.CreateTransitGatewayRouteInput) CreateTransitGatewayRouteRequest
}

// Send marshals and sends the CreateTransitGatewayRoute API request.
func (r CreateTransitGatewayRouteRequest) Send(ctx context.Context) (*CreateTransitGatewayRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTransitGatewayRouteResponse{
		CreateTransitGatewayRouteOutput: r.Request.Data.(*types.CreateTransitGatewayRouteOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTransitGatewayRouteResponse is the response type for the
// CreateTransitGatewayRoute API operation.
type CreateTransitGatewayRouteResponse struct {
	*types.CreateTransitGatewayRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTransitGatewayRoute request.
func (r *CreateTransitGatewayRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
