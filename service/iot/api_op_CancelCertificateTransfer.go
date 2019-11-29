// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opCancelCertificateTransfer = "CancelCertificateTransfer"

// CancelCertificateTransferRequest returns a request value for making API operation for
// AWS IoT.
//
// Cancels a pending transfer for the specified certificate.
//
// Note Only the transfer source account can use this operation to cancel a
// transfer. (Transfer destinations can use RejectCertificateTransfer instead.)
// After transfer, AWS IoT returns the certificate to the source account in
// the INACTIVE state. After the destination account has accepted the transfer,
// the transfer cannot be cancelled.
//
// After a certificate transfer is cancelled, the status of the certificate
// changes from PENDING_TRANSFER to INACTIVE.
//
//    // Example sending a request using CancelCertificateTransferRequest.
//    req := client.CancelCertificateTransferRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CancelCertificateTransferRequest(input *types.CancelCertificateTransferInput) CancelCertificateTransferRequest {
	op := &aws.Operation{
		Name:       opCancelCertificateTransfer,
		HTTPMethod: "PATCH",
		HTTPPath:   "/cancel-certificate-transfer/{certificateId}",
	}

	if input == nil {
		input = &types.CancelCertificateTransferInput{}
	}

	req := c.newRequest(op, input, &types.CancelCertificateTransferOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CancelCertificateTransferRequest{Request: req, Input: input, Copy: c.CancelCertificateTransferRequest}
}

// CancelCertificateTransferRequest is the request type for the
// CancelCertificateTransfer API operation.
type CancelCertificateTransferRequest struct {
	*aws.Request
	Input *types.CancelCertificateTransferInput
	Copy  func(*types.CancelCertificateTransferInput) CancelCertificateTransferRequest
}

// Send marshals and sends the CancelCertificateTransfer API request.
func (r CancelCertificateTransferRequest) Send(ctx context.Context) (*CancelCertificateTransferResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelCertificateTransferResponse{
		CancelCertificateTransferOutput: r.Request.Data.(*types.CancelCertificateTransferOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelCertificateTransferResponse is the response type for the
// CancelCertificateTransfer API operation.
type CancelCertificateTransferResponse struct {
	*types.CancelCertificateTransferOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelCertificateTransfer request.
func (r *CancelCertificateTransferResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
