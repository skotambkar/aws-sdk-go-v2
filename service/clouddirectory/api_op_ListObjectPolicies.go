// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opListObjectPolicies = "ListObjectPolicies"

// ListObjectPoliciesRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns policies attached to an object in pagination fashion.
//
//    // Example sending a request using ListObjectPoliciesRequest.
//    req := client.ListObjectPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectPolicies
func (c *Client) ListObjectPoliciesRequest(input *types.ListObjectPoliciesInput) ListObjectPoliciesRequest {
	op := &aws.Operation{
		Name:       opListObjectPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/policy",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListObjectPoliciesInput{}
	}

	req := c.newRequest(op, input, &types.ListObjectPoliciesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListObjectPoliciesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListObjectPoliciesRequest{Request: req, Input: input, Copy: c.ListObjectPoliciesRequest}
}

// ListObjectPoliciesRequest is the request type for the
// ListObjectPolicies API operation.
type ListObjectPoliciesRequest struct {
	*aws.Request
	Input *types.ListObjectPoliciesInput
	Copy  func(*types.ListObjectPoliciesInput) ListObjectPoliciesRequest
}

// Send marshals and sends the ListObjectPolicies API request.
func (r ListObjectPoliciesRequest) Send(ctx context.Context) (*ListObjectPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectPoliciesResponse{
		ListObjectPoliciesOutput: r.Request.Data.(*types.ListObjectPoliciesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectPoliciesRequestPaginator returns a paginator for ListObjectPolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectPoliciesRequest(input)
//   p := clouddirectory.NewListObjectPoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectPoliciesPaginator(req ListObjectPoliciesRequest) ListObjectPoliciesPaginator {
	return ListObjectPoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListObjectPoliciesInput
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

// ListObjectPoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectPoliciesPaginator struct {
	aws.Pager
}

func (p *ListObjectPoliciesPaginator) CurrentPage() *types.ListObjectPoliciesOutput {
	return p.Pager.CurrentPage().(*types.ListObjectPoliciesOutput)
}

// ListObjectPoliciesResponse is the response type for the
// ListObjectPolicies API operation.
type ListObjectPoliciesResponse struct {
	*types.ListObjectPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectPolicies request.
func (r *ListObjectPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
