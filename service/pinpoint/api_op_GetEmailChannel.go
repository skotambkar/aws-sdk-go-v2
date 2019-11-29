// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetEmailChannel = "GetEmailChannel"

// GetEmailChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves information about the status and settings of the email channel
// for an application.
//
//    // Example sending a request using GetEmailChannelRequest.
//    req := client.GetEmailChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEmailChannel
func (c *Client) GetEmailChannelRequest(input *types.GetEmailChannelInput) GetEmailChannelRequest {
	op := &aws.Operation{
		Name:       opGetEmailChannel,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/channels/email",
	}

	if input == nil {
		input = &types.GetEmailChannelInput{}
	}

	req := c.newRequest(op, input, &types.GetEmailChannelOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetEmailChannelMarshaler{Input: input}.GetNamedBuildHandler())

	return GetEmailChannelRequest{Request: req, Input: input, Copy: c.GetEmailChannelRequest}
}

// GetEmailChannelRequest is the request type for the
// GetEmailChannel API operation.
type GetEmailChannelRequest struct {
	*aws.Request
	Input *types.GetEmailChannelInput
	Copy  func(*types.GetEmailChannelInput) GetEmailChannelRequest
}

// Send marshals and sends the GetEmailChannel API request.
func (r GetEmailChannelRequest) Send(ctx context.Context) (*GetEmailChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEmailChannelResponse{
		GetEmailChannelOutput: r.Request.Data.(*types.GetEmailChannelOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEmailChannelResponse is the response type for the
// GetEmailChannel API operation.
type GetEmailChannelResponse struct {
	*types.GetEmailChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEmailChannel request.
func (r *GetEmailChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
