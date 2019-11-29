// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opReplaceRoute = "ReplaceRoute"

// ReplaceRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Replaces an existing route within a route table in a VPC. You must provide
// only one of the following: internet gateway or virtual private gateway, NAT
// instance, NAT gateway, VPC peering connection, network interface, or egress-only
// internet gateway.
//
// For more information, see Route Tables (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using ReplaceRouteRequest.
//    req := client.ReplaceRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReplaceRoute
func (c *Client) ReplaceRouteRequest(input *types.ReplaceRouteInput) ReplaceRouteRequest {
	op := &aws.Operation{
		Name:       opReplaceRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ReplaceRouteInput{}
	}

	req := c.newRequest(op, input, &types.ReplaceRouteOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.ReplaceRouteMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return ReplaceRouteRequest{Request: req, Input: input, Copy: c.ReplaceRouteRequest}
}

// ReplaceRouteRequest is the request type for the
// ReplaceRoute API operation.
type ReplaceRouteRequest struct {
	*aws.Request
	Input *types.ReplaceRouteInput
	Copy  func(*types.ReplaceRouteInput) ReplaceRouteRequest
}

// Send marshals and sends the ReplaceRoute API request.
func (r ReplaceRouteRequest) Send(ctx context.Context) (*ReplaceRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReplaceRouteResponse{
		ReplaceRouteOutput: r.Request.Data.(*types.ReplaceRouteOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReplaceRouteResponse is the response type for the
// ReplaceRoute API operation.
type ReplaceRouteResponse struct {
	*types.ReplaceRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReplaceRoute request.
func (r *ReplaceRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
