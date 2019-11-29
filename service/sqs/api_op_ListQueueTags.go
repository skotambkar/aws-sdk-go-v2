// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

const opListQueueTags = "ListQueueTags"

// ListQueueTagsRequest returns a request value for making API operation for
// Amazon Simple Queue Service.
//
// List all cost allocation tags added to the specified Amazon SQS queue. For
// an overview, see Tagging Your Amazon SQS Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
//    // Example sending a request using ListQueueTagsRequest.
//    req := client.ListQueueTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ListQueueTags
func (c *Client) ListQueueTagsRequest(input *types.ListQueueTagsInput) ListQueueTagsRequest {
	op := &aws.Operation{
		Name:       opListQueueTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListQueueTagsInput{}
	}

	req := c.newRequest(op, input, &types.ListQueueTagsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.ListQueueTagsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListQueueTagsRequest{Request: req, Input: input, Copy: c.ListQueueTagsRequest}
}

// ListQueueTagsRequest is the request type for the
// ListQueueTags API operation.
type ListQueueTagsRequest struct {
	*aws.Request
	Input *types.ListQueueTagsInput
	Copy  func(*types.ListQueueTagsInput) ListQueueTagsRequest
}

// Send marshals and sends the ListQueueTags API request.
func (r ListQueueTagsRequest) Send(ctx context.Context) (*ListQueueTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListQueueTagsResponse{
		ListQueueTagsOutput: r.Request.Data.(*types.ListQueueTagsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListQueueTagsResponse is the response type for the
// ListQueueTags API operation.
type ListQueueTagsResponse struct {
	*types.ListQueueTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListQueueTags request.
func (r *ListQueueTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
