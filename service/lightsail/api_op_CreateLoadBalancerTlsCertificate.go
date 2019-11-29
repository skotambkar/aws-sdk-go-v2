// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opCreateLoadBalancerTlsCertificate = "CreateLoadBalancerTlsCertificate"

// CreateLoadBalancerTlsCertificateRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Creates a Lightsail load balancer TLS certificate.
//
// TLS is just an updated, more secure version of Secure Socket Layer (SSL).
//
// The create load balancer tls certificate operation supports tag-based access
// control via resource tags applied to the resource identified by load balancer
// name. For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using CreateLoadBalancerTlsCertificateRequest.
//    req := client.CreateLoadBalancerTlsCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/CreateLoadBalancerTlsCertificate
func (c *Client) CreateLoadBalancerTlsCertificateRequest(input *types.CreateLoadBalancerTlsCertificateInput) CreateLoadBalancerTlsCertificateRequest {
	op := &aws.Operation{
		Name:       opCreateLoadBalancerTlsCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateLoadBalancerTlsCertificateInput{}
	}

	req := c.newRequest(op, input, &types.CreateLoadBalancerTlsCertificateOutput{})
	return CreateLoadBalancerTlsCertificateRequest{Request: req, Input: input, Copy: c.CreateLoadBalancerTlsCertificateRequest}
}

// CreateLoadBalancerTlsCertificateRequest is the request type for the
// CreateLoadBalancerTlsCertificate API operation.
type CreateLoadBalancerTlsCertificateRequest struct {
	*aws.Request
	Input *types.CreateLoadBalancerTlsCertificateInput
	Copy  func(*types.CreateLoadBalancerTlsCertificateInput) CreateLoadBalancerTlsCertificateRequest
}

// Send marshals and sends the CreateLoadBalancerTlsCertificate API request.
func (r CreateLoadBalancerTlsCertificateRequest) Send(ctx context.Context) (*CreateLoadBalancerTlsCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLoadBalancerTlsCertificateResponse{
		CreateLoadBalancerTlsCertificateOutput: r.Request.Data.(*types.CreateLoadBalancerTlsCertificateOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLoadBalancerTlsCertificateResponse is the response type for the
// CreateLoadBalancerTlsCertificate API operation.
type CreateLoadBalancerTlsCertificateResponse struct {
	*types.CreateLoadBalancerTlsCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLoadBalancerTlsCertificate request.
func (r *CreateLoadBalancerTlsCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
