// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opRemoveClientIDFromOpenIDConnectProvider = "RemoveClientIDFromOpenIDConnectProvider"

// RemoveClientIDFromOpenIDConnectProviderRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Removes the specified client ID (also known as audience) from the list of
// client IDs registered for the specified IAM OpenID Connect (OIDC) provider
// resource object.
//
// This operation is idempotent; it does not fail or return an error if you
// try to remove a client ID that does not exist.
//
//    // Example sending a request using RemoveClientIDFromOpenIDConnectProviderRequest.
//    req := client.RemoveClientIDFromOpenIDConnectProviderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/RemoveClientIDFromOpenIDConnectProvider
func (c *Client) RemoveClientIDFromOpenIDConnectProviderRequest(input *types.RemoveClientIDFromOpenIDConnectProviderInput) RemoveClientIDFromOpenIDConnectProviderRequest {
	op := &aws.Operation{
		Name:       opRemoveClientIDFromOpenIDConnectProvider,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RemoveClientIDFromOpenIDConnectProviderInput{}
	}

	req := c.newRequest(op, input, &types.RemoveClientIDFromOpenIDConnectProviderOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveClientIDFromOpenIDConnectProviderRequest{Request: req, Input: input, Copy: c.RemoveClientIDFromOpenIDConnectProviderRequest}
}

// RemoveClientIDFromOpenIDConnectProviderRequest is the request type for the
// RemoveClientIDFromOpenIDConnectProvider API operation.
type RemoveClientIDFromOpenIDConnectProviderRequest struct {
	*aws.Request
	Input *types.RemoveClientIDFromOpenIDConnectProviderInput
	Copy  func(*types.RemoveClientIDFromOpenIDConnectProviderInput) RemoveClientIDFromOpenIDConnectProviderRequest
}

// Send marshals and sends the RemoveClientIDFromOpenIDConnectProvider API request.
func (r RemoveClientIDFromOpenIDConnectProviderRequest) Send(ctx context.Context) (*RemoveClientIDFromOpenIDConnectProviderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveClientIDFromOpenIDConnectProviderResponse{
		RemoveClientIDFromOpenIDConnectProviderOutput: r.Request.Data.(*types.RemoveClientIDFromOpenIDConnectProviderOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveClientIDFromOpenIDConnectProviderResponse is the response type for the
// RemoveClientIDFromOpenIDConnectProvider API operation.
type RemoveClientIDFromOpenIDConnectProviderResponse struct {
	*types.RemoveClientIDFromOpenIDConnectProviderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveClientIDFromOpenIDConnectProvider request.
func (r *RemoveClientIDFromOpenIDConnectProviderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
