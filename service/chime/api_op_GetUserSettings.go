// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opGetUserSettings = "GetUserSettings"

// GetUserSettingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves settings for the specified user ID, such as any associated phone
// number settings.
//
//    // Example sending a request using GetUserSettingsRequest.
//    req := client.GetUserSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetUserSettings
func (c *Client) GetUserSettingsRequest(input *types.GetUserSettingsInput) GetUserSettingsRequest {
	op := &aws.Operation{
		Name:       opGetUserSettings,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{accountId}/users/{userId}/settings",
	}

	if input == nil {
		input = &types.GetUserSettingsInput{}
	}

	req := c.newRequest(op, input, &types.GetUserSettingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetUserSettingsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetUserSettingsRequest{Request: req, Input: input, Copy: c.GetUserSettingsRequest}
}

// GetUserSettingsRequest is the request type for the
// GetUserSettings API operation.
type GetUserSettingsRequest struct {
	*aws.Request
	Input *types.GetUserSettingsInput
	Copy  func(*types.GetUserSettingsInput) GetUserSettingsRequest
}

// Send marshals and sends the GetUserSettings API request.
func (r GetUserSettingsRequest) Send(ctx context.Context) (*GetUserSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUserSettingsResponse{
		GetUserSettingsOutput: r.Request.Data.(*types.GetUserSettingsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUserSettingsResponse is the response type for the
// GetUserSettings API operation.
type GetUserSettingsResponse struct {
	*types.GetUserSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUserSettings request.
func (r *GetUserSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
