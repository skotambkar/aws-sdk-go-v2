// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
)

const opListPendingInvitationResources = "ListPendingInvitationResources"

// ListPendingInvitationResourcesRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Lists the resources in a resource share that is shared with you but that
// the invitation is still pending for.
//
//    // Example sending a request using ListPendingInvitationResourcesRequest.
//    req := client.ListPendingInvitationResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListPendingInvitationResources
func (c *Client) ListPendingInvitationResourcesRequest(input *types.ListPendingInvitationResourcesInput) ListPendingInvitationResourcesRequest {
	op := &aws.Operation{
		Name:       opListPendingInvitationResources,
		HTTPMethod: "POST",
		HTTPPath:   "/listpendinginvitationresources",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPendingInvitationResourcesInput{}
	}

	req := c.newRequest(op, input, &types.ListPendingInvitationResourcesOutput{})
	return ListPendingInvitationResourcesRequest{Request: req, Input: input, Copy: c.ListPendingInvitationResourcesRequest}
}

// ListPendingInvitationResourcesRequest is the request type for the
// ListPendingInvitationResources API operation.
type ListPendingInvitationResourcesRequest struct {
	*aws.Request
	Input *types.ListPendingInvitationResourcesInput
	Copy  func(*types.ListPendingInvitationResourcesInput) ListPendingInvitationResourcesRequest
}

// Send marshals and sends the ListPendingInvitationResources API request.
func (r ListPendingInvitationResourcesRequest) Send(ctx context.Context) (*ListPendingInvitationResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPendingInvitationResourcesResponse{
		ListPendingInvitationResourcesOutput: r.Request.Data.(*types.ListPendingInvitationResourcesOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPendingInvitationResourcesRequestPaginator returns a paginator for ListPendingInvitationResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPendingInvitationResourcesRequest(input)
//   p := ram.NewListPendingInvitationResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPendingInvitationResourcesPaginator(req ListPendingInvitationResourcesRequest) ListPendingInvitationResourcesPaginator {
	return ListPendingInvitationResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPendingInvitationResourcesInput
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

// ListPendingInvitationResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPendingInvitationResourcesPaginator struct {
	aws.Pager
}

func (p *ListPendingInvitationResourcesPaginator) CurrentPage() *types.ListPendingInvitationResourcesOutput {
	return p.Pager.CurrentPage().(*types.ListPendingInvitationResourcesOutput)
}

// ListPendingInvitationResourcesResponse is the response type for the
// ListPendingInvitationResources API operation.
type ListPendingInvitationResourcesResponse struct {
	*types.ListPendingInvitationResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPendingInvitationResources request.
func (r *ListPendingInvitationResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
