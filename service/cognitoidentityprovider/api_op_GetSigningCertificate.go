// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opGetSigningCertificate = "GetSigningCertificate"

// GetSigningCertificateRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// This method takes a user pool ID, and returns the signing certificate.
//
//    // Example sending a request using GetSigningCertificateRequest.
//    req := client.GetSigningCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetSigningCertificate
func (c *Client) GetSigningCertificateRequest(input *types.GetSigningCertificateInput) GetSigningCertificateRequest {
	op := &aws.Operation{
		Name:       opGetSigningCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetSigningCertificateInput{}
	}

	req := c.newRequest(op, input, &types.GetSigningCertificateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetSigningCertificateMarshaler{Input: input}.GetNamedBuildHandler())

	return GetSigningCertificateRequest{Request: req, Input: input, Copy: c.GetSigningCertificateRequest}
}

// GetSigningCertificateRequest is the request type for the
// GetSigningCertificate API operation.
type GetSigningCertificateRequest struct {
	*aws.Request
	Input *types.GetSigningCertificateInput
	Copy  func(*types.GetSigningCertificateInput) GetSigningCertificateRequest
}

// Send marshals and sends the GetSigningCertificate API request.
func (r GetSigningCertificateRequest) Send(ctx context.Context) (*GetSigningCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSigningCertificateResponse{
		GetSigningCertificateOutput: r.Request.Data.(*types.GetSigningCertificateOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSigningCertificateResponse is the response type for the
// GetSigningCertificate API operation.
type GetSigningCertificateResponse struct {
	*types.GetSigningCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSigningCertificate request.
func (r *GetSigningCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
