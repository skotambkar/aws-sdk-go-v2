// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opTransferCertificate = "TransferCertificate"

// TransferCertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Transfers the specified certificate to the specified AWS account.
//
// You can cancel the transfer until it is acknowledged by the recipient.
//
// No notification is sent to the transfer destination's account. It is up to
// the caller to notify the transfer target.
//
// The certificate being transferred must not be in the ACTIVE state. You can
// use the UpdateCertificate API to deactivate it.
//
// The certificate must not have any policies attached to it. You can use the
// DetachPrincipalPolicy API to detach them.
//
//    // Example sending a request using TransferCertificateRequest.
//    req := client.TransferCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) TransferCertificateRequest(input *types.TransferCertificateInput) TransferCertificateRequest {
	op := &aws.Operation{
		Name:       opTransferCertificate,
		HTTPMethod: "PATCH",
		HTTPPath:   "/transfer-certificate/{certificateId}",
	}

	if input == nil {
		input = &types.TransferCertificateInput{}
	}

	req := c.newRequest(op, input, &types.TransferCertificateOutput{})
	return TransferCertificateRequest{Request: req, Input: input, Copy: c.TransferCertificateRequest}
}

// TransferCertificateRequest is the request type for the
// TransferCertificate API operation.
type TransferCertificateRequest struct {
	*aws.Request
	Input *types.TransferCertificateInput
	Copy  func(*types.TransferCertificateInput) TransferCertificateRequest
}

// Send marshals and sends the TransferCertificate API request.
func (r TransferCertificateRequest) Send(ctx context.Context) (*TransferCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TransferCertificateResponse{
		TransferCertificateOutput: r.Request.Data.(*types.TransferCertificateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TransferCertificateResponse is the response type for the
// TransferCertificate API operation.
type TransferCertificateResponse struct {
	*types.TransferCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TransferCertificate request.
func (r *TransferCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
