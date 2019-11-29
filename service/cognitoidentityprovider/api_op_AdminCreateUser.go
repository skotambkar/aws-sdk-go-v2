// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opAdminCreateUser = "AdminCreateUser"

// AdminCreateUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Creates a new user in the specified user pool.
//
// If MessageAction is not set, the default is to send a welcome message via
// email or phone (SMS).
//
// This message is based on a template that you configured in your call to or
// . This template includes your custom sign-up instructions and placeholders
// for user name and temporary password.
//
// Alternatively, you can call AdminCreateUser with “SUPPRESS” for the MessageAction
// parameter, and Amazon Cognito will not send any email.
//
// In either case, the user will be in the FORCE_CHANGE_PASSWORD state until
// they sign in and change their password.
//
// AdminCreateUser requires developer credentials.
//
//    // Example sending a request using AdminCreateUserRequest.
//    req := client.AdminCreateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminCreateUser
func (c *Client) AdminCreateUserRequest(input *types.AdminCreateUserInput) AdminCreateUserRequest {
	op := &aws.Operation{
		Name:       opAdminCreateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AdminCreateUserInput{}
	}

	req := c.newRequest(op, input, &types.AdminCreateUserOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AdminCreateUserMarshaler{Input: input}.GetNamedBuildHandler())

	return AdminCreateUserRequest{Request: req, Input: input, Copy: c.AdminCreateUserRequest}
}

// AdminCreateUserRequest is the request type for the
// AdminCreateUser API operation.
type AdminCreateUserRequest struct {
	*aws.Request
	Input *types.AdminCreateUserInput
	Copy  func(*types.AdminCreateUserInput) AdminCreateUserRequest
}

// Send marshals and sends the AdminCreateUser API request.
func (r AdminCreateUserRequest) Send(ctx context.Context) (*AdminCreateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminCreateUserResponse{
		AdminCreateUserOutput: r.Request.Data.(*types.AdminCreateUserOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminCreateUserResponse is the response type for the
// AdminCreateUser API operation.
type AdminCreateUserResponse struct {
	*types.AdminCreateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminCreateUser request.
func (r *AdminCreateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
