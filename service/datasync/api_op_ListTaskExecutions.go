// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
)

const opListTaskExecutions = "ListTaskExecutions"

// ListTaskExecutionsRequest returns a request value for making API operation for
// AWS DataSync.
//
// Returns a list of executed tasks.
//
//    // Example sending a request using ListTaskExecutionsRequest.
//    req := client.ListTaskExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/ListTaskExecutions
func (c *Client) ListTaskExecutionsRequest(input *types.ListTaskExecutionsInput) ListTaskExecutionsRequest {
	op := &aws.Operation{
		Name:       opListTaskExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListTaskExecutionsInput{}
	}

	req := c.newRequest(op, input, &types.ListTaskExecutionsOutput{})
	return ListTaskExecutionsRequest{Request: req, Input: input, Copy: c.ListTaskExecutionsRequest}
}

// ListTaskExecutionsRequest is the request type for the
// ListTaskExecutions API operation.
type ListTaskExecutionsRequest struct {
	*aws.Request
	Input *types.ListTaskExecutionsInput
	Copy  func(*types.ListTaskExecutionsInput) ListTaskExecutionsRequest
}

// Send marshals and sends the ListTaskExecutions API request.
func (r ListTaskExecutionsRequest) Send(ctx context.Context) (*ListTaskExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTaskExecutionsResponse{
		ListTaskExecutionsOutput: r.Request.Data.(*types.ListTaskExecutionsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTaskExecutionsRequestPaginator returns a paginator for ListTaskExecutions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTaskExecutionsRequest(input)
//   p := datasync.NewListTaskExecutionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTaskExecutionsPaginator(req ListTaskExecutionsRequest) ListTaskExecutionsPaginator {
	return ListTaskExecutionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTaskExecutionsInput
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

// ListTaskExecutionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTaskExecutionsPaginator struct {
	aws.Pager
}

func (p *ListTaskExecutionsPaginator) CurrentPage() *types.ListTaskExecutionsOutput {
	return p.Pager.CurrentPage().(*types.ListTaskExecutionsOutput)
}

// ListTaskExecutionsResponse is the response type for the
// ListTaskExecutions API operation.
type ListTaskExecutionsResponse struct {
	*types.ListTaskExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTaskExecutions request.
func (r *ListTaskExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
