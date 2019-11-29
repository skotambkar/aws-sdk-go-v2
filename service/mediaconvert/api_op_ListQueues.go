// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
)

const opListQueues = "ListQueues"

// ListQueuesRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Retrieve a JSON array of up to twenty of your queues. This will return the
// queues themselves, not just a list of them. To retrieve the next twenty queues,
// use the nextToken string returned with the array.
//
//    // Example sending a request using ListQueuesRequest.
//    req := client.ListQueuesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/ListQueues
func (c *Client) ListQueuesRequest(input *types.ListQueuesInput) ListQueuesRequest {
	op := &aws.Operation{
		Name:       opListQueues,
		HTTPMethod: "GET",
		HTTPPath:   "/2017-08-29/queues",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListQueuesInput{}
	}

	req := c.newRequest(op, input, &types.ListQueuesOutput{})
	return ListQueuesRequest{Request: req, Input: input, Copy: c.ListQueuesRequest}
}

// ListQueuesRequest is the request type for the
// ListQueues API operation.
type ListQueuesRequest struct {
	*aws.Request
	Input *types.ListQueuesInput
	Copy  func(*types.ListQueuesInput) ListQueuesRequest
}

// Send marshals and sends the ListQueues API request.
func (r ListQueuesRequest) Send(ctx context.Context) (*ListQueuesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListQueuesResponse{
		ListQueuesOutput: r.Request.Data.(*types.ListQueuesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListQueuesRequestPaginator returns a paginator for ListQueues.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListQueuesRequest(input)
//   p := mediaconvert.NewListQueuesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListQueuesPaginator(req ListQueuesRequest) ListQueuesPaginator {
	return ListQueuesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListQueuesInput
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

// ListQueuesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListQueuesPaginator struct {
	aws.Pager
}

func (p *ListQueuesPaginator) CurrentPage() *types.ListQueuesOutput {
	return p.Pager.CurrentPage().(*types.ListQueuesOutput)
}

// ListQueuesResponse is the response type for the
// ListQueues API operation.
type ListQueuesResponse struct {
	*types.ListQueuesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListQueues request.
func (r *ListQueuesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
