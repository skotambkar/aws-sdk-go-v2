// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opDeleteDeliveryChannel = "DeleteDeliveryChannel"

// DeleteDeliveryChannelRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the delivery channel.
//
// Before you can delete the delivery channel, you must stop the configuration
// recorder by using the StopConfigurationRecorder action.
//
//    // Example sending a request using DeleteDeliveryChannelRequest.
//    req := client.DeleteDeliveryChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteDeliveryChannel
func (c *Client) DeleteDeliveryChannelRequest(input *types.DeleteDeliveryChannelInput) DeleteDeliveryChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteDeliveryChannel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteDeliveryChannelInput{}
	}

	req := c.newRequest(op, input, &types.DeleteDeliveryChannelOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDeliveryChannelRequest{Request: req, Input: input, Copy: c.DeleteDeliveryChannelRequest}
}

// DeleteDeliveryChannelRequest is the request type for the
// DeleteDeliveryChannel API operation.
type DeleteDeliveryChannelRequest struct {
	*aws.Request
	Input *types.DeleteDeliveryChannelInput
	Copy  func(*types.DeleteDeliveryChannelInput) DeleteDeliveryChannelRequest
}

// Send marshals and sends the DeleteDeliveryChannel API request.
func (r DeleteDeliveryChannelRequest) Send(ctx context.Context) (*DeleteDeliveryChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDeliveryChannelResponse{
		DeleteDeliveryChannelOutput: r.Request.Data.(*types.DeleteDeliveryChannelOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDeliveryChannelResponse is the response type for the
// DeleteDeliveryChannel API operation.
type DeleteDeliveryChannelResponse struct {
	*types.DeleteDeliveryChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDeliveryChannel request.
func (r *DeleteDeliveryChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
