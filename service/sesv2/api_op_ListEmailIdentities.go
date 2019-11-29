// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sesv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

const opListEmailIdentities = "ListEmailIdentities"

// ListEmailIdentitiesRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Returns a list of all of the email identities that are associated with your
// AWS account. An identity can be either an email address or a domain. This
// operation returns identities that are verified as well as those that aren't.
// This operation returns identities that are associated with Amazon SES and
// Amazon Pinpoint.
//
//    // Example sending a request using ListEmailIdentitiesRequest.
//    req := client.ListEmailIdentitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sesv2-2019-09-27/ListEmailIdentities
func (c *Client) ListEmailIdentitiesRequest(input *types.ListEmailIdentitiesInput) ListEmailIdentitiesRequest {
	op := &aws.Operation{
		Name:       opListEmailIdentities,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/email/identities",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListEmailIdentitiesInput{}
	}

	req := c.newRequest(op, input, &types.ListEmailIdentitiesOutput{})
	return ListEmailIdentitiesRequest{Request: req, Input: input, Copy: c.ListEmailIdentitiesRequest}
}

// ListEmailIdentitiesRequest is the request type for the
// ListEmailIdentities API operation.
type ListEmailIdentitiesRequest struct {
	*aws.Request
	Input *types.ListEmailIdentitiesInput
	Copy  func(*types.ListEmailIdentitiesInput) ListEmailIdentitiesRequest
}

// Send marshals and sends the ListEmailIdentities API request.
func (r ListEmailIdentitiesRequest) Send(ctx context.Context) (*ListEmailIdentitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEmailIdentitiesResponse{
		ListEmailIdentitiesOutput: r.Request.Data.(*types.ListEmailIdentitiesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEmailIdentitiesRequestPaginator returns a paginator for ListEmailIdentities.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEmailIdentitiesRequest(input)
//   p := sesv2.NewListEmailIdentitiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEmailIdentitiesPaginator(req ListEmailIdentitiesRequest) ListEmailIdentitiesPaginator {
	return ListEmailIdentitiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListEmailIdentitiesInput
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

// ListEmailIdentitiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEmailIdentitiesPaginator struct {
	aws.Pager
}

func (p *ListEmailIdentitiesPaginator) CurrentPage() *types.ListEmailIdentitiesOutput {
	return p.Pager.CurrentPage().(*types.ListEmailIdentitiesOutput)
}

// ListEmailIdentitiesResponse is the response type for the
// ListEmailIdentities API operation.
type ListEmailIdentitiesResponse struct {
	*types.ListEmailIdentitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEmailIdentities request.
func (r *ListEmailIdentitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
