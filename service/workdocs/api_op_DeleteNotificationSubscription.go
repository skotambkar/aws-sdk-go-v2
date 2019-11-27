// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
)

const opDeleteNotificationSubscription = "DeleteNotificationSubscription"

// DeleteNotificationSubscriptionRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Deletes the specified subscription from the specified organization.
//
//    // Example sending a request using DeleteNotificationSubscriptionRequest.
//    req := client.DeleteNotificationSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteNotificationSubscription
func (c *Client) DeleteNotificationSubscriptionRequest(input *types.DeleteNotificationSubscriptionInput) DeleteNotificationSubscriptionRequest {
	op := &aws.Operation{
		Name:       opDeleteNotificationSubscription,
		HTTPMethod: "DELETE",
		HTTPPath:   "/api/v1/organizations/{OrganizationId}/subscriptions/{SubscriptionId}",
	}

	if input == nil {
		input = &types.DeleteNotificationSubscriptionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteNotificationSubscriptionOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteNotificationSubscriptionRequest{Request: req, Input: input, Copy: c.DeleteNotificationSubscriptionRequest}
}

// DeleteNotificationSubscriptionRequest is the request type for the
// DeleteNotificationSubscription API operation.
type DeleteNotificationSubscriptionRequest struct {
	*aws.Request
	Input *types.DeleteNotificationSubscriptionInput
	Copy  func(*types.DeleteNotificationSubscriptionInput) DeleteNotificationSubscriptionRequest
}

// Send marshals and sends the DeleteNotificationSubscription API request.
func (r DeleteNotificationSubscriptionRequest) Send(ctx context.Context) (*DeleteNotificationSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNotificationSubscriptionResponse{
		DeleteNotificationSubscriptionOutput: r.Request.Data.(*types.DeleteNotificationSubscriptionOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNotificationSubscriptionResponse is the response type for the
// DeleteNotificationSubscription API operation.
type DeleteNotificationSubscriptionResponse struct {
	*types.DeleteNotificationSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNotificationSubscription request.
func (r *DeleteNotificationSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
