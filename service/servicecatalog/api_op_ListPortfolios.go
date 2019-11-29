// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opListPortfolios = "ListPortfolios"

// ListPortfoliosRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists all portfolios in the catalog.
//
//    // Example sending a request using ListPortfoliosRequest.
//    req := client.ListPortfoliosRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListPortfolios
func (c *Client) ListPortfoliosRequest(input *types.ListPortfoliosInput) ListPortfoliosRequest {
	op := &aws.Operation{
		Name:       opListPortfolios,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPortfoliosInput{}
	}

	req := c.newRequest(op, input, &types.ListPortfoliosOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListPortfoliosMarshaler{Input: input}.GetNamedBuildHandler())

	return ListPortfoliosRequest{Request: req, Input: input, Copy: c.ListPortfoliosRequest}
}

// ListPortfoliosRequest is the request type for the
// ListPortfolios API operation.
type ListPortfoliosRequest struct {
	*aws.Request
	Input *types.ListPortfoliosInput
	Copy  func(*types.ListPortfoliosInput) ListPortfoliosRequest
}

// Send marshals and sends the ListPortfolios API request.
func (r ListPortfoliosRequest) Send(ctx context.Context) (*ListPortfoliosResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPortfoliosResponse{
		ListPortfoliosOutput: r.Request.Data.(*types.ListPortfoliosOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPortfoliosRequestPaginator returns a paginator for ListPortfolios.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPortfoliosRequest(input)
//   p := servicecatalog.NewListPortfoliosRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPortfoliosPaginator(req ListPortfoliosRequest) ListPortfoliosPaginator {
	return ListPortfoliosPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPortfoliosInput
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

// ListPortfoliosPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPortfoliosPaginator struct {
	aws.Pager
}

func (p *ListPortfoliosPaginator) CurrentPage() *types.ListPortfoliosOutput {
	return p.Pager.CurrentPage().(*types.ListPortfoliosOutput)
}

// ListPortfoliosResponse is the response type for the
// ListPortfolios API operation.
type ListPortfoliosResponse struct {
	*types.ListPortfoliosOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPortfolios request.
func (r *ListPortfoliosResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
