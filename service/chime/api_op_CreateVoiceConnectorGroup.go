// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opCreateVoiceConnectorGroup = "CreateVoiceConnectorGroup"

// CreateVoiceConnectorGroupRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates an Amazon Chime Voice Connector group under the administrator's AWS
// account. You can associate up to three existing Amazon Chime Voice Connectors
// with the Amazon Chime Voice Connector group by including VoiceConnectorItems
// in the request.
//
// You can include Amazon Chime Voice Connectors from different AWS Regions
// in your group. This creates a fault tolerant mechanism for fallback in case
// of availability events.
//
//    // Example sending a request using CreateVoiceConnectorGroupRequest.
//    req := client.CreateVoiceConnectorGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreateVoiceConnectorGroup
func (c *Client) CreateVoiceConnectorGroupRequest(input *types.CreateVoiceConnectorGroupInput) CreateVoiceConnectorGroupRequest {
	op := &aws.Operation{
		Name:       opCreateVoiceConnectorGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/voice-connector-groups",
	}

	if input == nil {
		input = &types.CreateVoiceConnectorGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateVoiceConnectorGroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateVoiceConnectorGroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateVoiceConnectorGroupRequest{Request: req, Input: input, Copy: c.CreateVoiceConnectorGroupRequest}
}

// CreateVoiceConnectorGroupRequest is the request type for the
// CreateVoiceConnectorGroup API operation.
type CreateVoiceConnectorGroupRequest struct {
	*aws.Request
	Input *types.CreateVoiceConnectorGroupInput
	Copy  func(*types.CreateVoiceConnectorGroupInput) CreateVoiceConnectorGroupRequest
}

// Send marshals and sends the CreateVoiceConnectorGroup API request.
func (r CreateVoiceConnectorGroupRequest) Send(ctx context.Context) (*CreateVoiceConnectorGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVoiceConnectorGroupResponse{
		CreateVoiceConnectorGroupOutput: r.Request.Data.(*types.CreateVoiceConnectorGroupOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVoiceConnectorGroupResponse is the response type for the
// CreateVoiceConnectorGroup API operation.
type CreateVoiceConnectorGroupResponse struct {
	*types.CreateVoiceConnectorGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVoiceConnectorGroup request.
func (r *CreateVoiceConnectorGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
