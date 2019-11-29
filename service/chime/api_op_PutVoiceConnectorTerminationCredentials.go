// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opPutVoiceConnectorTerminationCredentials = "PutVoiceConnectorTerminationCredentials"

// PutVoiceConnectorTerminationCredentialsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Adds termination SIP credentials for the specified Amazon Chime Voice Connector.
//
//    // Example sending a request using PutVoiceConnectorTerminationCredentialsRequest.
//    req := client.PutVoiceConnectorTerminationCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/PutVoiceConnectorTerminationCredentials
func (c *Client) PutVoiceConnectorTerminationCredentialsRequest(input *types.PutVoiceConnectorTerminationCredentialsInput) PutVoiceConnectorTerminationCredentialsRequest {
	op := &aws.Operation{
		Name:       opPutVoiceConnectorTerminationCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connectors/{voiceConnectorId}/termination/credentials?operation=put",
	}

	if input == nil {
		input = &types.PutVoiceConnectorTerminationCredentialsInput{}
	}

	req := c.newRequest(op, input, &types.PutVoiceConnectorTerminationCredentialsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.PutVoiceConnectorTerminationCredentialsMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutVoiceConnectorTerminationCredentialsRequest{Request: req, Input: input, Copy: c.PutVoiceConnectorTerminationCredentialsRequest}
}

// PutVoiceConnectorTerminationCredentialsRequest is the request type for the
// PutVoiceConnectorTerminationCredentials API operation.
type PutVoiceConnectorTerminationCredentialsRequest struct {
	*aws.Request
	Input *types.PutVoiceConnectorTerminationCredentialsInput
	Copy  func(*types.PutVoiceConnectorTerminationCredentialsInput) PutVoiceConnectorTerminationCredentialsRequest
}

// Send marshals and sends the PutVoiceConnectorTerminationCredentials API request.
func (r PutVoiceConnectorTerminationCredentialsRequest) Send(ctx context.Context) (*PutVoiceConnectorTerminationCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutVoiceConnectorTerminationCredentialsResponse{
		PutVoiceConnectorTerminationCredentialsOutput: r.Request.Data.(*types.PutVoiceConnectorTerminationCredentialsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutVoiceConnectorTerminationCredentialsResponse is the response type for the
// PutVoiceConnectorTerminationCredentials API operation.
type PutVoiceConnectorTerminationCredentialsResponse struct {
	*types.PutVoiceConnectorTerminationCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutVoiceConnectorTerminationCredentials request.
func (r *PutVoiceConnectorTerminationCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
