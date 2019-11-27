// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opUpdateNetworkProfile = "UpdateNetworkProfile"

// UpdateNetworkProfileRequest returns a request value for making API operation for
// Alexa For Business.
//
// Updates a network profile by the network profile ARN.
//
//    // Example sending a request using UpdateNetworkProfileRequest.
//    req := client.UpdateNetworkProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/UpdateNetworkProfile
func (c *Client) UpdateNetworkProfileRequest(input *types.UpdateNetworkProfileInput) UpdateNetworkProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateNetworkProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateNetworkProfileInput{}
	}

	req := c.newRequest(op, input, &types.UpdateNetworkProfileOutput{})
	return UpdateNetworkProfileRequest{Request: req, Input: input, Copy: c.UpdateNetworkProfileRequest}
}

// UpdateNetworkProfileRequest is the request type for the
// UpdateNetworkProfile API operation.
type UpdateNetworkProfileRequest struct {
	*aws.Request
	Input *types.UpdateNetworkProfileInput
	Copy  func(*types.UpdateNetworkProfileInput) UpdateNetworkProfileRequest
}

// Send marshals and sends the UpdateNetworkProfile API request.
func (r UpdateNetworkProfileRequest) Send(ctx context.Context) (*UpdateNetworkProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNetworkProfileResponse{
		UpdateNetworkProfileOutput: r.Request.Data.(*types.UpdateNetworkProfileOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNetworkProfileResponse is the response type for the
// UpdateNetworkProfile API operation.
type UpdateNetworkProfileResponse struct {
	*types.UpdateNetworkProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNetworkProfile request.
func (r *UpdateNetworkProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
