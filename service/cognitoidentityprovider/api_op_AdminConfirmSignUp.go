// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opAdminConfirmSignUp = "AdminConfirmSignUp"

// AdminConfirmSignUpRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Confirms user registration as an admin without using a confirmation code.
// Works on any user.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using AdminConfirmSignUpRequest.
//    req := client.AdminConfirmSignUpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminConfirmSignUp
func (c *Client) AdminConfirmSignUpRequest(input *types.AdminConfirmSignUpInput) AdminConfirmSignUpRequest {
	op := &aws.Operation{
		Name:       opAdminConfirmSignUp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AdminConfirmSignUpInput{}
	}

	req := c.newRequest(op, input, &types.AdminConfirmSignUpOutput{})
	return AdminConfirmSignUpRequest{Request: req, Input: input, Copy: c.AdminConfirmSignUpRequest}
}

// AdminConfirmSignUpRequest is the request type for the
// AdminConfirmSignUp API operation.
type AdminConfirmSignUpRequest struct {
	*aws.Request
	Input *types.AdminConfirmSignUpInput
	Copy  func(*types.AdminConfirmSignUpInput) AdminConfirmSignUpRequest
}

// Send marshals and sends the AdminConfirmSignUp API request.
func (r AdminConfirmSignUpRequest) Send(ctx context.Context) (*AdminConfirmSignUpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminConfirmSignUpResponse{
		AdminConfirmSignUpOutput: r.Request.Data.(*types.AdminConfirmSignUpOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminConfirmSignUpResponse is the response type for the
// AdminConfirmSignUp API operation.
type AdminConfirmSignUpResponse struct {
	*types.AdminConfirmSignUpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminConfirmSignUp request.
func (r *AdminConfirmSignUpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
