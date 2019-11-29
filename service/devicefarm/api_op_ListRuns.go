// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opListRuns = "ListRuns"

// ListRunsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about runs, given an AWS Device Farm project ARN.
//
//    // Example sending a request using ListRunsRequest.
//    req := client.ListRunsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListRuns
func (c *Client) ListRunsRequest(input *types.ListRunsInput) ListRunsRequest {
	op := &aws.Operation{
		Name:       opListRuns,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListRunsInput{}
	}

	req := c.newRequest(op, input, &types.ListRunsOutput{})
	return ListRunsRequest{Request: req, Input: input, Copy: c.ListRunsRequest}
}

// ListRunsRequest is the request type for the
// ListRuns API operation.
type ListRunsRequest struct {
	*aws.Request
	Input *types.ListRunsInput
	Copy  func(*types.ListRunsInput) ListRunsRequest
}

// Send marshals and sends the ListRuns API request.
func (r ListRunsRequest) Send(ctx context.Context) (*ListRunsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRunsResponse{
		ListRunsOutput: r.Request.Data.(*types.ListRunsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRunsRequestPaginator returns a paginator for ListRuns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRunsRequest(input)
//   p := devicefarm.NewListRunsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRunsPaginator(req ListRunsRequest) ListRunsPaginator {
	return ListRunsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListRunsInput
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

// ListRunsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRunsPaginator struct {
	aws.Pager
}

func (p *ListRunsPaginator) CurrentPage() *types.ListRunsOutput {
	return p.Pager.CurrentPage().(*types.ListRunsOutput)
}

// ListRunsResponse is the response type for the
// ListRuns API operation.
type ListRunsResponse struct {
	*types.ListRunsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRuns request.
func (r *ListRunsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
