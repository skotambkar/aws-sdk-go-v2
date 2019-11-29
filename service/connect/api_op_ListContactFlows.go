// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
)

const opListContactFlows = "ListContactFlows"

// ListContactFlowsRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Provides information about the contact flows for the specified Amazon Connect
// instance.
//
//    // Example sending a request using ListContactFlowsRequest.
//    req := client.ListContactFlowsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/ListContactFlows
func (c *Client) ListContactFlowsRequest(input *types.ListContactFlowsInput) ListContactFlowsRequest {
	op := &aws.Operation{
		Name:       opListContactFlows,
		HTTPMethod: "GET",
		HTTPPath:   "/contact-flows-summary/{InstanceId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListContactFlowsInput{}
	}

	req := c.newRequest(op, input, &types.ListContactFlowsOutput{})
	return ListContactFlowsRequest{Request: req, Input: input, Copy: c.ListContactFlowsRequest}
}

// ListContactFlowsRequest is the request type for the
// ListContactFlows API operation.
type ListContactFlowsRequest struct {
	*aws.Request
	Input *types.ListContactFlowsInput
	Copy  func(*types.ListContactFlowsInput) ListContactFlowsRequest
}

// Send marshals and sends the ListContactFlows API request.
func (r ListContactFlowsRequest) Send(ctx context.Context) (*ListContactFlowsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListContactFlowsResponse{
		ListContactFlowsOutput: r.Request.Data.(*types.ListContactFlowsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListContactFlowsRequestPaginator returns a paginator for ListContactFlows.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListContactFlowsRequest(input)
//   p := connect.NewListContactFlowsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListContactFlowsPaginator(req ListContactFlowsRequest) ListContactFlowsPaginator {
	return ListContactFlowsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListContactFlowsInput
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

// ListContactFlowsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListContactFlowsPaginator struct {
	aws.Pager
}

func (p *ListContactFlowsPaginator) CurrentPage() *types.ListContactFlowsOutput {
	return p.Pager.CurrentPage().(*types.ListContactFlowsOutput)
}

// ListContactFlowsResponse is the response type for the
// ListContactFlows API operation.
type ListContactFlowsResponse struct {
	*types.ListContactFlowsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListContactFlows request.
func (r *ListContactFlowsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
