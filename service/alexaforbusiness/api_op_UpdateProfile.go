// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opUpdateProfile = "UpdateProfile"

// UpdateProfileRequest returns a request value for making API operation for
// Alexa For Business.
//
// Updates an existing room profile by room profile ARN.
//
//    // Example sending a request using UpdateProfileRequest.
//    req := client.UpdateProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateProfile
func (c *Client) UpdateProfileRequest(input *types.UpdateProfileInput) UpdateProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateProfileInput{}
	}

	req := c.newRequest(op, input, &types.UpdateProfileOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateProfileMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateProfileRequest{Request: req, Input: input, Copy: c.UpdateProfileRequest}
}

// UpdateProfileRequest is the request type for the
// UpdateProfile API operation.
type UpdateProfileRequest struct {
	*aws.Request
	Input *types.UpdateProfileInput
	Copy  func(*types.UpdateProfileInput) UpdateProfileRequest
}

// Send marshals and sends the UpdateProfile API request.
func (r UpdateProfileRequest) Send(ctx context.Context) (*UpdateProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateProfileResponse{
		UpdateProfileOutput: r.Request.Data.(*types.UpdateProfileOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateProfileResponse is the response type for the
// UpdateProfile API operation.
type UpdateProfileResponse struct {
	*types.UpdateProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateProfile request.
func (r *UpdateProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
