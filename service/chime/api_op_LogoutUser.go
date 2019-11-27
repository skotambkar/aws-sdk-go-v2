// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opLogoutUser = "LogoutUser"

// LogoutUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Logs out the specified user from all of the devices they are currently logged
// into.
//
//    // Example sending a request using LogoutUserRequest.
//    req := client.LogoutUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/LogoutUser
func (c *Client) LogoutUserRequest(input *types.LogoutUserInput) LogoutUserRequest {
	op := &aws.Operation{
		Name:       opLogoutUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users/{userId}?operation=logout",
	}

	if input == nil {
		input = &types.LogoutUserInput{}
	}

	req := c.newRequest(op, input, &types.LogoutUserOutput{})
	return LogoutUserRequest{Request: req, Input: input, Copy: c.LogoutUserRequest}
}

// LogoutUserRequest is the request type for the
// LogoutUser API operation.
type LogoutUserRequest struct {
	*aws.Request
	Input *types.LogoutUserInput
	Copy  func(*types.LogoutUserInput) LogoutUserRequest
}

// Send marshals and sends the LogoutUser API request.
func (r LogoutUserRequest) Send(ctx context.Context) (*LogoutUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &LogoutUserResponse{
		LogoutUserOutput: r.Request.Data.(*types.LogoutUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// LogoutUserResponse is the response type for the
// LogoutUser API operation.
type LogoutUserResponse struct {
	*types.LogoutUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// LogoutUser request.
func (r *LogoutUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
