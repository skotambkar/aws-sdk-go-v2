// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opRevokeClientVpnIngress = "RevokeClientVpnIngress"

// RevokeClientVpnIngressRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Removes an ingress authorization rule from a Client VPN endpoint.
//
//    // Example sending a request using RevokeClientVpnIngressRequest.
//    req := client.RevokeClientVpnIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/RevokeClientVpnIngress
func (c *Client) RevokeClientVpnIngressRequest(input *types.RevokeClientVpnIngressInput) RevokeClientVpnIngressRequest {
	op := &aws.Operation{
		Name:       opRevokeClientVpnIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RevokeClientVpnIngressInput{}
	}

	req := c.newRequest(op, input, &types.RevokeClientVpnIngressOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.RevokeClientVpnIngressMarshaler{Input: input}.GetNamedBuildHandler())

	return RevokeClientVpnIngressRequest{Request: req, Input: input, Copy: c.RevokeClientVpnIngressRequest}
}

// RevokeClientVpnIngressRequest is the request type for the
// RevokeClientVpnIngress API operation.
type RevokeClientVpnIngressRequest struct {
	*aws.Request
	Input *types.RevokeClientVpnIngressInput
	Copy  func(*types.RevokeClientVpnIngressInput) RevokeClientVpnIngressRequest
}

// Send marshals and sends the RevokeClientVpnIngress API request.
func (r RevokeClientVpnIngressRequest) Send(ctx context.Context) (*RevokeClientVpnIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokeClientVpnIngressResponse{
		RevokeClientVpnIngressOutput: r.Request.Data.(*types.RevokeClientVpnIngressOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokeClientVpnIngressResponse is the response type for the
// RevokeClientVpnIngress API operation.
type RevokeClientVpnIngressResponse struct {
	*types.RevokeClientVpnIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokeClientVpnIngress request.
func (r *RevokeClientVpnIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
