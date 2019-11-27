// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackagevod

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod/types"
)

const opListAssets = "ListAssets"

// ListAssetsRequest returns a request value for making API operation for
// AWS Elemental MediaPackage VOD.
//
// Returns a collection of MediaPackage VOD Asset resources.
//
//    // Example sending a request using ListAssetsRequest.
//    req := client.ListAssetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-vod-2018-11-07/ListAssets
func (c *Client) ListAssetsRequest(input *types.ListAssetsInput) ListAssetsRequest {
	op := &aws.Operation{
		Name:       opListAssets,
		HTTPMethod: "GET",
		HTTPPath:   "/assets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListAssetsInput{}
	}

	req := c.newRequest(op, input, &types.ListAssetsOutput{})
	return ListAssetsRequest{Request: req, Input: input, Copy: c.ListAssetsRequest}
}

// ListAssetsRequest is the request type for the
// ListAssets API operation.
type ListAssetsRequest struct {
	*aws.Request
	Input *types.ListAssetsInput
	Copy  func(*types.ListAssetsInput) ListAssetsRequest
}

// Send marshals and sends the ListAssets API request.
func (r ListAssetsRequest) Send(ctx context.Context) (*ListAssetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAssetsResponse{
		ListAssetsOutput: r.Request.Data.(*types.ListAssetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAssetsRequestPaginator returns a paginator for ListAssets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAssetsRequest(input)
//   p := mediapackagevod.NewListAssetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAssetsPaginator(req ListAssetsRequest) ListAssetsPaginator {
	return ListAssetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListAssetsInput
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

// ListAssetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAssetsPaginator struct {
	aws.Pager
}

func (p *ListAssetsPaginator) CurrentPage() *types.ListAssetsOutput {
	return p.Pager.CurrentPage().(*types.ListAssetsOutput)
}

// ListAssetsResponse is the response type for the
// ListAssets API operation.
type ListAssetsResponse struct {
	*types.ListAssetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAssets request.
func (r *ListAssetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
