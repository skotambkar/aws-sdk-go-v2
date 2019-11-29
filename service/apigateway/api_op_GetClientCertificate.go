// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetClientCertificate = "GetClientCertificate"

// GetClientCertificateRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets information about the current ClientCertificate resource.
//
//    // Example sending a request using GetClientCertificateRequest.
//    req := client.GetClientCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetClientCertificateRequest(input *types.GetClientCertificateInput) GetClientCertificateRequest {
	op := &aws.Operation{
		Name:       opGetClientCertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/clientcertificates/{clientcertificate_id}",
	}

	if input == nil {
		input = &types.GetClientCertificateInput{}
	}

	req := c.newRequest(op, input, &types.GetClientCertificateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetClientCertificateMarshaler{Input: input}.GetNamedBuildHandler())

	return GetClientCertificateRequest{Request: req, Input: input, Copy: c.GetClientCertificateRequest}
}

// GetClientCertificateRequest is the request type for the
// GetClientCertificate API operation.
type GetClientCertificateRequest struct {
	*aws.Request
	Input *types.GetClientCertificateInput
	Copy  func(*types.GetClientCertificateInput) GetClientCertificateRequest
}

// Send marshals and sends the GetClientCertificate API request.
func (r GetClientCertificateRequest) Send(ctx context.Context) (*GetClientCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetClientCertificateResponse{
		GetClientCertificateOutput: r.Request.Data.(*types.GetClientCertificateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetClientCertificateResponse is the response type for the
// GetClientCertificate API operation.
type GetClientCertificateResponse struct {
	*types.GetClientCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetClientCertificate request.
func (r *GetClientCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
