// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opUpdateLoginProfile = "UpdateLoginProfile"

// UpdateLoginProfileRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Changes the password for the specified IAM user.
//
// IAM users can change their own passwords by calling ChangePassword. For more
// information about modifying passwords, see Managing Passwords (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html)
// in the IAM User Guide.
//
//    // Example sending a request using UpdateLoginProfileRequest.
//    req := client.UpdateLoginProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateLoginProfile
func (c *Client) UpdateLoginProfileRequest(input *types.UpdateLoginProfileInput) UpdateLoginProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateLoginProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateLoginProfileInput{}
	}

	req := c.newRequest(op, input, &types.UpdateLoginProfileOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateLoginProfileRequest{Request: req, Input: input, Copy: c.UpdateLoginProfileRequest}
}

// UpdateLoginProfileRequest is the request type for the
// UpdateLoginProfile API operation.
type UpdateLoginProfileRequest struct {
	*aws.Request
	Input *types.UpdateLoginProfileInput
	Copy  func(*types.UpdateLoginProfileInput) UpdateLoginProfileRequest
}

// Send marshals and sends the UpdateLoginProfile API request.
func (r UpdateLoginProfileRequest) Send(ctx context.Context) (*UpdateLoginProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLoginProfileResponse{
		UpdateLoginProfileOutput: r.Request.Data.(*types.UpdateLoginProfileOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLoginProfileResponse is the response type for the
// UpdateLoginProfile API operation.
type UpdateLoginProfileResponse struct {
	*types.UpdateLoginProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLoginProfile request.
func (r *UpdateLoginProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
