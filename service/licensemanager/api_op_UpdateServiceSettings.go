// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
)

const opUpdateServiceSettings = "UpdateServiceSettings"

// UpdateServiceSettingsRequest returns a request value for making API operation for
// AWS License Manager.
//
// Updates License Manager service settings.
//
//    // Example sending a request using UpdateServiceSettingsRequest.
//    req := client.UpdateServiceSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/UpdateServiceSettings
func (c *Client) UpdateServiceSettingsRequest(input *types.UpdateServiceSettingsInput) UpdateServiceSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateServiceSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateServiceSettingsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateServiceSettingsOutput{})
	return UpdateServiceSettingsRequest{Request: req, Input: input, Copy: c.UpdateServiceSettingsRequest}
}

// UpdateServiceSettingsRequest is the request type for the
// UpdateServiceSettings API operation.
type UpdateServiceSettingsRequest struct {
	*aws.Request
	Input *types.UpdateServiceSettingsInput
	Copy  func(*types.UpdateServiceSettingsInput) UpdateServiceSettingsRequest
}

// Send marshals and sends the UpdateServiceSettings API request.
func (r UpdateServiceSettingsRequest) Send(ctx context.Context) (*UpdateServiceSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateServiceSettingsResponse{
		UpdateServiceSettingsOutput: r.Request.Data.(*types.UpdateServiceSettingsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateServiceSettingsResponse is the response type for the
// UpdateServiceSettings API operation.
type UpdateServiceSettingsResponse struct {
	*types.UpdateServiceSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateServiceSettings request.
func (r *UpdateServiceSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
