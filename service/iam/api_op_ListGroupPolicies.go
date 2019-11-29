// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opListGroupPolicies = "ListGroupPolicies"

// ListGroupPoliciesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the names of the inline policies that are embedded in the specified
// IAM group.
//
// An IAM group can also have managed policies attached to it. To list the managed
// policies that are attached to a group, use ListAttachedGroupPolicies. For
// more information about policies, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. If
// there are no inline policies embedded with the specified group, the operation
// returns an empty list.
//
//    // Example sending a request using ListGroupPoliciesRequest.
//    req := client.ListGroupPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListGroupPolicies
func (c *Client) ListGroupPoliciesRequest(input *types.ListGroupPoliciesInput) ListGroupPoliciesRequest {
	op := &aws.Operation{
		Name:       opListGroupPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &types.ListGroupPoliciesInput{}
	}

	req := c.newRequest(op, input, &types.ListGroupPoliciesOutput{})
	return ListGroupPoliciesRequest{Request: req, Input: input, Copy: c.ListGroupPoliciesRequest}
}

// ListGroupPoliciesRequest is the request type for the
// ListGroupPolicies API operation.
type ListGroupPoliciesRequest struct {
	*aws.Request
	Input *types.ListGroupPoliciesInput
	Copy  func(*types.ListGroupPoliciesInput) ListGroupPoliciesRequest
}

// Send marshals and sends the ListGroupPolicies API request.
func (r ListGroupPoliciesRequest) Send(ctx context.Context) (*ListGroupPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupPoliciesResponse{
		ListGroupPoliciesOutput: r.Request.Data.(*types.ListGroupPoliciesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGroupPoliciesRequestPaginator returns a paginator for ListGroupPolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGroupPoliciesRequest(input)
//   p := iam.NewListGroupPoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGroupPoliciesPaginator(req ListGroupPoliciesRequest) ListGroupPoliciesPaginator {
	return ListGroupPoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListGroupPoliciesInput
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

// ListGroupPoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGroupPoliciesPaginator struct {
	aws.Pager
}

func (p *ListGroupPoliciesPaginator) CurrentPage() *types.ListGroupPoliciesOutput {
	return p.Pager.CurrentPage().(*types.ListGroupPoliciesOutput)
}

// ListGroupPoliciesResponse is the response type for the
// ListGroupPolicies API operation.
type ListGroupPoliciesResponse struct {
	*types.ListGroupPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroupPolicies request.
func (r *ListGroupPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
