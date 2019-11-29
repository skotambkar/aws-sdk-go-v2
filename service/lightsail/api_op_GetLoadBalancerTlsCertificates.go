// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetLoadBalancerTlsCertificates = "GetLoadBalancerTlsCertificates"

// GetLoadBalancerTlsCertificatesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about the TLS certificates that are associated with the
// specified Lightsail load balancer.
//
// TLS is just an updated, more secure version of Secure Socket Layer (SSL).
//
// You can have a maximum of 2 certificates associated with a Lightsail load
// balancer. One is active and the other is inactive.
//
//    // Example sending a request using GetLoadBalancerTlsCertificatesRequest.
//    req := client.GetLoadBalancerTlsCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetLoadBalancerTlsCertificates
func (c *Client) GetLoadBalancerTlsCertificatesRequest(input *types.GetLoadBalancerTlsCertificatesInput) GetLoadBalancerTlsCertificatesRequest {
	op := &aws.Operation{
		Name:       opGetLoadBalancerTlsCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetLoadBalancerTlsCertificatesInput{}
	}

	req := c.newRequest(op, input, &types.GetLoadBalancerTlsCertificatesOutput{})
	return GetLoadBalancerTlsCertificatesRequest{Request: req, Input: input, Copy: c.GetLoadBalancerTlsCertificatesRequest}
}

// GetLoadBalancerTlsCertificatesRequest is the request type for the
// GetLoadBalancerTlsCertificates API operation.
type GetLoadBalancerTlsCertificatesRequest struct {
	*aws.Request
	Input *types.GetLoadBalancerTlsCertificatesInput
	Copy  func(*types.GetLoadBalancerTlsCertificatesInput) GetLoadBalancerTlsCertificatesRequest
}

// Send marshals and sends the GetLoadBalancerTlsCertificates API request.
func (r GetLoadBalancerTlsCertificatesRequest) Send(ctx context.Context) (*GetLoadBalancerTlsCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoadBalancerTlsCertificatesResponse{
		GetLoadBalancerTlsCertificatesOutput: r.Request.Data.(*types.GetLoadBalancerTlsCertificatesOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoadBalancerTlsCertificatesResponse is the response type for the
// GetLoadBalancerTlsCertificates API operation.
type GetLoadBalancerTlsCertificatesResponse struct {
	*types.GetLoadBalancerTlsCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoadBalancerTlsCertificates request.
func (r *GetLoadBalancerTlsCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
