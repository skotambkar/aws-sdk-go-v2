// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opPutVoiceConnectorOrigination = "PutVoiceConnectorOrigination"

// PutVoiceConnectorOriginationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Adds origination settings for the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using PutVoiceConnectorOriginationRequest.
//    req := client.PutVoiceConnectorOriginationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorOrigination
func (c *Client) PutVoiceConnectorOriginationRequest(input *types.PutVoiceConnectorOriginationInput) PutVoiceConnectorOriginationRequest {
	op := &aws.Operation{
		Name:       opPutVoiceConnectorOrigination,
		HTTPMethod: "PUT",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/origination",
	}

	if input == nil {
		input = &types.PutVoiceConnectorOriginationInput{}
	}

	req := c.newRequest(op, input, &types.PutVoiceConnectorOriginationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.PutVoiceConnectorOriginationMarshaler{Input: input}.GetNamedBuildHandler())

	return PutVoiceConnectorOriginationRequest{Request: req, Input: input, Copy: c.PutVoiceConnectorOriginationRequest}
}

// PutVoiceConnectorOriginationRequest is the request type for the
// PutVoiceConnectorOrigination API operation.
type PutVoiceConnectorOriginationRequest struct {
	*aws.Request
	Input *types.PutVoiceConnectorOriginationInput
	Copy  func(*types.PutVoiceConnectorOriginationInput) PutVoiceConnectorOriginationRequest
}

// Send marshals and sends the PutVoiceConnectorOrigination API request.
func (r PutVoiceConnectorOriginationRequest) Send(ctx context.Context) (*PutVoiceConnectorOriginationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutVoiceConnectorOriginationResponse{
		PutVoiceConnectorOriginationOutput: r.Request.Data.(*types.PutVoiceConnectorOriginationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutVoiceConnectorOriginationResponse is the response type for the
// PutVoiceConnectorOrigination API operation.
type PutVoiceConnectorOriginationResponse struct {
	*types.PutVoiceConnectorOriginationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutVoiceConnectorOrigination request.
func (r *PutVoiceConnectorOriginationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
