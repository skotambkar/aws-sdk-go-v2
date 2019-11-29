// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opUpdateUserSettings = "UpdateUserSettings"

// UpdateUserSettingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates the settings for the specified user, such as phone number settings.
//
//    // Example sending a request using UpdateUserSettingsRequest.
//    req := client.UpdateUserSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateUserSettings
func (c *Client) UpdateUserSettingsRequest(input *types.UpdateUserSettingsInput) UpdateUserSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateUserSettings,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{accountId}/users/{userId}/settings",
	}

	if input == nil {
		input = &types.UpdateUserSettingsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateUserSettingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateUserSettingsMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateUserSettingsRequest{Request: req, Input: input, Copy: c.UpdateUserSettingsRequest}
}

// UpdateUserSettingsRequest is the request type for the
// UpdateUserSettings API operation.
type UpdateUserSettingsRequest struct {
	*aws.Request
	Input *types.UpdateUserSettingsInput
	Copy  func(*types.UpdateUserSettingsInput) UpdateUserSettingsRequest
}

// Send marshals and sends the UpdateUserSettings API request.
func (r UpdateUserSettingsRequest) Send(ctx context.Context) (*UpdateUserSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUserSettingsResponse{
		UpdateUserSettingsOutput: r.Request.Data.(*types.UpdateUserSettingsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUserSettingsResponse is the response type for the
// UpdateUserSettings API operation.
type UpdateUserSettingsResponse struct {
	*types.UpdateUserSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUserSettings request.
func (r *UpdateUserSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
