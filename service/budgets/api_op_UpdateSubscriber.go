// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
)

const opUpdateSubscriber = "UpdateSubscriber"

// UpdateSubscriberRequest returns a request value for making API operation for
// AWS Budgets.
//
// Updates a subscriber.
//
//    // Example sending a request using UpdateSubscriberRequest.
//    req := client.UpdateSubscriberRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateSubscriberRequest(input *types.UpdateSubscriberInput) UpdateSubscriberRequest {
	op := &aws.Operation{
		Name:       opUpdateSubscriber,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateSubscriberInput{}
	}

	req := c.newRequest(op, input, &types.UpdateSubscriberOutput{})
	return UpdateSubscriberRequest{Request: req, Input: input, Copy: c.UpdateSubscriberRequest}
}

// UpdateSubscriberRequest is the request type for the
// UpdateSubscriber API operation.
type UpdateSubscriberRequest struct {
	*aws.Request
	Input *types.UpdateSubscriberInput
	Copy  func(*types.UpdateSubscriberInput) UpdateSubscriberRequest
}

// Send marshals and sends the UpdateSubscriber API request.
func (r UpdateSubscriberRequest) Send(ctx context.Context) (*UpdateSubscriberResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSubscriberResponse{
		UpdateSubscriberOutput: r.Request.Data.(*types.UpdateSubscriberOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSubscriberResponse is the response type for the
// UpdateSubscriber API operation.
type UpdateSubscriberResponse struct {
	*types.UpdateSubscriberOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSubscriber request.
func (r *UpdateSubscriberResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
