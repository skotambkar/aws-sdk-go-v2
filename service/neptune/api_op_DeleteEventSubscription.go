// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opDeleteEventSubscription = "DeleteEventSubscription"

// DeleteEventSubscriptionRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Deletes an event notification subscription.
//
//    // Example sending a request using DeleteEventSubscriptionRequest.
//    req := client.DeleteEventSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DeleteEventSubscription
func (c *Client) DeleteEventSubscriptionRequest(input *types.DeleteEventSubscriptionInput) DeleteEventSubscriptionRequest {
	op := &aws.Operation{
		Name:       opDeleteEventSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteEventSubscriptionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteEventSubscriptionOutput{})
	return DeleteEventSubscriptionRequest{Request: req, Input: input, Copy: c.DeleteEventSubscriptionRequest}
}

// DeleteEventSubscriptionRequest is the request type for the
// DeleteEventSubscription API operation.
type DeleteEventSubscriptionRequest struct {
	*aws.Request
	Input *types.DeleteEventSubscriptionInput
	Copy  func(*types.DeleteEventSubscriptionInput) DeleteEventSubscriptionRequest
}

// Send marshals and sends the DeleteEventSubscription API request.
func (r DeleteEventSubscriptionRequest) Send(ctx context.Context) (*DeleteEventSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEventSubscriptionResponse{
		DeleteEventSubscriptionOutput: r.Request.Data.(*types.DeleteEventSubscriptionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEventSubscriptionResponse is the response type for the
// DeleteEventSubscription API operation.
type DeleteEventSubscriptionResponse struct {
	*types.DeleteEventSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEventSubscription request.
func (r *DeleteEventSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
