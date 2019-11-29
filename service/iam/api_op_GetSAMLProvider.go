// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opGetSAMLProvider = "GetSAMLProvider"

// GetSAMLProviderRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Returns the SAML provider metadocument that was uploaded when the IAM SAML
// provider resource object was created or updated.
//
// This operation requires Signature Version 4 (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
//    // Example sending a request using GetSAMLProviderRequest.
//    req := client.GetSAMLProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetSAMLProvider
func (c *Client) GetSAMLProviderRequest(input *types.GetSAMLProviderInput) GetSAMLProviderRequest {
	op := &aws.Operation{
		Name:       opGetSAMLProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetSAMLProviderInput{}
	}

	req := c.newRequest(op, input, &types.GetSAMLProviderOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.GetSAMLProviderMarshaler{Input: input}.GetNamedBuildHandler())

	return GetSAMLProviderRequest{Request: req, Input: input, Copy: c.GetSAMLProviderRequest}
}

// GetSAMLProviderRequest is the request type for the
// GetSAMLProvider API operation.
type GetSAMLProviderRequest struct {
	*aws.Request
	Input *types.GetSAMLProviderInput
	Copy  func(*types.GetSAMLProviderInput) GetSAMLProviderRequest
}

// Send marshals and sends the GetSAMLProvider API request.
func (r GetSAMLProviderRequest) Send(ctx context.Context) (*GetSAMLProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSAMLProviderResponse{
		GetSAMLProviderOutput: r.Request.Data.(*types.GetSAMLProviderOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSAMLProviderResponse is the response type for the
// GetSAMLProvider API operation.
type GetSAMLProviderResponse struct {
	*types.GetSAMLProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSAMLProvider request.
func (r *GetSAMLProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
