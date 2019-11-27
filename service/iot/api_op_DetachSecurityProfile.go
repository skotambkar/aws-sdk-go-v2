// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opDetachSecurityProfile = "DetachSecurityProfile"

// DetachSecurityProfileRequest returns a request value for making API operation for
// AWS IoT.
//
// Disassociates a Device Defender security profile from a thing group or from
// this account.
//
//    // Example sending a request using DetachSecurityProfileRequest.
//    req := client.DetachSecurityProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DetachSecurityProfileRequest(input *types.DetachSecurityProfileInput) DetachSecurityProfileRequest {
	op := &aws.Operation{
		Name:       opDetachSecurityProfile,
		HTTPMethod: "DELETE",
		HTTPPath:   "/security-profiles/{securityProfileName}/targets",
	}

	if input == nil {
		input = &types.DetachSecurityProfileInput{}
	}

	req := c.newRequest(op, input, &types.DetachSecurityProfileOutput{})
	return DetachSecurityProfileRequest{Request: req, Input: input, Copy: c.DetachSecurityProfileRequest}
}

// DetachSecurityProfileRequest is the request type for the
// DetachSecurityProfile API operation.
type DetachSecurityProfileRequest struct {
	*aws.Request
	Input *types.DetachSecurityProfileInput
	Copy  func(*types.DetachSecurityProfileInput) DetachSecurityProfileRequest
}

// Send marshals and sends the DetachSecurityProfile API request.
func (r DetachSecurityProfileRequest) Send(ctx context.Context) (*DetachSecurityProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachSecurityProfileResponse{
		DetachSecurityProfileOutput: r.Request.Data.(*types.DetachSecurityProfileOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachSecurityProfileResponse is the response type for the
// DetachSecurityProfile API operation.
type DetachSecurityProfileResponse struct {
	*types.DetachSecurityProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachSecurityProfile request.
func (r *DetachSecurityProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
