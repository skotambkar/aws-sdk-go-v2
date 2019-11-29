// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
)

const opListResolverEndpoints = "ListResolverEndpoints"

// ListResolverEndpointsRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Lists all the resolver endpoints that were created using the current AWS
// account.
//
//    // Example sending a request using ListResolverEndpointsRequest.
//    req := client.ListResolverEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ListResolverEndpoints
func (c *Client) ListResolverEndpointsRequest(input *types.ListResolverEndpointsInput) ListResolverEndpointsRequest {
	op := &aws.Operation{
		Name:       opListResolverEndpoints,
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
		input = &types.ListResolverEndpointsInput{}
	}

	req := c.newRequest(op, input, &types.ListResolverEndpointsOutput{})
	return ListResolverEndpointsRequest{Request: req, Input: input, Copy: c.ListResolverEndpointsRequest}
}

// ListResolverEndpointsRequest is the request type for the
// ListResolverEndpoints API operation.
type ListResolverEndpointsRequest struct {
	*aws.Request
	Input *types.ListResolverEndpointsInput
	Copy  func(*types.ListResolverEndpointsInput) ListResolverEndpointsRequest
}

// Send marshals and sends the ListResolverEndpoints API request.
func (r ListResolverEndpointsRequest) Send(ctx context.Context) (*ListResolverEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResolverEndpointsResponse{
		ListResolverEndpointsOutput: r.Request.Data.(*types.ListResolverEndpointsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResolverEndpointsRequestPaginator returns a paginator for ListResolverEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResolverEndpointsRequest(input)
//   p := route53resolver.NewListResolverEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResolverEndpointsPaginator(req ListResolverEndpointsRequest) ListResolverEndpointsPaginator {
	return ListResolverEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListResolverEndpointsInput
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

// ListResolverEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResolverEndpointsPaginator struct {
	aws.Pager
}

func (p *ListResolverEndpointsPaginator) CurrentPage() *types.ListResolverEndpointsOutput {
	return p.Pager.CurrentPage().(*types.ListResolverEndpointsOutput)
}

// ListResolverEndpointsResponse is the response type for the
// ListResolverEndpoints API operation.
type ListResolverEndpointsResponse struct {
	*types.ListResolverEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResolverEndpoints request.
func (r *ListResolverEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
