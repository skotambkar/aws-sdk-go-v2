// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetAdmChannel = "GetAdmChannel"

// GetAdmChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of the ADM channel for
// an application.
//
//    // Example sending a request using GetAdmChannelRequest.
//    req := client.GetAdmChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetAdmChannel
func (c *Client) GetAdmChannelRequest(input *types.GetAdmChannelInput) GetAdmChannelRequest {
	op := &aws.Operation{
		Name:       opGetAdmChannel,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/channels/adm",
	}

	if input == nil {
		input = &types.GetAdmChannelInput{}
	}

	req := c.newRequest(op, input, &types.GetAdmChannelOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetAdmChannelMarshaler{Input: input}.GetNamedBuildHandler())

	return GetAdmChannelRequest{Request: req, Input: input, Copy: c.GetAdmChannelRequest}
}

// GetAdmChannelRequest is the request type for the
// GetAdmChannel API operation.
type GetAdmChannelRequest struct {
	*aws.Request
	Input *types.GetAdmChannelInput
	Copy  func(*types.GetAdmChannelInput) GetAdmChannelRequest
}

// Send marshals and sends the GetAdmChannel API request.
func (r GetAdmChannelRequest) Send(ctx context.Context) (*GetAdmChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAdmChannelResponse{
		GetAdmChannelOutput: r.Request.Data.(*types.GetAdmChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAdmChannelResponse is the response type for the
// GetAdmChannel API operation.
type GetAdmChannelResponse struct {
	*types.GetAdmChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAdmChannel request.
func (r *GetAdmChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
