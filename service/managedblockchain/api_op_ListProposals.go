// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package managedblockchain

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
)

const opListProposals = "ListProposals"

// ListProposalsRequest returns a request value for making API operation for
// Amazon Managed Blockchain.
//
// Returns a listing of proposals for the network.
//
//    // Example sending a request using ListProposalsRequest.
//    req := client.ListProposalsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/managedblockchain-2018-09-24/ListProposals
func (c *Client) ListProposalsRequest(input *types.ListProposalsInput) ListProposalsRequest {
	op := &aws.Operation{
		Name:       opListProposals,
		HTTPMethod: "GET",
		HTTPPath:   "/networks/{networkId}/proposals",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListProposalsInput{}
	}

	req := c.newRequest(op, input, &types.ListProposalsOutput{})
	return ListProposalsRequest{Request: req, Input: input, Copy: c.ListProposalsRequest}
}

// ListProposalsRequest is the request type for the
// ListProposals API operation.
type ListProposalsRequest struct {
	*aws.Request
	Input *types.ListProposalsInput
	Copy  func(*types.ListProposalsInput) ListProposalsRequest
}

// Send marshals and sends the ListProposals API request.
func (r ListProposalsRequest) Send(ctx context.Context) (*ListProposalsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProposalsResponse{
		ListProposalsOutput: r.Request.Data.(*types.ListProposalsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProposalsRequestPaginator returns a paginator for ListProposals.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProposalsRequest(input)
//   p := managedblockchain.NewListProposalsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProposalsPaginator(req ListProposalsRequest) ListProposalsPaginator {
	return ListProposalsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListProposalsInput
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

// ListProposalsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProposalsPaginator struct {
	aws.Pager
}

func (p *ListProposalsPaginator) CurrentPage() *types.ListProposalsOutput {
	return p.Pager.CurrentPage().(*types.ListProposalsOutput)
}

// ListProposalsResponse is the response type for the
// ListProposals API operation.
type ListProposalsResponse struct {
	*types.ListProposalsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProposals request.
func (r *ListProposalsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
