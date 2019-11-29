// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opCreateKeysAndCertificate = "CreateKeysAndCertificate"

// CreateKeysAndCertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a 2048-bit RSA key pair and issues an X.509 certificate using the
// issued public key.
//
// Note This is the only time AWS IoT issues the private key for this certificate,
// so it is important to keep it in a secure location.
//
//    // Example sending a request using CreateKeysAndCertificateRequest.
//    req := client.CreateKeysAndCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateKeysAndCertificateRequest(input *types.CreateKeysAndCertificateInput) CreateKeysAndCertificateRequest {
	op := &aws.Operation{
		Name:       opCreateKeysAndCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/keys-and-certificate",
	}

	if input == nil {
		input = &types.CreateKeysAndCertificateInput{}
	}

	req := c.newRequest(op, input, &types.CreateKeysAndCertificateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateKeysAndCertificateMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateKeysAndCertificateRequest{Request: req, Input: input, Copy: c.CreateKeysAndCertificateRequest}
}

// CreateKeysAndCertificateRequest is the request type for the
// CreateKeysAndCertificate API operation.
type CreateKeysAndCertificateRequest struct {
	*aws.Request
	Input *types.CreateKeysAndCertificateInput
	Copy  func(*types.CreateKeysAndCertificateInput) CreateKeysAndCertificateRequest
}

// Send marshals and sends the CreateKeysAndCertificate API request.
func (r CreateKeysAndCertificateRequest) Send(ctx context.Context) (*CreateKeysAndCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateKeysAndCertificateResponse{
		CreateKeysAndCertificateOutput: r.Request.Data.(*types.CreateKeysAndCertificateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateKeysAndCertificateResponse is the response type for the
// CreateKeysAndCertificate API operation.
type CreateKeysAndCertificateResponse struct {
	*types.CreateKeysAndCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateKeysAndCertificate request.
func (r *CreateKeysAndCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
