// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opDescribeCACertificate = "DescribeCACertificate"

// DescribeCACertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Describes a registered CA certificate.
//
//    // Example sending a request using DescribeCACertificateRequest.
//    req := client.DescribeCACertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeCACertificateRequest(input *types.DescribeCACertificateInput) DescribeCACertificateRequest {
	op := &aws.Operation{
		Name:       opDescribeCACertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/cacertificate/{caCertificateId}",
	}

	if input == nil {
		input = &types.DescribeCACertificateInput{}
	}

	req := c.newRequest(op, input, &types.DescribeCACertificateOutput{})
	return DescribeCACertificateRequest{Request: req, Input: input, Copy: c.DescribeCACertificateRequest}
}

// DescribeCACertificateRequest is the request type for the
// DescribeCACertificate API operation.
type DescribeCACertificateRequest struct {
	*aws.Request
	Input *types.DescribeCACertificateInput
	Copy  func(*types.DescribeCACertificateInput) DescribeCACertificateRequest
}

// Send marshals and sends the DescribeCACertificate API request.
func (r DescribeCACertificateRequest) Send(ctx context.Context) (*DescribeCACertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCACertificateResponse{
		DescribeCACertificateOutput: r.Request.Data.(*types.DescribeCACertificateOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCACertificateResponse is the response type for the
// DescribeCACertificate API operation.
type DescribeCACertificateResponse struct {
	*types.DescribeCACertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCACertificate request.
func (r *DescribeCACertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
