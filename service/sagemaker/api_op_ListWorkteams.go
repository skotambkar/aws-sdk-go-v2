// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opListWorkteams = "ListWorkteams"

// ListWorkteamsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Gets a list of work teams that you have defined in a region. The list may
// be empty if no work team satisfies the filter specified in the NameContains
// parameter.
//
//    // Example sending a request using ListWorkteamsRequest.
//    req := client.ListWorkteamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListWorkteams
func (c *Client) ListWorkteamsRequest(input *types.ListWorkteamsInput) ListWorkteamsRequest {
	op := &aws.Operation{
		Name:       opListWorkteams,
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
		input = &types.ListWorkteamsInput{}
	}

	req := c.newRequest(op, input, &types.ListWorkteamsOutput{})
	return ListWorkteamsRequest{Request: req, Input: input, Copy: c.ListWorkteamsRequest}
}

// ListWorkteamsRequest is the request type for the
// ListWorkteams API operation.
type ListWorkteamsRequest struct {
	*aws.Request
	Input *types.ListWorkteamsInput
	Copy  func(*types.ListWorkteamsInput) ListWorkteamsRequest
}

// Send marshals and sends the ListWorkteams API request.
func (r ListWorkteamsRequest) Send(ctx context.Context) (*ListWorkteamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListWorkteamsResponse{
		ListWorkteamsOutput: r.Request.Data.(*types.ListWorkteamsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListWorkteamsRequestPaginator returns a paginator for ListWorkteams.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListWorkteamsRequest(input)
//   p := sagemaker.NewListWorkteamsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListWorkteamsPaginator(req ListWorkteamsRequest) ListWorkteamsPaginator {
	return ListWorkteamsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListWorkteamsInput
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

// ListWorkteamsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListWorkteamsPaginator struct {
	aws.Pager
}

func (p *ListWorkteamsPaginator) CurrentPage() *types.ListWorkteamsOutput {
	return p.Pager.CurrentPage().(*types.ListWorkteamsOutput)
}

// ListWorkteamsResponse is the response type for the
// ListWorkteams API operation.
type ListWorkteamsResponse struct {
	*types.ListWorkteamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListWorkteams request.
func (r *ListWorkteamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
