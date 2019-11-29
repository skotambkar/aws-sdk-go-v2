// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sns/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

const opListSubscriptionsByTopic = "ListSubscriptionsByTopic"

// ListSubscriptionsByTopicRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Returns a list of the subscriptions to a specific topic. Each call returns
// a limited list of subscriptions, up to 100. If there are more subscriptions,
// a NextToken is also returned. Use the NextToken parameter in a new ListSubscriptionsByTopic
// call to get further results.
//
// This action is throttled at 30 transactions per second (TPS).
//
//    // Example sending a request using ListSubscriptionsByTopicRequest.
//    req := client.ListSubscriptionsByTopicRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/ListSubscriptionsByTopic
func (c *Client) ListSubscriptionsByTopicRequest(input *types.ListSubscriptionsByTopicInput) ListSubscriptionsByTopicRequest {
	op := &aws.Operation{
		Name:       opListSubscriptionsByTopic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListSubscriptionsByTopicInput{}
	}

	req := c.newRequest(op, input, &types.ListSubscriptionsByTopicOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.ListSubscriptionsByTopicMarshaler{Input: input}.GetNamedBuildHandler())

	return ListSubscriptionsByTopicRequest{Request: req, Input: input, Copy: c.ListSubscriptionsByTopicRequest}
}

// ListSubscriptionsByTopicRequest is the request type for the
// ListSubscriptionsByTopic API operation.
type ListSubscriptionsByTopicRequest struct {
	*aws.Request
	Input *types.ListSubscriptionsByTopicInput
	Copy  func(*types.ListSubscriptionsByTopicInput) ListSubscriptionsByTopicRequest
}

// Send marshals and sends the ListSubscriptionsByTopic API request.
func (r ListSubscriptionsByTopicRequest) Send(ctx context.Context) (*ListSubscriptionsByTopicResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSubscriptionsByTopicResponse{
		ListSubscriptionsByTopicOutput: r.Request.Data.(*types.ListSubscriptionsByTopicOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSubscriptionsByTopicRequestPaginator returns a paginator for ListSubscriptionsByTopic.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSubscriptionsByTopicRequest(input)
//   p := sns.NewListSubscriptionsByTopicRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSubscriptionsByTopicPaginator(req ListSubscriptionsByTopicRequest) ListSubscriptionsByTopicPaginator {
	return ListSubscriptionsByTopicPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListSubscriptionsByTopicInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListSubscriptionsByTopicPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSubscriptionsByTopicPaginator struct {
	aws.Pager
}

func (p *ListSubscriptionsByTopicPaginator) CurrentPage() *types.ListSubscriptionsByTopicOutput {
	return p.Pager.CurrentPage().(*types.ListSubscriptionsByTopicOutput)
}

// ListSubscriptionsByTopicResponse is the response type for the
// ListSubscriptionsByTopic API operation.
type ListSubscriptionsByTopicResponse struct {
	*types.ListSubscriptionsByTopicOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSubscriptionsByTopic request.
func (r *ListSubscriptionsByTopicResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
