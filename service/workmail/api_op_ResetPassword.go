// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
)

const opResetPassword = "ResetPassword"

// ResetPasswordRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Allows the administrator to reset the password for a user.
//
//    // Example sending a request using ResetPasswordRequest.
//    req := client.ResetPasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/ResetPassword
func (c *Client) ResetPasswordRequest(input *types.ResetPasswordInput) ResetPasswordRequest {
	op := &aws.Operation{
		Name:       opResetPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResetPasswordInput{}
	}

	req := c.newRequest(op, input, &types.ResetPasswordOutput{})
	return ResetPasswordRequest{Request: req, Input: input, Copy: c.ResetPasswordRequest}
}

// ResetPasswordRequest is the request type for the
// ResetPassword API operation.
type ResetPasswordRequest struct {
	*aws.Request
	Input *types.ResetPasswordInput
	Copy  func(*types.ResetPasswordInput) ResetPasswordRequest
}

// Send marshals and sends the ResetPassword API request.
func (r ResetPasswordRequest) Send(ctx context.Context) (*ResetPasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetPasswordResponse{
		ResetPasswordOutput: r.Request.Data.(*types.ResetPasswordOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetPasswordResponse is the response type for the
// ResetPassword API operation.
type ResetPasswordResponse struct {
	*types.ResetPasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetPassword request.
func (r *ResetPasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
