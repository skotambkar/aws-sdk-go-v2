// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateRoute = "CreateRoute"

// CreateRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a route in a route table within a VPC.
//
// You must specify one of the following targets: internet gateway or virtual
// private gateway, NAT instance, NAT gateway, VPC peering connection, network
// interface, or egress-only internet gateway.
//
// When determining how to route traffic, we use the route with the most specific
// match. For example, traffic is destined for the IPv4 address 192.0.2.3, and
// the route table includes the following two IPv4 routes:
//
//    * 192.0.2.0/24 (goes to some target A)
//
//    * 192.0.2.0/28 (goes to some target B)
//
// Both routes apply to the traffic destined for 192.0.2.3. However, the second
// route in the list covers a smaller number of IP addresses and is therefore
// more specific, so we use that route to determine where to target the traffic.
//
// For more information about route tables, see Route Tables (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateRouteRequest.
//    req := client.CreateRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateRoute
func (c *Client) CreateRouteRequest(input *types.CreateRouteInput) CreateRouteRequest {
	op := &aws.Operation{
		Name:       opCreateRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateRouteInput{}
	}

	req := c.newRequest(op, input, &types.CreateRouteOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.CreateRouteMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateRouteRequest{Request: req, Input: input, Copy: c.CreateRouteRequest}
}

// CreateRouteRequest is the request type for the
// CreateRoute API operation.
type CreateRouteRequest struct {
	*aws.Request
	Input *types.CreateRouteInput
	Copy  func(*types.CreateRouteInput) CreateRouteRequest
}

// Send marshals and sends the CreateRoute API request.
func (r CreateRouteRequest) Send(ctx context.Context) (*CreateRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRouteResponse{
		CreateRouteOutput: r.Request.Data.(*types.CreateRouteOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRouteResponse is the response type for the
// CreateRoute API operation.
type CreateRouteResponse struct {
	*types.CreateRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRoute request.
func (r *CreateRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
