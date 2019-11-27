// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opSearchContacts = "SearchContacts"

// SearchContactsRequest returns a request value for making API operation for
// Alexa For Business.
//
// Searches contacts and lists the ones that meet a set of filter and sort criteria.
//
//    // Example sending a request using SearchContactsRequest.
//    req := client.SearchContactsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/SearchContacts
func (c *Client) SearchContactsRequest(input *types.SearchContactsInput) SearchContactsRequest {
	op := &aws.Operation{
		Name:       opSearchContacts,
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
		input = &types.SearchContactsInput{}
	}

	req := c.newRequest(op, input, &types.SearchContactsOutput{})
	return SearchContactsRequest{Request: req, Input: input, Copy: c.SearchContactsRequest}
}

// SearchContactsRequest is the request type for the
// SearchContacts API operation.
type SearchContactsRequest struct {
	*aws.Request
	Input *types.SearchContactsInput
	Copy  func(*types.SearchContactsInput) SearchContactsRequest
}

// Send marshals and sends the SearchContacts API request.
func (r SearchContactsRequest) Send(ctx context.Context) (*SearchContactsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchContactsResponse{
		SearchContactsOutput: r.Request.Data.(*types.SearchContactsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchContactsRequestPaginator returns a paginator for SearchContacts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchContactsRequest(input)
//   p := alexaforbusiness.NewSearchContactsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchContactsPaginator(req SearchContactsRequest) SearchContactsPaginator {
	return SearchContactsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.SearchContactsInput
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

// SearchContactsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchContactsPaginator struct {
	aws.Pager
}

func (p *SearchContactsPaginator) CurrentPage() *types.SearchContactsOutput {
	return p.Pager.CurrentPage().(*types.SearchContactsOutput)
}

// SearchContactsResponse is the response type for the
// SearchContacts API operation.
type SearchContactsResponse struct {
	*types.SearchContactsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchContacts request.
func (r *SearchContactsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
