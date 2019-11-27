// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
)

const opListWebsiteCertificateAuthorities = "ListWebsiteCertificateAuthorities"

// ListWebsiteCertificateAuthoritiesRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Retrieves a list of certificate authorities added for the current account
// and Region.
//
//    // Example sending a request using ListWebsiteCertificateAuthoritiesRequest.
//    req := client.ListWebsiteCertificateAuthoritiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/ListWebsiteCertificateAuthorities
func (c *Client) ListWebsiteCertificateAuthoritiesRequest(input *types.ListWebsiteCertificateAuthoritiesInput) ListWebsiteCertificateAuthoritiesRequest {
	op := &aws.Operation{
		Name:       opListWebsiteCertificateAuthorities,
		HTTPMethod: "POST",
		HTTPPath:   "/listWebsiteCertificateAuthorities",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListWebsiteCertificateAuthoritiesInput{}
	}

	req := c.newRequest(op, input, &types.ListWebsiteCertificateAuthoritiesOutput{})
	return ListWebsiteCertificateAuthoritiesRequest{Request: req, Input: input, Copy: c.ListWebsiteCertificateAuthoritiesRequest}
}

// ListWebsiteCertificateAuthoritiesRequest is the request type for the
// ListWebsiteCertificateAuthorities API operation.
type ListWebsiteCertificateAuthoritiesRequest struct {
	*aws.Request
	Input *types.ListWebsiteCertificateAuthoritiesInput
	Copy  func(*types.ListWebsiteCertificateAuthoritiesInput) ListWebsiteCertificateAuthoritiesRequest
}

// Send marshals and sends the ListWebsiteCertificateAuthorities API request.
func (r ListWebsiteCertificateAuthoritiesRequest) Send(ctx context.Context) (*ListWebsiteCertificateAuthoritiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListWebsiteCertificateAuthoritiesResponse{
		ListWebsiteCertificateAuthoritiesOutput: r.Request.Data.(*types.ListWebsiteCertificateAuthoritiesOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListWebsiteCertificateAuthoritiesRequestPaginator returns a paginator for ListWebsiteCertificateAuthorities.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListWebsiteCertificateAuthoritiesRequest(input)
//   p := worklink.NewListWebsiteCertificateAuthoritiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListWebsiteCertificateAuthoritiesPaginator(req ListWebsiteCertificateAuthoritiesRequest) ListWebsiteCertificateAuthoritiesPaginator {
	return ListWebsiteCertificateAuthoritiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListWebsiteCertificateAuthoritiesInput
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

// ListWebsiteCertificateAuthoritiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListWebsiteCertificateAuthoritiesPaginator struct {
	aws.Pager
}

func (p *ListWebsiteCertificateAuthoritiesPaginator) CurrentPage() *types.ListWebsiteCertificateAuthoritiesOutput {
	return p.Pager.CurrentPage().(*types.ListWebsiteCertificateAuthoritiesOutput)
}

// ListWebsiteCertificateAuthoritiesResponse is the response type for the
// ListWebsiteCertificateAuthorities API operation.
type ListWebsiteCertificateAuthoritiesResponse struct {
	*types.ListWebsiteCertificateAuthoritiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListWebsiteCertificateAuthorities request.
func (r *ListWebsiteCertificateAuthoritiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
