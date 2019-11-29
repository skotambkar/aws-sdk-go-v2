// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opGetPhoneNumberSettings = "GetPhoneNumberSettings"

// GetPhoneNumberSettingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Retrieves the phone number settings for the administrator's AWS account,
// such as the default outbound calling name.
//
//    // Example sending a request using GetPhoneNumberSettingsRequest.
//    req := client.GetPhoneNumberSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/GetPhoneNumberSettings
func (c *Client) GetPhoneNumberSettingsRequest(input *types.GetPhoneNumberSettingsInput) GetPhoneNumberSettingsRequest {
	op := &aws.Operation{
		Name:       opGetPhoneNumberSettings,
		HTTPMethod: "GET",
		HTTPPath:   "/settings/phone-number",
	}

	if input == nil {
		input = &types.GetPhoneNumberSettingsInput{}
	}

	req := c.newRequest(op, input, &types.GetPhoneNumberSettingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetPhoneNumberSettingsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetPhoneNumberSettingsRequest{Request: req, Input: input, Copy: c.GetPhoneNumberSettingsRequest}
}

// GetPhoneNumberSettingsRequest is the request type for the
// GetPhoneNumberSettings API operation.
type GetPhoneNumberSettingsRequest struct {
	*aws.Request
	Input *types.GetPhoneNumberSettingsInput
	Copy  func(*types.GetPhoneNumberSettingsInput) GetPhoneNumberSettingsRequest
}

// Send marshals and sends the GetPhoneNumberSettings API request.
func (r GetPhoneNumberSettingsRequest) Send(ctx context.Context) (*GetPhoneNumberSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetPhoneNumberSettingsResponse{
		GetPhoneNumberSettingsOutput: r.Request.Data.(*types.GetPhoneNumberSettingsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetPhoneNumberSettingsResponse is the response type for the
// GetPhoneNumberSettings API operation.
type GetPhoneNumberSettingsResponse struct {
	*types.GetPhoneNumberSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetPhoneNumberSettings request.
func (r *GetPhoneNumberSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
