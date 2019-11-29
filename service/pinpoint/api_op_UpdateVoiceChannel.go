// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opUpdateVoiceChannel = "UpdateVoiceChannel"

// UpdateVoiceChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Enables the voice channel for an application or updates the status and settings
// of the voice channel for an application.
//
//    // Example sending a request using UpdateVoiceChannelRequest.
//    req := client.UpdateVoiceChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateVoiceChannel
func (c *Client) UpdateVoiceChannelRequest(input *types.UpdateVoiceChannelInput) UpdateVoiceChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateVoiceChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/channels/voice",
	}

	if input == nil {
		input = &types.UpdateVoiceChannelInput{}
	}

	req := c.newRequest(op, input, &types.UpdateVoiceChannelOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateVoiceChannelMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateVoiceChannelRequest{Request: req, Input: input, Copy: c.UpdateVoiceChannelRequest}
}

// UpdateVoiceChannelRequest is the request type for the
// UpdateVoiceChannel API operation.
type UpdateVoiceChannelRequest struct {
	*aws.Request
	Input *types.UpdateVoiceChannelInput
	Copy  func(*types.UpdateVoiceChannelInput) UpdateVoiceChannelRequest
}

// Send marshals and sends the UpdateVoiceChannel API request.
func (r UpdateVoiceChannelRequest) Send(ctx context.Context) (*UpdateVoiceChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVoiceChannelResponse{
		UpdateVoiceChannelOutput: r.Request.Data.(*types.UpdateVoiceChannelOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVoiceChannelResponse is the response type for the
// UpdateVoiceChannel API operation.
type UpdateVoiceChannelResponse struct {
	*types.UpdateVoiceChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVoiceChannel request.
func (r *UpdateVoiceChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
