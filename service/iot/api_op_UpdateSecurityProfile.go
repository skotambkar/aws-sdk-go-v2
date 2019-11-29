// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opUpdateSecurityProfile = "UpdateSecurityProfile"

// UpdateSecurityProfileRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates a Device Defender security profile.
//
//    // Example sending a request using UpdateSecurityProfileRequest.
//    req := client.UpdateSecurityProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateSecurityProfileRequest(input *types.UpdateSecurityProfileInput) UpdateSecurityProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateSecurityProfile,
		HTTPMethod: "PATCH",
		HTTPPath:   "/security-profiles/{securityProfileName}",
	}

	if input == nil {
		input = &types.UpdateSecurityProfileInput{}
	}

	req := c.newRequest(op, input, &types.UpdateSecurityProfileOutput{})
	return UpdateSecurityProfileRequest{Request: req, Input: input, Copy: c.UpdateSecurityProfileRequest}
}

// UpdateSecurityProfileRequest is the request type for the
// UpdateSecurityProfile API operation.
type UpdateSecurityProfileRequest struct {
	*aws.Request
	Input *types.UpdateSecurityProfileInput
	Copy  func(*types.UpdateSecurityProfileInput) UpdateSecurityProfileRequest
}

// Send marshals and sends the UpdateSecurityProfile API request.
func (r UpdateSecurityProfileRequest) Send(ctx context.Context) (*UpdateSecurityProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSecurityProfileResponse{
		UpdateSecurityProfileOutput: r.Request.Data.(*types.UpdateSecurityProfileOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSecurityProfileResponse is the response type for the
// UpdateSecurityProfile API operation.
type UpdateSecurityProfileResponse struct {
	*types.UpdateSecurityProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSecurityProfile request.
func (r *UpdateSecurityProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
