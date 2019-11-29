// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opListObjectParents = "ListObjectParents"

// ListObjectParentsRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists parent objects that are associated with a given object in pagination
// fashion.
//
//    // Example sending a request using ListObjectParentsRequest.
//    req := client.ListObjectParentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListObjectParents
func (c *Client) ListObjectParentsRequest(input *types.ListObjectParentsInput) ListObjectParentsRequest {
	op := &aws.Operation{
		Name:       opListObjectParents,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/parent",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListObjectParentsInput{}
	}

	req := c.newRequest(op, input, &types.ListObjectParentsOutput{})
	return ListObjectParentsRequest{Request: req, Input: input, Copy: c.ListObjectParentsRequest}
}

// ListObjectParentsRequest is the request type for the
// ListObjectParents API operation.
type ListObjectParentsRequest struct {
	*aws.Request
	Input *types.ListObjectParentsInput
	Copy  func(*types.ListObjectParentsInput) ListObjectParentsRequest
}

// Send marshals and sends the ListObjectParents API request.
func (r ListObjectParentsRequest) Send(ctx context.Context) (*ListObjectParentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListObjectParentsResponse{
		ListObjectParentsOutput: r.Request.Data.(*types.ListObjectParentsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListObjectParentsRequestPaginator returns a paginator for ListObjectParents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListObjectParentsRequest(input)
//   p := clouddirectory.NewListObjectParentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListObjectParentsPaginator(req ListObjectParentsRequest) ListObjectParentsPaginator {
	return ListObjectParentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListObjectParentsInput
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

// ListObjectParentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListObjectParentsPaginator struct {
	aws.Pager
}

func (p *ListObjectParentsPaginator) CurrentPage() *types.ListObjectParentsOutput {
	return p.Pager.CurrentPage().(*types.ListObjectParentsOutput)
}

// ListObjectParentsResponse is the response type for the
// ListObjectParents API operation.
type ListObjectParentsResponse struct {
	*types.ListObjectParentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListObjectParents request.
func (r *ListObjectParentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
