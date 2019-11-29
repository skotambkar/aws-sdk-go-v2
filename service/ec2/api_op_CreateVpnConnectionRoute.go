// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateVpnConnectionRoute = "CreateVpnConnectionRoute"

// CreateVpnConnectionRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a static route associated with a VPN connection between an existing
// virtual private gateway and a VPN customer gateway. The static route allows
// traffic to be routed from the virtual private gateway to the VPN customer
// gateway.
//
// For more information, see AWS Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html)
// in the AWS Site-to-Site VPN User Guide.
//
//    // Example sending a request using CreateVpnConnectionRouteRequest.
//    req := client.CreateVpnConnectionRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateVpnConnectionRoute
func (c *Client) CreateVpnConnectionRouteRequest(input *types.CreateVpnConnectionRouteInput) CreateVpnConnectionRouteRequest {
	op := &aws.Operation{
		Name:       opCreateVpnConnectionRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateVpnConnectionRouteInput{}
	}

	req := c.newRequest(op, input, &types.CreateVpnConnectionRouteOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateVpnConnectionRouteRequest{Request: req, Input: input, Copy: c.CreateVpnConnectionRouteRequest}
}

// CreateVpnConnectionRouteRequest is the request type for the
// CreateVpnConnectionRoute API operation.
type CreateVpnConnectionRouteRequest struct {
	*aws.Request
	Input *types.CreateVpnConnectionRouteInput
	Copy  func(*types.CreateVpnConnectionRouteInput) CreateVpnConnectionRouteRequest
}

// Send marshals and sends the CreateVpnConnectionRoute API request.
func (r CreateVpnConnectionRouteRequest) Send(ctx context.Context) (*CreateVpnConnectionRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVpnConnectionRouteResponse{
		CreateVpnConnectionRouteOutput: r.Request.Data.(*types.CreateVpnConnectionRouteOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVpnConnectionRouteResponse is the response type for the
// CreateVpnConnectionRoute API operation.
type CreateVpnConnectionRouteResponse struct {
	*types.CreateVpnConnectionRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVpnConnectionRoute request.
func (r *CreateVpnConnectionRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
