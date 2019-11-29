// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
)

const opListResources = "ListResources"

// ListResourcesRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Lists the resources registered to be managed by the Data Catalog.
//
//    // Example sending a request using ListResourcesRequest.
//    req := client.ListResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/ListResources
func (c *Client) ListResourcesRequest(input *types.ListResourcesInput) ListResourcesRequest {
	op := &aws.Operation{
		Name:       opListResources,
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
		input = &types.ListResourcesInput{}
	}

	req := c.newRequest(op, input, &types.ListResourcesOutput{})
	return ListResourcesRequest{Request: req, Input: input, Copy: c.ListResourcesRequest}
}

// ListResourcesRequest is the request type for the
// ListResources API operation.
type ListResourcesRequest struct {
	*aws.Request
	Input *types.ListResourcesInput
	Copy  func(*types.ListResourcesInput) ListResourcesRequest
}

// Send marshals and sends the ListResources API request.
func (r ListResourcesRequest) Send(ctx context.Context) (*ListResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResourcesResponse{
		ListResourcesOutput: r.Request.Data.(*types.ListResourcesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResourcesRequestPaginator returns a paginator for ListResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResourcesRequest(input)
//   p := lakeformation.NewListResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResourcesPaginator(req ListResourcesRequest) ListResourcesPaginator {
	return ListResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListResourcesInput
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

// ListResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResourcesPaginator struct {
	aws.Pager
}

func (p *ListResourcesPaginator) CurrentPage() *types.ListResourcesOutput {
	return p.Pager.CurrentPage().(*types.ListResourcesOutput)
}

// ListResourcesResponse is the response type for the
// ListResources API operation.
type ListResourcesResponse struct {
	*types.ListResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResources request.
func (r *ListResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
