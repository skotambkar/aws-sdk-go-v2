// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetVoiceTemplate = "GetVoiceTemplate"

// GetVoiceTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves the content and settings for a message template that you can use
// in messages that are sent through the voice channel.
//
//    // Example sending a request using GetVoiceTemplateRequest.
//    req := client.GetVoiceTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetVoiceTemplate
func (c *Client) GetVoiceTemplateRequest(input *types.GetVoiceTemplateInput) GetVoiceTemplateRequest {
	op := &aws.Operation{
		Name:       opGetVoiceTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/templates/{template-name}/voice",
	}

	if input == nil {
		input = &types.GetVoiceTemplateInput{}
	}

	req := c.newRequest(op, input, &types.GetVoiceTemplateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetVoiceTemplateMarshaler{Input: input}.GetNamedBuildHandler())

	return GetVoiceTemplateRequest{Request: req, Input: input, Copy: c.GetVoiceTemplateRequest}
}

// GetVoiceTemplateRequest is the request type for the
// GetVoiceTemplate API operation.
type GetVoiceTemplateRequest struct {
	*aws.Request
	Input *types.GetVoiceTemplateInput
	Copy  func(*types.GetVoiceTemplateInput) GetVoiceTemplateRequest
}

// Send marshals and sends the GetVoiceTemplate API request.
func (r GetVoiceTemplateRequest) Send(ctx context.Context) (*GetVoiceTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVoiceTemplateResponse{
		GetVoiceTemplateOutput: r.Request.Data.(*types.GetVoiceTemplateOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVoiceTemplateResponse is the response type for the
// GetVoiceTemplate API operation.
type GetVoiceTemplateResponse struct {
	*types.GetVoiceTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVoiceTemplate request.
func (r *GetVoiceTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
