// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opListProtectedResources = "ListProtectedResources"

// ListProtectedResourcesRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns an array of resources successfully backed up by AWS Backup, including
// the time the resource was saved, an Amazon Resource Name (ARN) of the resource,
// and a resource type.
//
//    // Example sending a request using ListProtectedResourcesRequest.
//    req := client.ListProtectedResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListProtectedResources
func (c *Client) ListProtectedResourcesRequest(input *types.ListProtectedResourcesInput) ListProtectedResourcesRequest {
	op := &aws.Operation{
		Name:       opListProtectedResources,
		HTTPMethod: "GET",
		HTTPPath:   "/resources/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListProtectedResourcesInput{}
	}

	req := c.newRequest(op, input, &types.ListProtectedResourcesOutput{})
	return ListProtectedResourcesRequest{Request: req, Input: input, Copy: c.ListProtectedResourcesRequest}
}

// ListProtectedResourcesRequest is the request type for the
// ListProtectedResources API operation.
type ListProtectedResourcesRequest struct {
	*aws.Request
	Input *types.ListProtectedResourcesInput
	Copy  func(*types.ListProtectedResourcesInput) ListProtectedResourcesRequest
}

// Send marshals and sends the ListProtectedResources API request.
func (r ListProtectedResourcesRequest) Send(ctx context.Context) (*ListProtectedResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProtectedResourcesResponse{
		ListProtectedResourcesOutput: r.Request.Data.(*types.ListProtectedResourcesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProtectedResourcesRequestPaginator returns a paginator for ListProtectedResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProtectedResourcesRequest(input)
//   p := backup.NewListProtectedResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProtectedResourcesPaginator(req ListProtectedResourcesRequest) ListProtectedResourcesPaginator {
	return ListProtectedResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListProtectedResourcesInput
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

// ListProtectedResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProtectedResourcesPaginator struct {
	aws.Pager
}

func (p *ListProtectedResourcesPaginator) CurrentPage() *types.ListProtectedResourcesOutput {
	return p.Pager.CurrentPage().(*types.ListProtectedResourcesOutput)
}

// ListProtectedResourcesResponse is the response type for the
// ListProtectedResources API operation.
type ListProtectedResourcesResponse struct {
	*types.ListProtectedResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProtectedResources request.
func (r *ListProtectedResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
