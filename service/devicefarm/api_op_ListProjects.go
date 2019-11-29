// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opListProjects = "ListProjects"

// ListProjectsRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about projects.
//
//    // Example sending a request using ListProjectsRequest.
//    req := client.ListProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListProjects
func (c *Client) ListProjectsRequest(input *types.ListProjectsInput) ListProjectsRequest {
	op := &aws.Operation{
		Name:       opListProjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListProjectsInput{}
	}

	req := c.newRequest(op, input, &types.ListProjectsOutput{})
	return ListProjectsRequest{Request: req, Input: input, Copy: c.ListProjectsRequest}
}

// ListProjectsRequest is the request type for the
// ListProjects API operation.
type ListProjectsRequest struct {
	*aws.Request
	Input *types.ListProjectsInput
	Copy  func(*types.ListProjectsInput) ListProjectsRequest
}

// Send marshals and sends the ListProjects API request.
func (r ListProjectsRequest) Send(ctx context.Context) (*ListProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProjectsResponse{
		ListProjectsOutput: r.Request.Data.(*types.ListProjectsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProjectsRequestPaginator returns a paginator for ListProjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProjectsRequest(input)
//   p := devicefarm.NewListProjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProjectsPaginator(req ListProjectsRequest) ListProjectsPaginator {
	return ListProjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListProjectsInput
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

// ListProjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProjectsPaginator struct {
	aws.Pager
}

func (p *ListProjectsPaginator) CurrentPage() *types.ListProjectsOutput {
	return p.Pager.CurrentPage().(*types.ListProjectsOutput)
}

// ListProjectsResponse is the response type for the
// ListProjects API operation.
type ListProjectsResponse struct {
	*types.ListProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProjects request.
func (r *ListProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
