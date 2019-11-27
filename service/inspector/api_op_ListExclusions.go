// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
)

const opListExclusions = "ListExclusions"

// ListExclusionsRequest returns a request value for making API operation for
// Amazon Inspector.
//
// List exclusions that are generated by the assessment run.
//
//    // Example sending a request using ListExclusionsRequest.
//    req := client.ListExclusionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/ListExclusions
func (c *Client) ListExclusionsRequest(input *types.ListExclusionsInput) ListExclusionsRequest {
	op := &aws.Operation{
		Name:       opListExclusions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListExclusionsInput{}
	}

	req := c.newRequest(op, input, &types.ListExclusionsOutput{})
	return ListExclusionsRequest{Request: req, Input: input, Copy: c.ListExclusionsRequest}
}

// ListExclusionsRequest is the request type for the
// ListExclusions API operation.
type ListExclusionsRequest struct {
	*aws.Request
	Input *types.ListExclusionsInput
	Copy  func(*types.ListExclusionsInput) ListExclusionsRequest
}

// Send marshals and sends the ListExclusions API request.
func (r ListExclusionsRequest) Send(ctx context.Context) (*ListExclusionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListExclusionsResponse{
		ListExclusionsOutput: r.Request.Data.(*types.ListExclusionsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListExclusionsRequestPaginator returns a paginator for ListExclusions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListExclusionsRequest(input)
//   p := inspector.NewListExclusionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListExclusionsPaginator(req ListExclusionsRequest) ListExclusionsPaginator {
	return ListExclusionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListExclusionsInput
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

// ListExclusionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListExclusionsPaginator struct {
	aws.Pager
}

func (p *ListExclusionsPaginator) CurrentPage() *types.ListExclusionsOutput {
	return p.Pager.CurrentPage().(*types.ListExclusionsOutput)
}

// ListExclusionsResponse is the response type for the
// ListExclusions API operation.
type ListExclusionsResponse struct {
	*types.ListExclusionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListExclusions request.
func (r *ListExclusionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
