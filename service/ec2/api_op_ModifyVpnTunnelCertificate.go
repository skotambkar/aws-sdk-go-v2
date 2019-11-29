// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opModifyVpnTunnelCertificate = "ModifyVpnTunnelCertificate"

// ModifyVpnTunnelCertificateRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the VPN tunnel endpoint certificate.
//
//    // Example sending a request using ModifyVpnTunnelCertificateRequest.
//    req := client.ModifyVpnTunnelCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpnTunnelCertificate
func (c *Client) ModifyVpnTunnelCertificateRequest(input *types.ModifyVpnTunnelCertificateInput) ModifyVpnTunnelCertificateRequest {
	op := &aws.Operation{
		Name:       opModifyVpnTunnelCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyVpnTunnelCertificateInput{}
	}

	req := c.newRequest(op, input, &types.ModifyVpnTunnelCertificateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.ModifyVpnTunnelCertificateMarshaler{Input: input}.GetNamedBuildHandler())

	return ModifyVpnTunnelCertificateRequest{Request: req, Input: input, Copy: c.ModifyVpnTunnelCertificateRequest}
}

// ModifyVpnTunnelCertificateRequest is the request type for the
// ModifyVpnTunnelCertificate API operation.
type ModifyVpnTunnelCertificateRequest struct {
	*aws.Request
	Input *types.ModifyVpnTunnelCertificateInput
	Copy  func(*types.ModifyVpnTunnelCertificateInput) ModifyVpnTunnelCertificateRequest
}

// Send marshals and sends the ModifyVpnTunnelCertificate API request.
func (r ModifyVpnTunnelCertificateRequest) Send(ctx context.Context) (*ModifyVpnTunnelCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpnTunnelCertificateResponse{
		ModifyVpnTunnelCertificateOutput: r.Request.Data.(*types.ModifyVpnTunnelCertificateOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpnTunnelCertificateResponse is the response type for the
// ModifyVpnTunnelCertificate API operation.
type ModifyVpnTunnelCertificateResponse struct {
	*types.ModifyVpnTunnelCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpnTunnelCertificate request.
func (r *ModifyVpnTunnelCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
