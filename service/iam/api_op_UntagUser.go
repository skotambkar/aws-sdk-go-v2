// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opUntagUser = "UntagUser"

// UntagUserRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Removes the specified tags from the user. For more information about tagging,
// see Tagging IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
// in the IAM User Guide.
//
//    // Example sending a request using UntagUserRequest.
//    req := client.UntagUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UntagUser
func (c *Client) UntagUserRequest(input *types.UntagUserInput) UntagUserRequest {
	op := &aws.Operation{
		Name:       opUntagUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UntagUserInput{}
	}

	req := c.newRequest(op, input, &types.UntagUserOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UntagUserRequest{Request: req, Input: input, Copy: c.UntagUserRequest}
}

// UntagUserRequest is the request type for the
// UntagUser API operation.
type UntagUserRequest struct {
	*aws.Request
	Input *types.UntagUserInput
	Copy  func(*types.UntagUserInput) UntagUserRequest
}

// Send marshals and sends the UntagUser API request.
func (r UntagUserRequest) Send(ctx context.Context) (*UntagUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagUserResponse{
		UntagUserOutput: r.Request.Data.(*types.UntagUserOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagUserResponse is the response type for the
// UntagUser API operation.
type UntagUserResponse struct {
	*types.UntagUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagUser request.
func (r *UntagUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
