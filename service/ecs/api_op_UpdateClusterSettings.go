// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const opUpdateClusterSettings = "UpdateClusterSettings"

// UpdateClusterSettingsRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Modifies the settings to use for a cluster.
//
//    // Example sending a request using UpdateClusterSettingsRequest.
//    req := client.UpdateClusterSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/UpdateClusterSettings
func (c *Client) UpdateClusterSettingsRequest(input *types.UpdateClusterSettingsInput) UpdateClusterSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateClusterSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateClusterSettingsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateClusterSettingsOutput{})
	return UpdateClusterSettingsRequest{Request: req, Input: input, Copy: c.UpdateClusterSettingsRequest}
}

// UpdateClusterSettingsRequest is the request type for the
// UpdateClusterSettings API operation.
type UpdateClusterSettingsRequest struct {
	*aws.Request
	Input *types.UpdateClusterSettingsInput
	Copy  func(*types.UpdateClusterSettingsInput) UpdateClusterSettingsRequest
}

// Send marshals and sends the UpdateClusterSettings API request.
func (r UpdateClusterSettingsRequest) Send(ctx context.Context) (*UpdateClusterSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClusterSettingsResponse{
		UpdateClusterSettingsOutput: r.Request.Data.(*types.UpdateClusterSettingsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClusterSettingsResponse is the response type for the
// UpdateClusterSettings API operation.
type UpdateClusterSettingsResponse struct {
	*types.UpdateClusterSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClusterSettings request.
func (r *UpdateClusterSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
