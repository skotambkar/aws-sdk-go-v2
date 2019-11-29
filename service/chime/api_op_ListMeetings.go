// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opListMeetings = "ListMeetings"

// ListMeetingsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists up to 100 active Amazon Chime SDK meetings. For more information about
// the Amazon Chime SDK, see Using the Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide.
//
//    // Example sending a request using ListMeetingsRequest.
//    req := client.ListMeetingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListMeetings
func (c *Client) ListMeetingsRequest(input *types.ListMeetingsInput) ListMeetingsRequest {
	op := &aws.Operation{
		Name:       opListMeetings,
		HTTPMethod: "GET",
		HTTPPath:   "/meetings",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListMeetingsInput{}
	}

	req := c.newRequest(op, input, &types.ListMeetingsOutput{})
	return ListMeetingsRequest{Request: req, Input: input, Copy: c.ListMeetingsRequest}
}

// ListMeetingsRequest is the request type for the
// ListMeetings API operation.
type ListMeetingsRequest struct {
	*aws.Request
	Input *types.ListMeetingsInput
	Copy  func(*types.ListMeetingsInput) ListMeetingsRequest
}

// Send marshals and sends the ListMeetings API request.
func (r ListMeetingsRequest) Send(ctx context.Context) (*ListMeetingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMeetingsResponse{
		ListMeetingsOutput: r.Request.Data.(*types.ListMeetingsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMeetingsRequestPaginator returns a paginator for ListMeetings.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMeetingsRequest(input)
//   p := chime.NewListMeetingsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMeetingsPaginator(req ListMeetingsRequest) ListMeetingsPaginator {
	return ListMeetingsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListMeetingsInput
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

// ListMeetingsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMeetingsPaginator struct {
	aws.Pager
}

func (p *ListMeetingsPaginator) CurrentPage() *types.ListMeetingsOutput {
	return p.Pager.CurrentPage().(*types.ListMeetingsOutput)
}

// ListMeetingsResponse is the response type for the
// ListMeetings API operation.
type ListMeetingsResponse struct {
	*types.ListMeetingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMeetings request.
func (r *ListMeetingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
