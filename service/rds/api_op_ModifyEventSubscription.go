// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/rds/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opModifyEventSubscription = "ModifyEventSubscription"

// ModifyEventSubscriptionRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Modifies an existing RDS event notification subscription. Note that you can't
// modify the source identifiers using this call; to change source identifiers
// for a subscription, use the AddSourceIdentifierToSubscription and RemoveSourceIdentifierFromSubscription
// calls.
//
// You can see a list of the event categories for a given SourceType in the
// Events (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html)
// topic in the Amazon RDS User Guide or by using the DescribeEventCategories
// action.
//
//    // Example sending a request using ModifyEventSubscriptionRequest.
//    req := client.ModifyEventSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/ModifyEventSubscription
func (c *Client) ModifyEventSubscriptionRequest(input *types.ModifyEventSubscriptionInput) ModifyEventSubscriptionRequest {
	op := &aws.Operation{
		Name:       opModifyEventSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyEventSubscriptionInput{}
	}

	req := c.newRequest(op, input, &types.ModifyEventSubscriptionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.ModifyEventSubscriptionMarshaler{Input: input}.GetNamedBuildHandler())

	return ModifyEventSubscriptionRequest{Request: req, Input: input, Copy: c.ModifyEventSubscriptionRequest}
}

// ModifyEventSubscriptionRequest is the request type for the
// ModifyEventSubscription API operation.
type ModifyEventSubscriptionRequest struct {
	*aws.Request
	Input *types.ModifyEventSubscriptionInput
	Copy  func(*types.ModifyEventSubscriptionInput) ModifyEventSubscriptionRequest
}

// Send marshals and sends the ModifyEventSubscription API request.
func (r ModifyEventSubscriptionRequest) Send(ctx context.Context) (*ModifyEventSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyEventSubscriptionResponse{
		ModifyEventSubscriptionOutput: r.Request.Data.(*types.ModifyEventSubscriptionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyEventSubscriptionResponse is the response type for the
// ModifyEventSubscription API operation.
type ModifyEventSubscriptionResponse struct {
	*types.ModifyEventSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyEventSubscription request.
func (r *ModifyEventSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
