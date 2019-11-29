// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opRemoveSourceIdentifierFromSubscription = "RemoveSourceIdentifierFromSubscription"

// RemoveSourceIdentifierFromSubscriptionRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Removes a source identifier from an existing event notification subscription.
//
//    // Example sending a request using RemoveSourceIdentifierFromSubscriptionRequest.
//    req := client.RemoveSourceIdentifierFromSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RemoveSourceIdentifierFromSubscription
func (c *Client) RemoveSourceIdentifierFromSubscriptionRequest(input *types.RemoveSourceIdentifierFromSubscriptionInput) RemoveSourceIdentifierFromSubscriptionRequest {
	op := &aws.Operation{
		Name:       opRemoveSourceIdentifierFromSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RemoveSourceIdentifierFromSubscriptionInput{}
	}

	req := c.newRequest(op, input, &types.RemoveSourceIdentifierFromSubscriptionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.RemoveSourceIdentifierFromSubscriptionMarshaler{Input: input}.GetNamedBuildHandler())

	return RemoveSourceIdentifierFromSubscriptionRequest{Request: req, Input: input, Copy: c.RemoveSourceIdentifierFromSubscriptionRequest}
}

// RemoveSourceIdentifierFromSubscriptionRequest is the request type for the
// RemoveSourceIdentifierFromSubscription API operation.
type RemoveSourceIdentifierFromSubscriptionRequest struct {
	*aws.Request
	Input *types.RemoveSourceIdentifierFromSubscriptionInput
	Copy  func(*types.RemoveSourceIdentifierFromSubscriptionInput) RemoveSourceIdentifierFromSubscriptionRequest
}

// Send marshals and sends the RemoveSourceIdentifierFromSubscription API request.
func (r RemoveSourceIdentifierFromSubscriptionRequest) Send(ctx context.Context) (*RemoveSourceIdentifierFromSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveSourceIdentifierFromSubscriptionResponse{
		RemoveSourceIdentifierFromSubscriptionOutput: r.Request.Data.(*types.RemoveSourceIdentifierFromSubscriptionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveSourceIdentifierFromSubscriptionResponse is the response type for the
// RemoveSourceIdentifierFromSubscription API operation.
type RemoveSourceIdentifierFromSubscriptionResponse struct {
	*types.RemoveSourceIdentifierFromSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveSourceIdentifierFromSubscription request.
func (r *RemoveSourceIdentifierFromSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
