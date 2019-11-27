// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
)

const opDescribeNotificationSubscriptions = "DescribeNotificationSubscriptions"

// DescribeNotificationSubscriptionsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Lists the specified notification subscriptions.
//
//    // Example sending a request using DescribeNotificationSubscriptionsRequest.
//    req := client.DescribeNotificationSubscriptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeNotificationSubscriptions
func (c *Client) DescribeNotificationSubscriptionsRequest(input *types.DescribeNotificationSubscriptionsInput) DescribeNotificationSubscriptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeNotificationSubscriptions,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/organizations/{OrganizationId}/subscriptions",
	}

	if input == nil {
		input = &types.DescribeNotificationSubscriptionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeNotificationSubscriptionsOutput{})
	return DescribeNotificationSubscriptionsRequest{Request: req, Input: input, Copy: c.DescribeNotificationSubscriptionsRequest}
}

// DescribeNotificationSubscriptionsRequest is the request type for the
// DescribeNotificationSubscriptions API operation.
type DescribeNotificationSubscriptionsRequest struct {
	*aws.Request
	Input *types.DescribeNotificationSubscriptionsInput
	Copy  func(*types.DescribeNotificationSubscriptionsInput) DescribeNotificationSubscriptionsRequest
}

// Send marshals and sends the DescribeNotificationSubscriptions API request.
func (r DescribeNotificationSubscriptionsRequest) Send(ctx context.Context) (*DescribeNotificationSubscriptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNotificationSubscriptionsResponse{
		DescribeNotificationSubscriptionsOutput: r.Request.Data.(*types.DescribeNotificationSubscriptionsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeNotificationSubscriptionsResponse is the response type for the
// DescribeNotificationSubscriptions API operation.
type DescribeNotificationSubscriptionsResponse struct {
	*types.DescribeNotificationSubscriptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNotificationSubscriptions request.
func (r *DescribeNotificationSubscriptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
