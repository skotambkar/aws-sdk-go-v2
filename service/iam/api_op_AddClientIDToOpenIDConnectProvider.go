// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opAddClientIDToOpenIDConnectProvider = "AddClientIDToOpenIDConnectProvider"

// AddClientIDToOpenIDConnectProviderRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Adds a new client ID (also known as audience) to the list of client IDs already
// registered for the specified IAM OpenID Connect (OIDC) provider resource.
//
// This operation is idempotent; it does not fail or return an error if you
// add an existing client ID to the provider.
//
//    // Example sending a request using AddClientIDToOpenIDConnectProviderRequest.
//    req := client.AddClientIDToOpenIDConnectProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/AddClientIDToOpenIDConnectProvider
func (c *Client) AddClientIDToOpenIDConnectProviderRequest(input *types.AddClientIDToOpenIDConnectProviderInput) AddClientIDToOpenIDConnectProviderRequest {
	op := &aws.Operation{
		Name:       opAddClientIDToOpenIDConnectProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AddClientIDToOpenIDConnectProviderInput{}
	}

	req := c.newRequest(op, input, &types.AddClientIDToOpenIDConnectProviderOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AddClientIDToOpenIDConnectProviderRequest{Request: req, Input: input, Copy: c.AddClientIDToOpenIDConnectProviderRequest}
}

// AddClientIDToOpenIDConnectProviderRequest is the request type for the
// AddClientIDToOpenIDConnectProvider API operation.
type AddClientIDToOpenIDConnectProviderRequest struct {
	*aws.Request
	Input *types.AddClientIDToOpenIDConnectProviderInput
	Copy  func(*types.AddClientIDToOpenIDConnectProviderInput) AddClientIDToOpenIDConnectProviderRequest
}

// Send marshals and sends the AddClientIDToOpenIDConnectProvider API request.
func (r AddClientIDToOpenIDConnectProviderRequest) Send(ctx context.Context) (*AddClientIDToOpenIDConnectProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddClientIDToOpenIDConnectProviderResponse{
		AddClientIDToOpenIDConnectProviderOutput: r.Request.Data.(*types.AddClientIDToOpenIDConnectProviderOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddClientIDToOpenIDConnectProviderResponse is the response type for the
// AddClientIDToOpenIDConnectProvider API operation.
type AddClientIDToOpenIDConnectProviderResponse struct {
	*types.AddClientIDToOpenIDConnectProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddClientIDToOpenIDConnectProvider request.
func (r *AddClientIDToOpenIDConnectProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
