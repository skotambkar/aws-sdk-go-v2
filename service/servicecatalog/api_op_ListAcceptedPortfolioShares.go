// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opListAcceptedPortfolioShares = "ListAcceptedPortfolioShares"

// ListAcceptedPortfolioSharesRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists all portfolios for which sharing was accepted by this account.
//
//    // Example sending a request using ListAcceptedPortfolioSharesRequest.
//    req := client.ListAcceptedPortfolioSharesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListAcceptedPortfolioShares
func (c *Client) ListAcceptedPortfolioSharesRequest(input *types.ListAcceptedPortfolioSharesInput) ListAcceptedPortfolioSharesRequest {
	op := &aws.Operation{
		Name:       opListAcceptedPortfolioShares,
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
		input = &types.ListAcceptedPortfolioSharesInput{}
	}

	req := c.newRequest(op, input, &types.ListAcceptedPortfolioSharesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListAcceptedPortfolioSharesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListAcceptedPortfolioSharesRequest{Request: req, Input: input, Copy: c.ListAcceptedPortfolioSharesRequest}
}

// ListAcceptedPortfolioSharesRequest is the request type for the
// ListAcceptedPortfolioShares API operation.
type ListAcceptedPortfolioSharesRequest struct {
	*aws.Request
	Input *types.ListAcceptedPortfolioSharesInput
	Copy  func(*types.ListAcceptedPortfolioSharesInput) ListAcceptedPortfolioSharesRequest
}

// Send marshals and sends the ListAcceptedPortfolioShares API request.
func (r ListAcceptedPortfolioSharesRequest) Send(ctx context.Context) (*ListAcceptedPortfolioSharesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAcceptedPortfolioSharesResponse{
		ListAcceptedPortfolioSharesOutput: r.Request.Data.(*types.ListAcceptedPortfolioSharesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAcceptedPortfolioSharesRequestPaginator returns a paginator for ListAcceptedPortfolioShares.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAcceptedPortfolioSharesRequest(input)
//   p := servicecatalog.NewListAcceptedPortfolioSharesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAcceptedPortfolioSharesPaginator(req ListAcceptedPortfolioSharesRequest) ListAcceptedPortfolioSharesPaginator {
	return ListAcceptedPortfolioSharesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListAcceptedPortfolioSharesInput
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

// ListAcceptedPortfolioSharesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAcceptedPortfolioSharesPaginator struct {
	aws.Pager
}

func (p *ListAcceptedPortfolioSharesPaginator) CurrentPage() *types.ListAcceptedPortfolioSharesOutput {
	return p.Pager.CurrentPage().(*types.ListAcceptedPortfolioSharesOutput)
}

// ListAcceptedPortfolioSharesResponse is the response type for the
// ListAcceptedPortfolioShares API operation.
type ListAcceptedPortfolioSharesResponse struct {
	*types.ListAcceptedPortfolioSharesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAcceptedPortfolioShares request.
func (r *ListAcceptedPortfolioSharesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
