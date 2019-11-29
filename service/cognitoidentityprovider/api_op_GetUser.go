// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

const opGetUser = "GetUser"

// GetUserRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Gets the user attributes and metadata for a user.
//
//    // Example sending a request using GetUserRequest.
//    req := client.GetUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetUser
func (c *Client) GetUserRequest(input *types.GetUserInput) GetUserRequest {
	op := &aws.Operation{
		Name:       opGetUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetUserInput{}
	}

	req := c.newRequest(op, input, &types.GetUserOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetUserMarshaler{Input: input}.GetNamedBuildHandler())

	req.Config.Credentials = aws.AnonymousCredentials
	return GetUserRequest{Request: req, Input: input, Copy: c.GetUserRequest}
}

// GetUserRequest is the request type for the
// GetUser API operation.
type GetUserRequest struct {
	*aws.Request
	Input *types.GetUserInput
	Copy  func(*types.GetUserInput) GetUserRequest
}

// Send marshals and sends the GetUser API request.
func (r GetUserRequest) Send(ctx context.Context) (*GetUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUserResponse{
		GetUserOutput: r.Request.Data.(*types.GetUserOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUserResponse is the response type for the
// GetUser API operation.
type GetUserResponse struct {
	*types.GetUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUser request.
func (r *GetUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
