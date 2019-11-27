// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opDeleteEventsConfiguration = "DeleteEventsConfiguration"

// DeleteEventsConfigurationRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the events configuration that allows a bot to receive outgoing events.
//
//    // Example sending a request using DeleteEventsConfigurationRequest.
//    req := client.DeleteEventsConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteEventsConfiguration
func (c *Client) DeleteEventsConfigurationRequest(input *types.DeleteEventsConfigurationInput) DeleteEventsConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteEventsConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{accountId}/bots/{botId}/events-configuration",
	}

	if input == nil {
		input = &types.DeleteEventsConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteEventsConfigurationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteEventsConfigurationRequest{Request: req, Input: input, Copy: c.DeleteEventsConfigurationRequest}
}

// DeleteEventsConfigurationRequest is the request type for the
// DeleteEventsConfiguration API operation.
type DeleteEventsConfigurationRequest struct {
	*aws.Request
	Input *types.DeleteEventsConfigurationInput
	Copy  func(*types.DeleteEventsConfigurationInput) DeleteEventsConfigurationRequest
}

// Send marshals and sends the DeleteEventsConfiguration API request.
func (r DeleteEventsConfigurationRequest) Send(ctx context.Context) (*DeleteEventsConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEventsConfigurationResponse{
		DeleteEventsConfigurationOutput: r.Request.Data.(*types.DeleteEventsConfigurationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEventsConfigurationResponse is the response type for the
// DeleteEventsConfiguration API operation.
type DeleteEventsConfigurationResponse struct {
	*types.DeleteEventsConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEventsConfiguration request.
func (r *DeleteEventsConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
