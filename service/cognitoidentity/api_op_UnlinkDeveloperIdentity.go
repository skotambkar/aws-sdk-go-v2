// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
)

const opUnlinkDeveloperIdentity = "UnlinkDeveloperIdentity"

// UnlinkDeveloperIdentityRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Unlinks a DeveloperUserIdentifier from an existing identity. Unlinked developer
// users will be considered new identities next time they are seen. If, for
// a given Cognito identity, you remove all federated identities as well as
// the developer user identifier, the Cognito identity becomes inaccessible.
//
// You must use AWS Developer credentials to call this API.
//
//    // Example sending a request using UnlinkDeveloperIdentityRequest.
//    req := client.UnlinkDeveloperIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/UnlinkDeveloperIdentity
func (c *Client) UnlinkDeveloperIdentityRequest(input *types.UnlinkDeveloperIdentityInput) UnlinkDeveloperIdentityRequest {
	op := &aws.Operation{
		Name:       opUnlinkDeveloperIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UnlinkDeveloperIdentityInput{}
	}

	req := c.newRequest(op, input, &types.UnlinkDeveloperIdentityOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UnlinkDeveloperIdentityRequest{Request: req, Input: input, Copy: c.UnlinkDeveloperIdentityRequest}
}

// UnlinkDeveloperIdentityRequest is the request type for the
// UnlinkDeveloperIdentity API operation.
type UnlinkDeveloperIdentityRequest struct {
	*aws.Request
	Input *types.UnlinkDeveloperIdentityInput
	Copy  func(*types.UnlinkDeveloperIdentityInput) UnlinkDeveloperIdentityRequest
}

// Send marshals and sends the UnlinkDeveloperIdentity API request.
func (r UnlinkDeveloperIdentityRequest) Send(ctx context.Context) (*UnlinkDeveloperIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UnlinkDeveloperIdentityResponse{
		UnlinkDeveloperIdentityOutput: r.Request.Data.(*types.UnlinkDeveloperIdentityOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UnlinkDeveloperIdentityResponse is the response type for the
// UnlinkDeveloperIdentity API operation.
type UnlinkDeveloperIdentityResponse struct {
	*types.UnlinkDeveloperIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UnlinkDeveloperIdentity request.
func (r *UnlinkDeveloperIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
