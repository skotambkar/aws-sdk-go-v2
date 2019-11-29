// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opListSuites = "ListSuites"

// ListSuitesRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about test suites for a given job.
//
//    // Example sending a request using ListSuitesRequest.
//    req := client.ListSuitesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListSuites
func (c *Client) ListSuitesRequest(input *types.ListSuitesInput) ListSuitesRequest {
	op := &aws.Operation{
		Name:       opListSuites,
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
		input = &types.ListSuitesInput{}
	}

	req := c.newRequest(op, input, &types.ListSuitesOutput{})
	return ListSuitesRequest{Request: req, Input: input, Copy: c.ListSuitesRequest}
}

// ListSuitesRequest is the request type for the
// ListSuites API operation.
type ListSuitesRequest struct {
	*aws.Request
	Input *types.ListSuitesInput
	Copy  func(*types.ListSuitesInput) ListSuitesRequest
}

// Send marshals and sends the ListSuites API request.
func (r ListSuitesRequest) Send(ctx context.Context) (*ListSuitesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSuitesResponse{
		ListSuitesOutput: r.Request.Data.(*types.ListSuitesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSuitesRequestPaginator returns a paginator for ListSuites.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSuitesRequest(input)
//   p := devicefarm.NewListSuitesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSuitesPaginator(req ListSuitesRequest) ListSuitesPaginator {
	return ListSuitesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListSuitesInput
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

// ListSuitesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSuitesPaginator struct {
	aws.Pager
}

func (p *ListSuitesPaginator) CurrentPage() *types.ListSuitesOutput {
	return p.Pager.CurrentPage().(*types.ListSuitesOutput)
}

// ListSuitesResponse is the response type for the
// ListSuites API operation.
type ListSuitesResponse struct {
	*types.ListSuitesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSuites request.
func (r *ListSuitesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
