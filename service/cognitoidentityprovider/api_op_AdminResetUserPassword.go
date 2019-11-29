// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opAdminResetUserPassword = "AdminResetUserPassword"

// AdminResetUserPasswordRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Resets the specified user's password in a user pool as an administrator.
// Works on any user.
//
// When a developer calls this API, the current password is invalidated, so
// it must be changed. If a user tries to sign in after the API is called, the
// app will get a PasswordResetRequiredException exception back and should direct
// the user down the flow to reset the password, which is the same as the forgot
// password flow. In addition, if the user pool has phone verification selected
// and a verified phone number exists for the user, or if email verification
// is selected and a verified email exists for the user, calling this API will
// also result in sending a message to the end user with the code to change
// their password.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using AdminResetUserPasswordRequest.
//    req := client.AdminResetUserPasswordRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminResetUserPassword
func (c *Client) AdminResetUserPasswordRequest(input *types.AdminResetUserPasswordInput) AdminResetUserPasswordRequest {
	op := &aws.Operation{
		Name:       opAdminResetUserPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AdminResetUserPasswordInput{}
	}

	req := c.newRequest(op, input, &types.AdminResetUserPasswordOutput{})
	return AdminResetUserPasswordRequest{Request: req, Input: input, Copy: c.AdminResetUserPasswordRequest}
}

// AdminResetUserPasswordRequest is the request type for the
// AdminResetUserPassword API operation.
type AdminResetUserPasswordRequest struct {
	*aws.Request
	Input *types.AdminResetUserPasswordInput
	Copy  func(*types.AdminResetUserPasswordInput) AdminResetUserPasswordRequest
}

// Send marshals and sends the AdminResetUserPassword API request.
func (r AdminResetUserPasswordRequest) Send(ctx context.Context) (*AdminResetUserPasswordResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminResetUserPasswordResponse{
		AdminResetUserPasswordOutput: r.Request.Data.(*types.AdminResetUserPasswordOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminResetUserPasswordResponse is the response type for the
// AdminResetUserPassword API operation.
type AdminResetUserPasswordResponse struct {
	*types.AdminResetUserPasswordOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminResetUserPassword request.
func (r *AdminResetUserPasswordResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
