// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
)

const opSearchFlowTemplates = "SearchFlowTemplates"

// SearchFlowTemplatesRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Searches for summary information about workflows.
//
//    // Example sending a request using SearchFlowTemplatesRequest.
//    req := client.SearchFlowTemplatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/SearchFlowTemplates
func (c *Client) SearchFlowTemplatesRequest(input *types.SearchFlowTemplatesInput) SearchFlowTemplatesRequest {
	op := &aws.Operation{
		Name:       opSearchFlowTemplates,
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
		input = &types.SearchFlowTemplatesInput{}
	}

	req := c.newRequest(op, input, &types.SearchFlowTemplatesOutput{})
	return SearchFlowTemplatesRequest{Request: req, Input: input, Copy: c.SearchFlowTemplatesRequest}
}

// SearchFlowTemplatesRequest is the request type for the
// SearchFlowTemplates API operation.
type SearchFlowTemplatesRequest struct {
	*aws.Request
	Input *types.SearchFlowTemplatesInput
	Copy  func(*types.SearchFlowTemplatesInput) SearchFlowTemplatesRequest
}

// Send marshals and sends the SearchFlowTemplates API request.
func (r SearchFlowTemplatesRequest) Send(ctx context.Context) (*SearchFlowTemplatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchFlowTemplatesResponse{
		SearchFlowTemplatesOutput: r.Request.Data.(*types.SearchFlowTemplatesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchFlowTemplatesRequestPaginator returns a paginator for SearchFlowTemplates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchFlowTemplatesRequest(input)
//   p := iotthingsgraph.NewSearchFlowTemplatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchFlowTemplatesPaginator(req SearchFlowTemplatesRequest) SearchFlowTemplatesPaginator {
	return SearchFlowTemplatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.SearchFlowTemplatesInput
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

// SearchFlowTemplatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchFlowTemplatesPaginator struct {
	aws.Pager
}

func (p *SearchFlowTemplatesPaginator) CurrentPage() *types.SearchFlowTemplatesOutput {
	return p.Pager.CurrentPage().(*types.SearchFlowTemplatesOutput)
}

// SearchFlowTemplatesResponse is the response type for the
// SearchFlowTemplates API operation.
type SearchFlowTemplatesResponse struct {
	*types.SearchFlowTemplatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchFlowTemplates request.
func (r *SearchFlowTemplatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
