// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opListAttachedRolePolicies = "ListAttachedRolePolicies"

// ListAttachedRolePoliciesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists all managed policies that are attached to the specified IAM role.
//
// An IAM role can also have inline policies embedded with it. To list the inline
// policies for a role, use the ListRolePolicies API. For information about
// policies, see Managed Policies and Inline Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
// You can paginate the results using the MaxItems and Marker parameters. You
// can use the PathPrefix parameter to limit the list of policies to only those
// matching the specified path prefix. If there are no policies attached to
// the specified role (or none that match the specified path prefix), the operation
// returns an empty list.
//
//    // Example sending a request using ListAttachedRolePoliciesRequest.
//    req := client.ListAttachedRolePoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListAttachedRolePolicies
func (c *Client) ListAttachedRolePoliciesRequest(input *types.ListAttachedRolePoliciesInput) ListAttachedRolePoliciesRequest {
	op := &aws.Operation{
		Name:       opListAttachedRolePolicies,
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
		input = &types.ListAttachedRolePoliciesInput{}
	}

	req := c.newRequest(op, input, &types.ListAttachedRolePoliciesOutput{})
	return ListAttachedRolePoliciesRequest{Request: req, Input: input, Copy: c.ListAttachedRolePoliciesRequest}
}

// ListAttachedRolePoliciesRequest is the request type for the
// ListAttachedRolePolicies API operation.
type ListAttachedRolePoliciesRequest struct {
	*aws.Request
	Input *types.ListAttachedRolePoliciesInput
	Copy  func(*types.ListAttachedRolePoliciesInput) ListAttachedRolePoliciesRequest
}

// Send marshals and sends the ListAttachedRolePolicies API request.
func (r ListAttachedRolePoliciesRequest) Send(ctx context.Context) (*ListAttachedRolePoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAttachedRolePoliciesResponse{
		ListAttachedRolePoliciesOutput: r.Request.Data.(*types.ListAttachedRolePoliciesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAttachedRolePoliciesRequestPaginator returns a paginator for ListAttachedRolePolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAttachedRolePoliciesRequest(input)
//   p := iam.NewListAttachedRolePoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAttachedRolePoliciesPaginator(req ListAttachedRolePoliciesRequest) ListAttachedRolePoliciesPaginator {
	return ListAttachedRolePoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListAttachedRolePoliciesInput
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

// ListAttachedRolePoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAttachedRolePoliciesPaginator struct {
	aws.Pager
}

func (p *ListAttachedRolePoliciesPaginator) CurrentPage() *types.ListAttachedRolePoliciesOutput {
	return p.Pager.CurrentPage().(*types.ListAttachedRolePoliciesOutput)
}

// ListAttachedRolePoliciesResponse is the response type for the
// ListAttachedRolePolicies API operation.
type ListAttachedRolePoliciesResponse struct {
	*types.ListAttachedRolePoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAttachedRolePolicies request.
func (r *ListAttachedRolePoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
