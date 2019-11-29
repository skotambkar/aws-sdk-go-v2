// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
)

const opGetCertificateAuthorityCertificate = "GetCertificateAuthorityCertificate"

// GetCertificateAuthorityCertificateRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Retrieves the certificate and certificate chain for your private certificate
// authority (CA). Both the certificate and the chain are base64 PEM-encoded.
// The chain does not include the CA certificate. Each certificate in the chain
// signs the one before it.
//
//    // Example sending a request using GetCertificateAuthorityCertificateRequest.
//    req := client.GetCertificateAuthorityCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/GetCertificateAuthorityCertificate
func (c *Client) GetCertificateAuthorityCertificateRequest(input *types.GetCertificateAuthorityCertificateInput) GetCertificateAuthorityCertificateRequest {
	op := &aws.Operation{
		Name:       opGetCertificateAuthorityCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetCertificateAuthorityCertificateInput{}
	}

	req := c.newRequest(op, input, &types.GetCertificateAuthorityCertificateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetCertificateAuthorityCertificateMarshaler{Input: input}.GetNamedBuildHandler())

	return GetCertificateAuthorityCertificateRequest{Request: req, Input: input, Copy: c.GetCertificateAuthorityCertificateRequest}
}

// GetCertificateAuthorityCertificateRequest is the request type for the
// GetCertificateAuthorityCertificate API operation.
type GetCertificateAuthorityCertificateRequest struct {
	*aws.Request
	Input *types.GetCertificateAuthorityCertificateInput
	Copy  func(*types.GetCertificateAuthorityCertificateInput) GetCertificateAuthorityCertificateRequest
}

// Send marshals and sends the GetCertificateAuthorityCertificate API request.
func (r GetCertificateAuthorityCertificateRequest) Send(ctx context.Context) (*GetCertificateAuthorityCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCertificateAuthorityCertificateResponse{
		GetCertificateAuthorityCertificateOutput: r.Request.Data.(*types.GetCertificateAuthorityCertificateOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCertificateAuthorityCertificateResponse is the response type for the
// GetCertificateAuthorityCertificate API operation.
type GetCertificateAuthorityCertificateResponse struct {
	*types.GetCertificateAuthorityCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCertificateAuthorityCertificate request.
func (r *GetCertificateAuthorityCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
