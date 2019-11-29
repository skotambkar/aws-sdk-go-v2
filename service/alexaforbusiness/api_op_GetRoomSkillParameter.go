// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opGetRoomSkillParameter = "GetRoomSkillParameter"

// GetRoomSkillParameterRequest returns a request value for making API operation for
// Alexa For Business.
//
// Gets room skill parameter details by room, skill, and parameter key ARN.
//
//    // Example sending a request using GetRoomSkillParameterRequest.
//    req := client.GetRoomSkillParameterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/GetRoomSkillParameter
func (c *Client) GetRoomSkillParameterRequest(input *types.GetRoomSkillParameterInput) GetRoomSkillParameterRequest {
	op := &aws.Operation{
		Name:       opGetRoomSkillParameter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetRoomSkillParameterInput{}
	}

	req := c.newRequest(op, input, &types.GetRoomSkillParameterOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetRoomSkillParameterMarshaler{Input: input}.GetNamedBuildHandler())

	return GetRoomSkillParameterRequest{Request: req, Input: input, Copy: c.GetRoomSkillParameterRequest}
}

// GetRoomSkillParameterRequest is the request type for the
// GetRoomSkillParameter API operation.
type GetRoomSkillParameterRequest struct {
	*aws.Request
	Input *types.GetRoomSkillParameterInput
	Copy  func(*types.GetRoomSkillParameterInput) GetRoomSkillParameterRequest
}

// Send marshals and sends the GetRoomSkillParameter API request.
func (r GetRoomSkillParameterRequest) Send(ctx context.Context) (*GetRoomSkillParameterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRoomSkillParameterResponse{
		GetRoomSkillParameterOutput: r.Request.Data.(*types.GetRoomSkillParameterOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRoomSkillParameterResponse is the response type for the
// GetRoomSkillParameter API operation.
type GetRoomSkillParameterResponse struct {
	*types.GetRoomSkillParameterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRoomSkillParameter request.
func (r *GetRoomSkillParameterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
