// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opDeleteVoiceConnectorStreamingConfiguration = "DeleteVoiceConnectorStreamingConfiguration"

// DeleteVoiceConnectorStreamingConfigurationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the streaming configuration for the specified Amazon Chime Voice
// Connector.
//
//    // Example sending a request using DeleteVoiceConnectorStreamingConfigurationRequest.
//    req := client.DeleteVoiceConnectorStreamingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteVoiceConnectorStreamingConfiguration
func (c *Client) DeleteVoiceConnectorStreamingConfigurationRequest(input *types.DeleteVoiceConnectorStreamingConfigurationInput) DeleteVoiceConnectorStreamingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteVoiceConnectorStreamingConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/streaming-configuration",
	}

	if input == nil {
		input = &types.DeleteVoiceConnectorStreamingConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteVoiceConnectorStreamingConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteVoiceConnectorStreamingConfigurationRequest{Request: req, Input: input, Copy: c.DeleteVoiceConnectorStreamingConfigurationRequest}
}

// DeleteVoiceConnectorStreamingConfigurationRequest is the request type for the
// DeleteVoiceConnectorStreamingConfiguration API operation.
type DeleteVoiceConnectorStreamingConfigurationRequest struct {
	*aws.Request
	Input *types.DeleteVoiceConnectorStreamingConfigurationInput
	Copy  func(*types.DeleteVoiceConnectorStreamingConfigurationInput) DeleteVoiceConnectorStreamingConfigurationRequest
}

// Send marshals and sends the DeleteVoiceConnectorStreamingConfiguration API request.
func (r DeleteVoiceConnectorStreamingConfigurationRequest) Send(ctx context.Context) (*DeleteVoiceConnectorStreamingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVoiceConnectorStreamingConfigurationResponse{
		DeleteVoiceConnectorStreamingConfigurationOutput: r.Request.Data.(*types.DeleteVoiceConnectorStreamingConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVoiceConnectorStreamingConfigurationResponse is the response type for the
// DeleteVoiceConnectorStreamingConfiguration API operation.
type DeleteVoiceConnectorStreamingConfigurationResponse struct {
	*types.DeleteVoiceConnectorStreamingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVoiceConnectorStreamingConfiguration request.
func (r *DeleteVoiceConnectorStreamingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
