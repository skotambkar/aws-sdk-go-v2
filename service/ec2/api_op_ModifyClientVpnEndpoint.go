// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opModifyClientVpnEndpoint = "ModifyClientVpnEndpoint"

// ModifyClientVpnEndpointRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the specified Client VPN endpoint. You can only modify an endpoint's
// server certificate information, client connection logging information, DNS
// server, and description. Modifying the DNS server resets existing client
// connections.
//
//    // Example sending a request using ModifyClientVpnEndpointRequest.
//    req := client.ModifyClientVpnEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyClientVpnEndpoint
func (c *Client) ModifyClientVpnEndpointRequest(input *types.ModifyClientVpnEndpointInput) ModifyClientVpnEndpointRequest {
	op := &aws.Operation{
		Name:       opModifyClientVpnEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyClientVpnEndpointInput{}
	}

	req := c.newRequest(op, input, &types.ModifyClientVpnEndpointOutput{})
	return ModifyClientVpnEndpointRequest{Request: req, Input: input, Copy: c.ModifyClientVpnEndpointRequest}
}

// ModifyClientVpnEndpointRequest is the request type for the
// ModifyClientVpnEndpoint API operation.
type ModifyClientVpnEndpointRequest struct {
	*aws.Request
	Input *types.ModifyClientVpnEndpointInput
	Copy  func(*types.ModifyClientVpnEndpointInput) ModifyClientVpnEndpointRequest
}

// Send marshals and sends the ModifyClientVpnEndpoint API request.
func (r ModifyClientVpnEndpointRequest) Send(ctx context.Context) (*ModifyClientVpnEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyClientVpnEndpointResponse{
		ModifyClientVpnEndpointOutput: r.Request.Data.(*types.ModifyClientVpnEndpointOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyClientVpnEndpointResponse is the response type for the
// ModifyClientVpnEndpoint API operation.
type ModifyClientVpnEndpointResponse struct {
	*types.ModifyClientVpnEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyClientVpnEndpoint request.
func (r *ModifyClientVpnEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
