// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opAdminGetUser = "AdminGetUser"

// AdminGetUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Gets the specified user by user name in a user pool as an administrator.
// Works on any user.
//
// Calling this action requires developer credentials.
//
//    // Example sending a request using AdminGetUserRequest.
//    req := client.AdminGetUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminGetUser
func (c *Client) AdminGetUserRequest(input *types.AdminGetUserInput) AdminGetUserRequest {
	op := &aws.Operation{
		Name:       opAdminGetUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AdminGetUserInput{}
	}

	req := c.newRequest(op, input, &types.AdminGetUserOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AdminGetUserMarshaler{Input: input}.GetNamedBuildHandler())

	return AdminGetUserRequest{Request: req, Input: input, Copy: c.AdminGetUserRequest}
}

// AdminGetUserRequest is the request type for the
// AdminGetUser API operation.
type AdminGetUserRequest struct {
	*aws.Request
	Input *types.AdminGetUserInput
	Copy  func(*types.AdminGetUserInput) AdminGetUserRequest
}

// Send marshals and sends the AdminGetUser API request.
func (r AdminGetUserRequest) Send(ctx context.Context) (*AdminGetUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminGetUserResponse{
		AdminGetUserOutput: r.Request.Data.(*types.AdminGetUserOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminGetUserResponse is the response type for the
// AdminGetUser API operation.
type AdminGetUserResponse struct {
	*types.AdminGetUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminGetUser request.
func (r *AdminGetUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
