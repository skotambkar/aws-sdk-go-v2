// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opAcceptVpcEndpointConnections = "AcceptVpcEndpointConnections"

// AcceptVpcEndpointConnectionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Accepts one or more interface VPC endpoint connection requests to your VPC
// endpoint service.
//
//    // Example sending a request using AcceptVpcEndpointConnectionsRequest.
//    req := client.AcceptVpcEndpointConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AcceptVpcEndpointConnections
func (c *Client) AcceptVpcEndpointConnectionsRequest(input *types.AcceptVpcEndpointConnectionsInput) AcceptVpcEndpointConnectionsRequest {
	op := &aws.Operation{
		Name:       opAcceptVpcEndpointConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AcceptVpcEndpointConnectionsInput{}
	}

	req := c.newRequest(op, input, &types.AcceptVpcEndpointConnectionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.AcceptVpcEndpointConnectionsMarshaler{Input: input}.GetNamedBuildHandler())

	return AcceptVpcEndpointConnectionsRequest{Request: req, Input: input, Copy: c.AcceptVpcEndpointConnectionsRequest}
}

// AcceptVpcEndpointConnectionsRequest is the request type for the
// AcceptVpcEndpointConnections API operation.
type AcceptVpcEndpointConnectionsRequest struct {
	*aws.Request
	Input *types.AcceptVpcEndpointConnectionsInput
	Copy  func(*types.AcceptVpcEndpointConnectionsInput) AcceptVpcEndpointConnectionsRequest
}

// Send marshals and sends the AcceptVpcEndpointConnections API request.
func (r AcceptVpcEndpointConnectionsRequest) Send(ctx context.Context) (*AcceptVpcEndpointConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptVpcEndpointConnectionsResponse{
		AcceptVpcEndpointConnectionsOutput: r.Request.Data.(*types.AcceptVpcEndpointConnectionsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptVpcEndpointConnectionsResponse is the response type for the
// AcceptVpcEndpointConnections API operation.
type AcceptVpcEndpointConnectionsResponse struct {
	*types.AcceptVpcEndpointConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptVpcEndpointConnections request.
func (r *AcceptVpcEndpointConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
