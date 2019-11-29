// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
)

const opListResolverRuleAssociations = "ListResolverRuleAssociations"

// ListResolverRuleAssociationsRequest returns a request value for making API operation for
// Amazon Route 53 Resolver.
//
// Lists the associations that were created between resolver rules and VPCs
// using the current AWS account.
//
//    // Example sending a request using ListResolverRuleAssociationsRequest.
//    req := client.ListResolverRuleAssociationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53resolver-2018-04-01/ListResolverRuleAssociations
func (c *Client) ListResolverRuleAssociationsRequest(input *types.ListResolverRuleAssociationsInput) ListResolverRuleAssociationsRequest {
	op := &aws.Operation{
		Name:       opListResolverRuleAssociations,
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
		input = &types.ListResolverRuleAssociationsInput{}
	}

	req := c.newRequest(op, input, &types.ListResolverRuleAssociationsOutput{})
	return ListResolverRuleAssociationsRequest{Request: req, Input: input, Copy: c.ListResolverRuleAssociationsRequest}
}

// ListResolverRuleAssociationsRequest is the request type for the
// ListResolverRuleAssociations API operation.
type ListResolverRuleAssociationsRequest struct {
	*aws.Request
	Input *types.ListResolverRuleAssociationsInput
	Copy  func(*types.ListResolverRuleAssociationsInput) ListResolverRuleAssociationsRequest
}

// Send marshals and sends the ListResolverRuleAssociations API request.
func (r ListResolverRuleAssociationsRequest) Send(ctx context.Context) (*ListResolverRuleAssociationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListResolverRuleAssociationsResponse{
		ListResolverRuleAssociationsOutput: r.Request.Data.(*types.ListResolverRuleAssociationsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListResolverRuleAssociationsRequestPaginator returns a paginator for ListResolverRuleAssociations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListResolverRuleAssociationsRequest(input)
//   p := route53resolver.NewListResolverRuleAssociationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListResolverRuleAssociationsPaginator(req ListResolverRuleAssociationsRequest) ListResolverRuleAssociationsPaginator {
	return ListResolverRuleAssociationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListResolverRuleAssociationsInput
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

// ListResolverRuleAssociationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListResolverRuleAssociationsPaginator struct {
	aws.Pager
}

func (p *ListResolverRuleAssociationsPaginator) CurrentPage() *types.ListResolverRuleAssociationsOutput {
	return p.Pager.CurrentPage().(*types.ListResolverRuleAssociationsOutput)
}

// ListResolverRuleAssociationsResponse is the response type for the
// ListResolverRuleAssociations API operation.
type ListResolverRuleAssociationsResponse struct {
	*types.ListResolverRuleAssociationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListResolverRuleAssociations request.
func (r *ListResolverRuleAssociationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
