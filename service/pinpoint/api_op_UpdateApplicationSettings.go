// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opUpdateApplicationSettings = "UpdateApplicationSettings"

// UpdateApplicationSettingsRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the settings for an application.
//
//    // Example sending a request using UpdateApplicationSettingsRequest.
//    req := client.UpdateApplicationSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateApplicationSettings
func (c *Client) UpdateApplicationSettingsRequest(input *types.UpdateApplicationSettingsInput) UpdateApplicationSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateApplicationSettings,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/settings",
	}

	if input == nil {
		input = &types.UpdateApplicationSettingsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateApplicationSettingsOutput{})
	return UpdateApplicationSettingsRequest{Request: req, Input: input, Copy: c.UpdateApplicationSettingsRequest}
}

// UpdateApplicationSettingsRequest is the request type for the
// UpdateApplicationSettings API operation.
type UpdateApplicationSettingsRequest struct {
	*aws.Request
	Input *types.UpdateApplicationSettingsInput
	Copy  func(*types.UpdateApplicationSettingsInput) UpdateApplicationSettingsRequest
}

// Send marshals and sends the UpdateApplicationSettings API request.
func (r UpdateApplicationSettingsRequest) Send(ctx context.Context) (*UpdateApplicationSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApplicationSettingsResponse{
		UpdateApplicationSettingsOutput: r.Request.Data.(*types.UpdateApplicationSettingsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApplicationSettingsResponse is the response type for the
// UpdateApplicationSettings API operation.
type UpdateApplicationSettingsResponse struct {
	*types.UpdateApplicationSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApplicationSettings request.
func (r *UpdateApplicationSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
