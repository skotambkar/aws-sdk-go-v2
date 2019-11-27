// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

const opListUpdates = "ListUpdates"

// ListUpdatesRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Lists the updates associated with an Amazon EKS cluster in your AWS account,
// in the specified Region.
//
//    // Example sending a request using ListUpdatesRequest.
//    req := client.ListUpdatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/ListUpdates
func (c *Client) ListUpdatesRequest(input *types.ListUpdatesInput) ListUpdatesRequest {
	op := &aws.Operation{
		Name:       opListUpdates,
		HTTPMethod: "GET",
		HTTPPath:   "/clusters/{name}/updates",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListUpdatesInput{}
	}

	req := c.newRequest(op, input, &types.ListUpdatesOutput{})
	return ListUpdatesRequest{Request: req, Input: input, Copy: c.ListUpdatesRequest}
}

// ListUpdatesRequest is the request type for the
// ListUpdates API operation.
type ListUpdatesRequest struct {
	*aws.Request
	Input *types.ListUpdatesInput
	Copy  func(*types.ListUpdatesInput) ListUpdatesRequest
}

// Send marshals and sends the ListUpdates API request.
func (r ListUpdatesRequest) Send(ctx context.Context) (*ListUpdatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUpdatesResponse{
		ListUpdatesOutput: r.Request.Data.(*types.ListUpdatesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListUpdatesRequestPaginator returns a paginator for ListUpdates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListUpdatesRequest(input)
//   p := eks.NewListUpdatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListUpdatesPaginator(req ListUpdatesRequest) ListUpdatesPaginator {
	return ListUpdatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListUpdatesInput
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

// ListUpdatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListUpdatesPaginator struct {
	aws.Pager
}

func (p *ListUpdatesPaginator) CurrentPage() *types.ListUpdatesOutput {
	return p.Pager.CurrentPage().(*types.ListUpdatesOutput)
}

// ListUpdatesResponse is the response type for the
// ListUpdates API operation.
type ListUpdatesResponse struct {
	*types.ListUpdatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUpdates request.
func (r *ListUpdatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
