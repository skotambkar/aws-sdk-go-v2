// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
)

const opListEntitlements = "ListEntitlements"

// ListEntitlementsRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Displays a list of all entitlements that have been granted to this account.
// This request returns 20 results per page.
//
//    // Example sending a request using ListEntitlementsRequest.
//    req := client.ListEntitlementsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/ListEntitlements
func (c *Client) ListEntitlementsRequest(input *types.ListEntitlementsInput) ListEntitlementsRequest {
	op := &aws.Operation{
		Name:       opListEntitlements,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/entitlements",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListEntitlementsInput{}
	}

	req := c.newRequest(op, input, &types.ListEntitlementsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListEntitlementsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListEntitlementsRequest{Request: req, Input: input, Copy: c.ListEntitlementsRequest}
}

// ListEntitlementsRequest is the request type for the
// ListEntitlements API operation.
type ListEntitlementsRequest struct {
	*aws.Request
	Input *types.ListEntitlementsInput
	Copy  func(*types.ListEntitlementsInput) ListEntitlementsRequest
}

// Send marshals and sends the ListEntitlements API request.
func (r ListEntitlementsRequest) Send(ctx context.Context) (*ListEntitlementsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEntitlementsResponse{
		ListEntitlementsOutput: r.Request.Data.(*types.ListEntitlementsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEntitlementsRequestPaginator returns a paginator for ListEntitlements.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEntitlementsRequest(input)
//   p := mediaconnect.NewListEntitlementsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEntitlementsPaginator(req ListEntitlementsRequest) ListEntitlementsPaginator {
	return ListEntitlementsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListEntitlementsInput
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

// ListEntitlementsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEntitlementsPaginator struct {
	aws.Pager
}

func (p *ListEntitlementsPaginator) CurrentPage() *types.ListEntitlementsOutput {
	return p.Pager.CurrentPage().(*types.ListEntitlementsOutput)
}

// ListEntitlementsResponse is the response type for the
// ListEntitlements API operation.
type ListEntitlementsResponse struct {
	*types.ListEntitlementsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEntitlements request.
func (r *ListEntitlementsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
