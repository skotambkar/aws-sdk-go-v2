// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
)

const opUpdateWebhook = "UpdateWebhook"

// UpdateWebhookRequest returns a request value for making API operation for
// AWS Amplify.
//
// Update a webhook.
//
//    // Example sending a request using UpdateWebhookRequest.
//    req := client.UpdateWebhookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateWebhook
func (c *Client) UpdateWebhookRequest(input *types.UpdateWebhookInput) UpdateWebhookRequest {
	op := &aws.Operation{
		Name:       opUpdateWebhook,
		HTTPMethod: "POST",
		HTTPPath:   "/webhooks/{webhookId}",
	}

	if input == nil {
		input = &types.UpdateWebhookInput{}
	}

	req := c.newRequest(op, input, &types.UpdateWebhookOutput{})
	return UpdateWebhookRequest{Request: req, Input: input, Copy: c.UpdateWebhookRequest}
}

// UpdateWebhookRequest is the request type for the
// UpdateWebhook API operation.
type UpdateWebhookRequest struct {
	*aws.Request
	Input *types.UpdateWebhookInput
	Copy  func(*types.UpdateWebhookInput) UpdateWebhookRequest
}

// Send marshals and sends the UpdateWebhook API request.
func (r UpdateWebhookRequest) Send(ctx context.Context) (*UpdateWebhookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateWebhookResponse{
		UpdateWebhookOutput: r.Request.Data.(*types.UpdateWebhookOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateWebhookResponse is the response type for the
// UpdateWebhook API operation.
type UpdateWebhookResponse struct {
	*types.UpdateWebhookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateWebhook request.
func (r *UpdateWebhookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
