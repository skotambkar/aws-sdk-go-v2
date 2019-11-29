// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codestar/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/codestar/types"
)

const opCreateUserProfile = "CreateUserProfile"

// CreateUserProfileRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Creates a profile for a user that includes user preferences, such as the
// display name and email address assocciated with the user, in AWS CodeStar.
// The user profile is not project-specific. Information in the user profile
// is displayed wherever the user's information appears to other users in AWS
// CodeStar.
//
//    // Example sending a request using CreateUserProfileRequest.
//    req := client.CreateUserProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/CreateUserProfile
func (c *Client) CreateUserProfileRequest(input *types.CreateUserProfileInput) CreateUserProfileRequest {
	op := &aws.Operation{
		Name:       opCreateUserProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateUserProfileInput{}
	}

	req := c.newRequest(op, input, &types.CreateUserProfileOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateUserProfileMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateUserProfileRequest{Request: req, Input: input, Copy: c.CreateUserProfileRequest}
}

// CreateUserProfileRequest is the request type for the
// CreateUserProfile API operation.
type CreateUserProfileRequest struct {
	*aws.Request
	Input *types.CreateUserProfileInput
	Copy  func(*types.CreateUserProfileInput) CreateUserProfileRequest
}

// Send marshals and sends the CreateUserProfile API request.
func (r CreateUserProfileRequest) Send(ctx context.Context) (*CreateUserProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserProfileResponse{
		CreateUserProfileOutput: r.Request.Data.(*types.CreateUserProfileOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserProfileResponse is the response type for the
// CreateUserProfile API operation.
type CreateUserProfileResponse struct {
	*types.CreateUserProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUserProfile request.
func (r *CreateUserProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
