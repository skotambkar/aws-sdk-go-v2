// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
)

const opListDomains = "ListDomains"

// ListDomainsRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns the list of domains registered in the account. The results may be
// split into multiple pages. To retrieve subsequent pages, make the call again
// using the nextPageToken returned by the initial call.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains. The element must be set to arn:aws:swf::AccountID:domain/*,
//    where AccountID is the account ID, with no dashes.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using ListDomainsRequest.
//    req := client.ListDomainsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListDomainsRequest(input *types.ListDomainsInput) ListDomainsRequest {
	op := &aws.Operation{
		Name:       opListDomains,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextPageToken"},
			OutputTokens:    []string{"nextPageToken"},
			LimitToken:      "maximumPageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListDomainsInput{}
	}

	req := c.newRequest(op, input, &types.ListDomainsOutput{})
	return ListDomainsRequest{Request: req, Input: input, Copy: c.ListDomainsRequest}
}

// ListDomainsRequest is the request type for the
// ListDomains API operation.
type ListDomainsRequest struct {
	*aws.Request
	Input *types.ListDomainsInput
	Copy  func(*types.ListDomainsInput) ListDomainsRequest
}

// Send marshals and sends the ListDomains API request.
func (r ListDomainsRequest) Send(ctx context.Context) (*ListDomainsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDomainsResponse{
		ListDomainsOutput: r.Request.Data.(*types.ListDomainsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDomainsRequestPaginator returns a paginator for ListDomains.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDomainsRequest(input)
//   p := swf.NewListDomainsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDomainsPaginator(req ListDomainsRequest) ListDomainsPaginator {
	return ListDomainsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListDomainsInput
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

// ListDomainsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDomainsPaginator struct {
	aws.Pager
}

func (p *ListDomainsPaginator) CurrentPage() *types.ListDomainsOutput {
	return p.Pager.CurrentPage().(*types.ListDomainsOutput)
}

// ListDomainsResponse is the response type for the
// ListDomains API operation.
type ListDomainsResponse struct {
	*types.ListDomainsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDomains request.
func (r *ListDomainsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
