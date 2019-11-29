// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
)

const opListExecutions = "ListExecutions"

// ListExecutionsRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Lists the executions of a state machine that meet the filtering criteria.
// Results are sorted by time, with the most recent execution first.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again
// using the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This operation is eventually consistent. The results are best effort and
// may not reflect very recent updates and changes.
//
//    // Example sending a request using ListExecutionsRequest.
//    req := client.ListExecutionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/ListExecutions
func (c *Client) ListExecutionsRequest(input *types.ListExecutionsInput) ListExecutionsRequest {
	op := &aws.Operation{
		Name:       opListExecutions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListExecutionsInput{}
	}

	req := c.newRequest(op, input, &types.ListExecutionsOutput{})
	return ListExecutionsRequest{Request: req, Input: input, Copy: c.ListExecutionsRequest}
}

// ListExecutionsRequest is the request type for the
// ListExecutions API operation.
type ListExecutionsRequest struct {
	*aws.Request
	Input *types.ListExecutionsInput
	Copy  func(*types.ListExecutionsInput) ListExecutionsRequest
}

// Send marshals and sends the ListExecutions API request.
func (r ListExecutionsRequest) Send(ctx context.Context) (*ListExecutionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListExecutionsResponse{
		ListExecutionsOutput: r.Request.Data.(*types.ListExecutionsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListExecutionsRequestPaginator returns a paginator for ListExecutions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListExecutionsRequest(input)
//   p := sfn.NewListExecutionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListExecutionsPaginator(req ListExecutionsRequest) ListExecutionsPaginator {
	return ListExecutionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListExecutionsInput
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

// ListExecutionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListExecutionsPaginator struct {
	aws.Pager
}

func (p *ListExecutionsPaginator) CurrentPage() *types.ListExecutionsOutput {
	return p.Pager.CurrentPage().(*types.ListExecutionsOutput)
}

// ListExecutionsResponse is the response type for the
// ListExecutions API operation.
type ListExecutionsResponse struct {
	*types.ListExecutionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListExecutions request.
func (r *ListExecutionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
