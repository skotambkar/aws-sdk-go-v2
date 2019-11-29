// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
)

const opUpdateNotificationSettings = "UpdateNotificationSettings"

// UpdateNotificationSettingsRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The UpdateNotificationSettings operation creates, updates, disables or re-enables
// notifications for a HIT type. If you call the UpdateNotificationSettings
// operation for a HIT type that already has a notification specification, the
// operation replaces the old specification with a new one. You can call the
// UpdateNotificationSettings operation to enable or disable notifications for
// the HIT type, without having to modify the notification specification itself
// by providing updates to the Active status without specifying a new notification
// specification. To change the Active status of a HIT type's notifications,
// the HIT type must already have a notification specification, or one must
// be provided in the same call to UpdateNotificationSettings.
//
//    // Example sending a request using UpdateNotificationSettingsRequest.
//    req := client.UpdateNotificationSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/UpdateNotificationSettings
func (c *Client) UpdateNotificationSettingsRequest(input *types.UpdateNotificationSettingsInput) UpdateNotificationSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateNotificationSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateNotificationSettingsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateNotificationSettingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateNotificationSettingsMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateNotificationSettingsRequest{Request: req, Input: input, Copy: c.UpdateNotificationSettingsRequest}
}

// UpdateNotificationSettingsRequest is the request type for the
// UpdateNotificationSettings API operation.
type UpdateNotificationSettingsRequest struct {
	*aws.Request
	Input *types.UpdateNotificationSettingsInput
	Copy  func(*types.UpdateNotificationSettingsInput) UpdateNotificationSettingsRequest
}

// Send marshals and sends the UpdateNotificationSettings API request.
func (r UpdateNotificationSettingsRequest) Send(ctx context.Context) (*UpdateNotificationSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNotificationSettingsResponse{
		UpdateNotificationSettingsOutput: r.Request.Data.(*types.UpdateNotificationSettingsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNotificationSettingsResponse is the response type for the
// UpdateNotificationSettings API operation.
type UpdateNotificationSettingsResponse struct {
	*types.UpdateNotificationSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNotificationSettings request.
func (r *UpdateNotificationSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
