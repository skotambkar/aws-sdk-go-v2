// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opAcceptCertificateTransfer = "AcceptCertificateTransfer"

// AcceptCertificateTransferRequest returns a request value for making API operation for
// AWS IoT.
//
// Accepts a pending certificate transfer. The default state of the certificate
// is INACTIVE.
//
// To check for pending certificate transfers, call ListCertificates to enumerate
// your certificates.
//
//    // Example sending a request using AcceptCertificateTransferRequest.
//    req := client.AcceptCertificateTransferRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) AcceptCertificateTransferRequest(input *types.AcceptCertificateTransferInput) AcceptCertificateTransferRequest {
	op := &aws.Operation{
		Name:       opAcceptCertificateTransfer,
		HTTPMethod: "PATCH",
		HTTPPath:   "/accept-certificate-transfer/{certificateId}",
	}

	if input == nil {
		input = &types.AcceptCertificateTransferInput{}
	}

	req := c.newRequest(op, input, &types.AcceptCertificateTransferOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.AcceptCertificateTransferMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AcceptCertificateTransferRequest{Request: req, Input: input, Copy: c.AcceptCertificateTransferRequest}
}

// AcceptCertificateTransferRequest is the request type for the
// AcceptCertificateTransfer API operation.
type AcceptCertificateTransferRequest struct {
	*aws.Request
	Input *types.AcceptCertificateTransferInput
	Copy  func(*types.AcceptCertificateTransferInput) AcceptCertificateTransferRequest
}

// Send marshals and sends the AcceptCertificateTransfer API request.
func (r AcceptCertificateTransferRequest) Send(ctx context.Context) (*AcceptCertificateTransferResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptCertificateTransferResponse{
		AcceptCertificateTransferOutput: r.Request.Data.(*types.AcceptCertificateTransferOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptCertificateTransferResponse is the response type for the
// AcceptCertificateTransfer API operation.
type AcceptCertificateTransferResponse struct {
	*types.AcceptCertificateTransferOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptCertificateTransfer request.
func (r *AcceptCertificateTransferResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
